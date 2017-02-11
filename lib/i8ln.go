package lib

import (
	"log"
	"os"
	"path/filepath"

	"github.com/leonelquinteros/gotext"
)

var (
	dom  = "default"
	lang = "en_UK"
	lib  = "locale"
)

func init() {
	initGetText()
}

func initGetText() {
	cd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot determine path of current directory")
	}
	cd, err = filepath.Abs(cd)
	if err != nil {
		log.Fatal("cannot determine absolute path of current directory")
	}
	lib = filepath.Clean(filepath.Join(cd, lib))
	gotext.Configure(lib, lang, dom)
}

// POLocalizer ...
type POLocalizer struct{}

// GetText ...
func (l POLocalizer) GetText(word string, lc string) string {
	gotext.SetLanguage(lc)
	return gotext.Get(word)
}
