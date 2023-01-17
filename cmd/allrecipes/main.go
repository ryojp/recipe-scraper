package main

import (
	"flag"
	"log"
	"regexp"
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

	flag.Parse()

	// Instantiate the collector
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(`https://www.allrecipes.com/recipe/\d+/.*`),
		),
		colly.CacheDir("./.cache/allrecipes"),
		colly.MaxDepth(*maxDepth),
		colly.Async(),
	)

	extensions.RandomMobileUserAgent(c)
	extensions.Referer(c)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
	})

	recipes := allrecipes.Recipes{}

	// Extract details of the course
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		log.Println("Recipe found", e.Request.URL)

		recipe := allrecipes.ParseRecipe(e)
		if recipe == nil {
			return
		}

		recipes.Add(recipe)

		// recursively visit recipe pages shown in section "You'll Also Love"
		for _, url := range e.ChildAttrs("#recirc-section__card-list-1_1-0 a", "href") {
			e.Request.Visit(url)
		}
	})

	// Start scraping
	c.Visit(*startURL)

	// Wait until all the threads exit
	c.Wait()

	// Export the collected recipes to JSON
	recipes.DumpJSON("allrecipes.json")
}
