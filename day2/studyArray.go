package main

import "fmt"
import "errors"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	book := Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}
	testFunc(&book)





}

func testFunc(book1 *Books) {
	book1.author = "yanzhaoyu"
	book1.book_id = 8888888

}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
