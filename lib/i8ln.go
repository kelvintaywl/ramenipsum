package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/leonelquinteros/gotext"
)

var (
	dom  = "default"
	lang = "en_UK"
	lib  = fmt.Sprintf("%s/src/github.com/kelvintaywl/ramenipsum/locale", os.Getenv("GOPATH"))
)

func init() {
	initGetText()
}

func initGetText() {
	gotext.Configure(lib, lang, dom)
	log.Printf("lang: %s", gotext.GetLanguage())
	log.Printf("salt: %s", gotext.Get("salt"))
}

// Localizer ...
type Localizer interface {
	GetText(word string, lc locale) string
}

// POLocalizer ...
type POLocalizer struct{}

// GetText ...
func (l POLocalizer) GetText(word string, lc string) string {
	gotext.SetLanguage(lc)
	return gotext.Get(word)
}
