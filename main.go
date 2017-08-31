package main

import (
	"fmt"

	"github.com/wangbin/jiebago"
	"github.com/sajari/word2vec"
	"log"
)

var seg jiebago.Segmenter

func init() {
	seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
	for word := range ch {
		fmt.Printf(" %s /", word)
	}
	fmt.Println()
}
func wordanalyse(){
	c := word2vec.Client{Addr: "localhost:1234"}

	// Create an expression.
	expr := word2vec.Expr{}
	expr.Add(1, "War")

	// Find the most similar result by cosine similarity.
	matches, err := c.CosN(expr, 10)
	if err != nil {
		log.Fatalf("error evaluating cosine similarity: %v", err)
	}else{
		fmt.Println(matches)
	}
}
func fenci() {
	fmt.Print("【全模式】：")
	print(seg.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	print(seg.Cut("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	print(seg.Cut("他来到了网易杭研大厦", true))

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("国家旅游局责成云南查明女大学生旅游被打事件", true))
}
func main()  {
	wordanalyse()
}