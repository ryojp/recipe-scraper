package allrecipes

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	pluralize "github.com/gertd/go-pluralize"
	"github.com/gocolly/colly/v2"
)

// plu is an instance of `pluraliez` pkg to convert plural nouns to singular
var plu = pluralize.NewClient()

// ParseRecipe is the main parser function that parses the detailed recipe page
func ParseRecipe(e *colly.HTMLElement) *Recipe {
	imageURL := e.ChildAttr("div.article-content img", "src")
	if imageURL == "" {
		imageURL = e.ChildAttr(".recipe__steps-content img", "data-src")
	}

	recipe := &Recipe{
		Title:    e.ChildText("h1"),
		Summary:  e.ChildText(".article-subheading"),
		URL:      e.Request.URL.String(),
		ImageURL: imageURL,
		Author:   e.ChildText(".mntl-bylines__item a"),
	}
	recipe.parseStats(e)
	recipe.parseTimeAndServings(e)
	recipe.parseIngredients(e)
	recipe.parseDirections(e)
	recipe.parseNutrition(e)

	return recipe
}

// parseStats populates `Rating`, `NumRatings`, `NumReviews`, and `NumPhotos“
func (recipe *Recipe) parseStats(e *colly.HTMLElement) error {
	numRatings := e.ChildText(".mntl-recipe-review-bar__rating-count")
	if numRatings != "" {
		numRatings = numRatings[1 : len(numRatings)-1] // tirm parenthesis
	}

	recipe.Rating = strings.TrimSpace(e.ChildText(".mntl-recipe-review-bar__rating"))
	recipe.NumRatings = toInt(numRatings)
	recipe.NumReviews = toInt(strings.Split(e.ChildText(".mntl-recipe-review-bar__comment-count"), " ")[0])
	recipe.NumPhotos = toInt(strings.Split(e.ChildText(".recipe-review-bar__photo-count"), " ")[0])

	return nil
}

// parseTimeAndServings populates `PrepTime`, `CookTime`, `TotalTime`, and `Servings`
func (recipe *Recipe) parseTimeAndServings(e *colly.HTMLElement) error {
	texts := strings.Split(e.ChildText(".recipe-details .mntl-recipe-details__content"), "\n")
	mode := ""
	for _, text := range texts {
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		} else if strings.Contains(text, ":") {
			mode = text
		} else {
			switch mode {
			case "Prep Time:":
				recipe.PrepTime = toMinutes(text)
			case "Cook Time:":
				recipe.CookTime = toMinutes(text)
			case "Total Time:":
				recipe.TotalTime = toMinutes(text)
			case "Servings:":
				servings, _ := strconv.Atoi(text)
				recipe.Servings = servings
			default:
				// do nothing
			}
		}
	}
	return nil
}

// parseIngredients populates `ingredients` field
func (recipe *Recipe) parseIngredients(e *colly.HTMLElement) error {
	selection := e.DOM.Find(".mntl-structured-ingredients__list")

	var ingredients []Ingredient
	selection.Children().Each(func(i int, s *goquery.Selection) {
		q := s.Find("span:nth-of-type(1)").Text()
		q = vulgarToDecimal(q)

		unit := s.Find("span:nth-of-type(2)").Text()
		unit = regexp.MustCompile(`\(.*\)`).ReplaceAllString(unit, "")
		unit = strings.TrimSpace(unit)
		unit = plu.Singular(unit)

		name := s.Find("span:nth-of-type(3)").Text()
		name = plu.Singular(name)

		ingredients = append(ingredients, Ingredient{
			Quantity: q,
			Unit:     unit,
			Name:     name,
		})
	})
	recipe.Ingredients = ingredients

	return nil
}

// parseDirections populates `directions` field
func (recipe *Recipe) parseDirections(e *colly.HTMLElement) error {
	selection := e.DOM.Find(".recipe__steps-content ol")

	var directions []string
	selection.Children().Each(func(_ int, s *goquery.Selection) {
		direction := strings.TrimSpace(s.Find("p").Text())
		directions = append(directions, direction)
	})
	recipe.Directions = directions

	return nil
}

// parseNutrition populates `Calories`, `Fat`, `Carbs`, and `Protein`
func (recipe *Recipe) parseNutrition(e *colly.HTMLElement) error {
	selection := e.DOM.Find(".mntl-nutrition-facts-summary tbody")

	selection.Children().Each(func(_ int, s *goquery.Selection) {
		value := s.Find("td:first-of-type").Text()
		value = strings.TrimSpace(value)
		value = strings.ReplaceAll(value, "g", "")
		v := toInt(value)

		switch s.Find("td:last-of-type").Text() {
		case "Calories":
			recipe.Calories = v
		case "Fat":
			recipe.Fat = v
		case "Carbs":
			recipe.Carbs = v
		case "Protein":
			recipe.Protein = v
		default:
			// do nothing
		}
	})

	return nil
}
