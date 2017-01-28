package lib

import (
	"math/rand"
	"strings"
)

var ipsum = [...]string{
	"soy sauce",
	"salt",
	"miso",
	"butter",
	"roasted pork slices",
	"flavoured oil",
	"yuzu",
	"seasoned egg",
	"nori",
	"pork bones",
	"bean sprouts",
	"chopped onions",
	"chilli",
	"chicken stock",
	"kamaboko",
	"spicy bean paste",
	"bamboo slices",
	"sesame oil",
	"minced garlic",
	"corn",
	"pork cubes",
	"scallions",
	"spinach",
	"leek",
	"curry",
	"vinegar",
	"ginger",
	"lard",
	"fish broth",
	"rice",
	"ground black pepper",
	"soy milk",
	"wood ear mushroom",
	"toasted sesame seeds",
	"mustard greens",
	"Asahikawa",
	"Hakodate",
	"Sapporo",
	"Yokohama",
	"Hakata",
	"Tokyo",
	"Tokushima",
	"Kagoshima",
	"Wakayama",
	"Nagoya",
	"Kumamoto",
	"ramen burger",
	"tsukemen",
	"abura soba",
	"Nissin instant cup ramen",
}

// Generate implements a generator of a randomized text
type Generate func() string

// Operation implements an operation onto an input text
type Operation func(string) string

// randIntFromRange returns a random integer between
// the range [min, max]. Note that it assumes min < max
func RandIntFromRange(min, max int) int {
	return rand.Intn(max-min) + min + 1
}

func RandomWord() string {
	// pick a random word from the ipsum list
	return ipsum[rand.Intn(len(ipsum))]
}

// stylizeSentence implements the Operation function
// and formats a text with capitalization and punctuation
func stylizeSentence(s string) string {
	return strings.ToUpper(s[:1]) + s[1:] + "."
}

// gen returns a specified length of text elements, conjoined by
// a separator. It uses the input Generate function to build elements.
// This is an attempt at using compositions when simplfying the task
// of generating paragraphs of sentences of random words.
func gen(g Generate, o Operation, size int, sep string) string {
	var elems []string
	for i := 0; i < size; i++ {
		elems = append(elems, g())
	}
	res := strings.Join(elems, sep)
	if o != nil {
		return o(res)
	}
	return res
}

func RandomSentence() string {
	s := RandIntFromRange(7, 19)
	return gen(RandomWord, stylizeSentence, s, " ")
}

func RandomParagraph() string {
	s := RandIntFromRange(3, 7)
	return gen(RandomSentence, nil, s, " ")
}

func RandomText(x int) string {
	return gen(RandomParagraph, nil, x, "\n\n")
}
