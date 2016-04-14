package main

import (
	"io/ioutil"
	"strings"
)

const delim = "%"

// Interface that represents a "fortune cookie jar": it can be implemented
// either as Unix file, database, etc.
type CookieJar interface {
  // Loads the contents of the jar (i.e. the cookies)
	Load()

	// Returns the total number of cookies from the jar.
	Size() int

	// Returns the cookie from the position "pos", where "pos"
	GetAt(pos int) string
}

// Represents a disk fortune jarsitory - e.g. Unix fortune file.
type DiskCookieJar struct {
	name, path string
	fortunes   []string
}

// Creates a new disk fortune jarsitory.
func NewDiskCookieJar(name, path string) CookieJar {
	jar := &DiskCookieJar{name: name, path: path}

	return jar
}

func (jarsitory *DiskCookieJar) Load() {
	f, err := ioutil.ReadFile(jarsitory.path)
	if err != nil {
		panic("Unable to open the fortune file: " + jarsitory.path)
	}

	fileContents := string(f)
	if len(fileContents) == 0 {
		panic("There are no fortune cookies.")
	}

	fortuneCookies := strings.Split(fileContents, delim)

	jarsitory.fortunes = fortuneCookies
}

func (jarsitory *DiskCookieJar) Size() int {
	return len(jarsitory.fortunes)
}

func (jarsitory *DiskCookieJar) GetAt(pos int) string {
	return jarsitory.fortunes[pos]
}
