package main

import (
	"fmt"
	"log"
	"os"

	"code.sajari.com/word2vec"
)

func main() {
	file, err := os.Open("")
	if err != nil {
		panic(err)
	}

	model, err := word2vec.FromReader(file)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	expr := word2vec.Expr{}
	expr.Add(1, "king")
	expr.Add(-1, "man")
	expr.Add(1, "woman")

	matches, err := model.CosN(expr, 1)
	if err != nil {
		panic(err)
	}

	fmt.Sprintln(matches)
	fmt.Sprintln("HI")
}
