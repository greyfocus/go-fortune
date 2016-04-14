package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

func getFortuneRepositories(baseDir string) (fortuneRepositories []CookieJar) {
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		panic("unable to open directory: " + baseDir)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		jar := NewDiskCookieJar(filepath.Base(file.Name()),
			filepath.Join(baseDir, file.Name()))
		fortuneRepositories = append(fortuneRepositories, jar)
	}

	return fortuneRepositories
}

func pickFortune(jars []CookieJar) string {
	if len(jars) == 0 {
		panic("No fortune files found.")
	}

	// pick a random jar
	jar := jars[rand.Intn(len(jars))]
	jar.Load()

	return jar.GetAt(rand.Intn(jar.Size()))
}

func handleErrors() {
	if r := recover(); r != nil {
		fmt.Printf("Error: %s\n", r)
	}
}

func main() {
	// Error handling logic.
	defer handleErrors()

	fortuneRepositoriesDir := flag.String("d", "repo",
		"the location of fortune jarsitory directory")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	fortuneRepositories := getFortuneRepositories(*fortuneRepositoriesDir)
	fortune := pickFortune(fortuneRepositories)

	fmt.Println(fortune)
}
