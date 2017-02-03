package lib

import (
	"math/rand"
	"strings"
)

// Words is a list of ramen-related words
// Feel free to add more words to this list
// more ingredients, more flavour!
var Words = [...]string{
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

// generate implements a generator of a randomized text
type generate func() string

// operation implements an operation onto an input text
type operation func(string) string

func stylizeSentence(s string) string {
	// formats a text with capitalization and punctuation
	return strings.ToUpper(s[:1]) + s[1:] + "."
}

// gen returns a specified length of text elements, conjoined by
// a separator. It uses the input Generate function to build elements.
// This is an attempt at using compositions when simplfying the task
// of generating paragraphs of sentences of random words.
func gen(g generate, o operation, size int, sep string) string {
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

// RandIntFromRange returns a random integer between
// the range [min, max]. Note that it assumes min < max
func RandIntFromRange(min, max int) int {
	return rand.Intn(max-min) + min + 1
}

// RandomWord picks and returns a random word from a predefined list
func RandomWord() string {
	// pick a random word from the list
	return Words[rand.Intn(len(Words))]
}

// RandomSentence returns a sentence of random words, complete with punctuation.
func RandomSentence() string {
	s := RandIntFromRange(7, 19)
	return gen(RandomWord, stylizeSentence, s, " ")
}

// RandomParagraph returns a paragraphs made of random words.
func RandomParagraph() string {
	s := RandIntFromRange(3, 7)
	return gen(RandomSentence, nil, s, " ")
}

// RandomText returns a blob of text (N paragraphs) of random words
func RandomText(n int) string {
	return gen(RandomParagraph, nil, n, "\n\n")
}
