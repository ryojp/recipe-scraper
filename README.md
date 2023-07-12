# Recipe Scraper

Using [Colly](https://github.com/gocolly/colly) to scrape recipe websites.  
Currently only supporting [www.allrecipes.com](https://www.allrecipes.com/).

Inspired by [shaansubbaiah/allrecipes-scraper](https://github.com/shaansubbaiah/allrecipes-scraper)

## Prerequisites
* Installation of `go`

## Usage
* Example:
  ```sh
  go run ./cmd/allrecipes -depth 3 -parallel 2 -delay 500 -start https://www.allrecipes.com/recipe/20144/banana-banana-bread/
  ```
  When I ran this command, I got 77 recipes. The output JSON is located at `allrecipes.json` by default.
* `help` available:
  ```sh
  go run ./cmd/allrecipes -help
  ```

## Example
```json
[
  {
    "title": "Fried Chicken Sandwich",
    "summary": "These fried chicken sandwiches are pretty darn close to those Chick-fil-A sandwiches found in the mall food court. I like to make them on Sundays. Serve with waffle fries!",
    "url": "https://www.allrecipes.com/recipe/275490/fried-chicken-sandwich/",
    "image_url": "https://www.allrecipes.com/thmb/a-vjxj02rR2r16Epq5ul6sBBub0=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/2505573-32909af6c23843ab92d4fe6666580956.jpg",
    "author": "RainbowJewels",
    "ingredients": [
      {
        "quantity": "3",
        "unit": "",
        "name": "skinless, boneless chicken breast half"
      },
      {
        "quantity": "12",
        "unit": "",
        "name": "dill pickle slices, with brine"
      },
      {
        "quantity": "",
        "unit": "",
        "name": "peanut oil for frying"
      },
      {
        "quantity": "1",
        "unit": "cup",
        "name": "all-purpose flour"
      },
      {
        "quantity": "2",
        "unit": "tablespoon",
        "name": "powdered sugar"
      },
      {
        "quantity": "1",
        "unit": "teaspoon",
        "name": "salt"
      },
      {
        "quantity": "",
        "unit": "",
        "name": "paprika"
      },
      {
        "quantity": "1/2",
        "unit": "teaspoon",
        "name": "ground black pepper"
      },
      {
        "quantity": "1/2",
        "unit": "teaspoon",
        "name": "celery salt"
      },
      {
        "quantity": "1/2",
        "unit": "teaspoon",
        "name": "dried basil"
      },
      {
        "quantity": "2",
        "unit": "large",
        "name": "egg"
      },
      {
        "quantity": "",
        "unit": "",
        "name": "milk"
      },
      {
        "quantity": "2",
        "unit": "tablespoon",
        "name": "butter, softened"
      },
      {
        "quantity": "6",
        "unit": "",
        "name": "sandwich bun"
      }
    ],
    "directions": [
      "Place each chicken breast between two sheets of heavy plastic on a solid, level surface. Firmly pound chicken with the smooth side of a meat mallet to an even thickness. Cut each breast in half and place in a bowl. Pour pickle brine on top and let sit for 1 hour.",
      "Meanwhile, mix flour, powdered sugar, salt, paprika, pepper, celery salt, and basil in a bowl.",
      "Whisk eggs and milk together in another bowl.",
      "Drain liquid from the chicken and rinse off briefly. Shake any excess liquid from chicken.",
      "Heat oil in a deep-fryer or large skillet to 350 degrees F (175 degrees C).",
      "Dip each piece of chicken into the flour mixture, then into the egg-milk mixture, then coat again with the flour. Dip each piece again, if desired.",
      "Deep-fry chicken, a few pieces at a time, until golden brown, 2 to 3 minutes per side. An instant-read thermometer inserted into the center should read at least 165 degrees F (74 degrees C).",
      "Let chicken rest on a cooling rack set over a baking sheet for 5 minutes. While chicken is resting, lightly butter each sandwich bun and toast them in a skillet.",
      "Add 2 pickle slices to the bottom of each bun, then a chicken piece, then top with the other bun. Repeat to make remaining sandwiches."
    ],
    "prep_time_min": 20,
    "cook_time_min": 10,
    "total_time_min": 95,
    "servings": 6,
    "rating": "4.8",
    "num_ratings": 12,
    "num_reviews": 9,
    "num_photos": 4,
    "calories_kcal": 405,
    "fat_g": 20,
    "carbs_g": 183,
    "protein_g": 17
  }
]
```
