package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type randomFn func() string
type mapFn func(string) string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randIntFromRange(min, max int) int {
	return rand.Intn(max-min) + min + 1
}

func randomWord() string {
	return ipsum[rand.Intn(len(ipsum))]
}

func stylizeSentence(s string) string {
	return strings.ToUpper(s[:1]) + s[1:] + "."
}

func ran(size int, rf randomFn, mf mapFn, sep string) string {
	var elems []string
	for i := 0; i < size; i++ {
		elems = append(elems, rf())
	}
	res := strings.Join(elems, sep)
	if mf != nil {
		return mf(res)
	}
	return res
}

func randomSentence() string {
	s := randIntFromRange(7, 19)
	return ran(s, randomWord, stylizeSentence, " ")
}

func randomParagraph() string {
	s := randIntFromRange(3, 7)
	return ran(s, randomSentence, nil, " ")
}

func randomText(x int) string {
	return ran(x, randomParagraph, nil, "\n\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
	var str string
	i, err := strconv.Atoi(r.URL.Path[len("/ramen/"):])
	if err != nil {
		str = randomText(1)
	} else {
		str = randomText(i)
	}
	fmt.Fprint(w, str)
}

func main() {
	http.HandleFunc("/ramen/", handler)
	http.ListenAndServe(":9000", nil)
}
