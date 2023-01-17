package allrecipes

import (
	"regexp"
	"strconv"
	"strings"
)

// toMinutes converts a string like `6 hrs 10 mins` to `370`
func toMinutes(text string) int {
	hrStr := strings.Split(regexp.MustCompile(`\d+ hrs?`).FindString(text), " ")[0]
	hr, _ := strconv.Atoi(hrStr) // Atoi returns 0 if the arg is empty ("")
	minStr := strings.Split(regexp.MustCompile(`\d+ mins?`).FindString(text), " ")[0]
	min, _ := strconv.Atoi(minStr)
	return 60*hr + min
}

// vulgarToDecimal converts e.g., ¼ to 1/4
func vulgarToDecimal(vulgar string) string {
	r := strings.NewReplacer(
		"¼", "1/4",
		"½", "1/2",
		"¾", "3/4",
		"⅐", "1/7",
		"⅑", "1/9",
		"⅒", "1/10",
		"⅓", "1/3",
		"⅔", "2/3",
		"⅕", "1/5",
		"⅖", "2/5",
		"⅗", "3/5",
		"⅘", "4/5",
		"⅙", "1/6",
		"⅚", "5/6",
		"⅛", "1/8",
		"⅜", "3/8",
		"⅝", "5/8",
		"⅞", "7/8",
	)
	return r.Replace(vulgar)
}

// toInt converts the string to int, ignoring potential errors
// Returns 0 if the argument is an empty string
func toInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
