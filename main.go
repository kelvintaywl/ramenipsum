package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func handler(c *gin.Context) {
	para := c.Param("num")
	n, err := strconv.Atoi(para)
	if err != nil {
		n = 1
	}
	c.String(http.StatusOK, randomText(n))
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl",
		gin.H{
			"title": "Ramen Ipsum: A brothier Lorem Ipsum generator",
		})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/assets", "assets")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)
	router.GET("/paragraphs/:num", handler)

	router.Run(":" + port)
}
