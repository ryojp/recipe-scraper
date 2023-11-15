package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/ryojp/recipe_scraper/internal/apps/allrecipes"
)

func main() {
	maxDepth := flag.Int("depth", 4, "Max depth for the crawler")
	startURL := flag.String(
		"start",
		"https://www.allrecipes.com/recipe/176132/slow-cooker-buffalo-chicken-sandwiches/",
		"Initial URL to crawl",
	)
	delay := flag.Int("delay", 2000, "Milliseconds between requests")
	parallel := flag.Int("parallel", 2, "Number of concurrnet requests")
	outfile := flag.String("out", "allrecipes.json", "Filename for the output JSON")
	cacheDir := flag.String("cache", "./.cache/allrecipes", "Cache directory for scraping")

	flag.Parse()

	// Instantiate the collector
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(`https://www.allrecipes.com/recipe/\d+/.*`),
		),
		colly.CacheDir(*cacheDir),
		colly.MaxDepth(*maxDepth),
		colly.Async(),
	)

	extensions.RandomMobileUserAgent(c)
	extensions.Referer(c)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: *parallel,
		RandomDelay: time.Duration(*delay) * time.Millisecond,
	})

	recipes := allrecipes.Recipes{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Dump collected recipes and exit upon receiving signals
	go func() {
		sig := <-sigs
		log.Println("Received signal", sig.String())
		recipes.DumpJSON(*outfile)
		os.Exit(1)
	}()

	// Extract details of the course
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		log.Println("Recipe found", e.Request.URL)

		recipe := allrecipes.ParseRecipe(e)
		if recipe == nil {
			return
		}

		recipes.Add(recipe)

		// recursively visit recipe pages shown in section "You'll Also Love"
		for _, url := range e.ChildAttrs(".recirc-section a", "href") {
			if !recipes.Exists(url) {
				e.Request.Visit(url) // visit only if the url is new
			}
		}
	})

	// Start scraping
	c.Visit(*startURL)

	// Wait until all the threads exit
	c.Wait()

	// Export the collected recipes to JSON
	recipes.DumpJSON(*outfile)
}
