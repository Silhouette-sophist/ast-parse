package main

import (
	"ast-parse/parse"
	"log"
	"os"
)

func main() {
	path := "/Users/silhouette/codeworks/ast-parse"
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalf("read stat error %v", err)
		return
	}
	if stat.IsDir() {
		log.Println("dir parse")
		err := parse.GoDir(path)
		if err != nil {
			log.Fatalf("parseGoDir error %v", err)
			return
		}
	} else {
		log.Println("file parse")
		err := parse.GoFile(path)
		if err != nil {
			log.Fatalf("parseGoFile error %v", err)
			return
		}
	}
}
