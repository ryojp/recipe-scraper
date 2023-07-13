package allrecipes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// Recipe stores information about a recipe page
type Recipe struct {
	Title       string       `json:"title"`
	Summary     string       `json:"summary"`
	URL         string       `json:"url"`
	ImageURL    string       `json:"image_url"`
	Author      string       `json:"author"`
	Ingredients []Ingredient `json:"ingredients"`
	Directions  []string     `json:"directions"`
	// time and servings
	PrepTime  int `json:"prep_time_min"`
	CookTime  int `json:"cook_time_min"`
	TotalTime int `json:"total_time_min"`
	Servings  int `json:"servings"`
	// stats
	Rating     string `json:"rating"`
	NumRatings int    `json:"num_ratings"`
	NumReviews int    `json:"num_reviews"`
	NumPhotos  int    `json:"num_photos"`
	// nutritions (per serving)
	Calories int `json:"calories_kcal"`
	Fat      int `json:"fat_g"`
	Carbs    int `json:"carbs_g"`
	Protein  int `json:"protein_g"`
}

// Ingredient stores the name, quantity, and unit
type Ingredient struct {
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
	Name     string `json:"name"`
}

// Recipes is a thread-safe recipe map
type Recipes struct {
	recipes map[string]Recipe
	mux     sync.Mutex
}

// Add adds a given recipe to the recipes map if not yet added
func (recipes *Recipes) Add(recipe *Recipe) {
	recipes.mux.Lock()
	if recipes.recipes == nil {
		recipes.recipes = make(map[string]Recipe)
	}
	_, ok := recipes.recipes[recipe.URL]
	if !ok {
		recipes.recipes[recipe.URL] = *recipe
	}
	recipes.mux.Unlock()
}

// DumpJSON dumps the recipes into a json file
func (recipes *Recipes) DumpJSON(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", filename, err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// To dump recipes as an array, convert recipes map to an array
	recipes.mux.Lock()
	var recipeSlice []Recipe
	for _, recipe := range recipes.recipes {
		recipeSlice = append(recipeSlice, recipe)
	}
	recipes.mux.Unlock()

	fmt.Printf("Dumping the collected %v recipes into %q...\n", len(recipeSlice), filename)

	// Dump json to the standard output
	err = enc.Encode(recipeSlice)
	if err != nil {
		log.Fatalf("Failed to dump to JSON: %s\n", err)
		return
	}
}
