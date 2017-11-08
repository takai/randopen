package main

import (
	"flag"
	"os"
	"fmt"
	"path/filepath"
	"log"
	"time"
	"math/rand"
	"os/exec"
)

func selectFileRandomly(path *string, suffix *string) string {
	glob := fmt.Sprintf("%s/*.%s", *path, *suffix)
	files, err := filepath.Glob(glob)
	if err != nil {
		log.Fatal(err)
	}
	file := files[rand.Intn(len(files))]
	return file
}

func main() {
	rand.Seed(time.Now().Unix())

	path := flag.String("path", "", "Target directory (required)")
	suffix := flag.String("suffix", "wmv", "Target suffix")
	flag.Parse()

	if *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file := selectFileRandomly(path, suffix)

	cmd := exec.Command("open", file)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
