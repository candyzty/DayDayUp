package main

import (
	"DayDayUp/JUN/day_04/struct_01"
	"fmt"
	//"day_04/struct_01/Books"
)

func main() {
	book := struct_01.NewBook(1000, "Golang", "James", "Golang Book")
	//book.Id = 1000
	//book.Title = "java"
	fmt.Println(book.String())
	//struct_01.Books.Id = 1000

	struct_01.RefTag(*book, 2)
	struct_01.InitTechBook()
}
