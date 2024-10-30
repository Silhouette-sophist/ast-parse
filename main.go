package main

import (
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
		err := parseGoDir(path)
		if err != nil {
			log.Fatalf("parseGoDir error %v", err)
			return
		}
	} else {
		log.Println("file parse")
		err := parseGoFile(path)
		if err != nil {
			log.Fatalf("parseGoFile error %v", err)
			return
		}
	}
}
