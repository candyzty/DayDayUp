package struct_01

import (
	"fmt"
	"reflect"
)

type Books struct {
	Id      int    "this is a book id"
	Title   string "this is  a title"
	Author  string "this is a author"
	Subject string "this is s subject"
}

type techBook struct {
	Cat   string
	int   // 匿名字段
	Books // 匿名字段，内嵌结构体
}

func NewBook(Id int, Title, Author, Subject string) *Books {
	return &Books{Id, Title, Author, Subject}
}

func (book *Books) String() string {
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n", book.Id, book.Title, book.Author, book.Subject)

}

func (book *Books) GetTitle() string {
	return book.Title
}

func printBook(book *Books) {
	fmt.Printf("id=%d,title=%s,author=%s,subject=%s\n", book.Id, book.Title, book.Author, book.Subject)
	book.Id = 3434
}

func RefTag(b Books, idx int) {
	bType := reflect.TypeOf(b)
	ixField := bType.Field(idx)
	fmt.Printf("%v\n", ixField)
}

func InitTechBook() {
	bk := NewBook(1011, "Java", "Tom", "this is an java")
	tb := new(techBook)
	tb.Cat = "tech"
	tb.int = 10
	tb.Books = *bk
	fmt.Printf("techBook cat = %s\n", tb.Cat)
	fmt.Printf("techBook int = %d\n", tb.int)
	fmt.Printf("techBook Books = %v\n", tb.Books.String())
}

//func main() {
//	var book1 *Books
//	book1 = new(Books)
//	book1.Id = 1
//	book1.Title = "Golang"
//	book1.Author = "Rob Pike"
//	book1.Subject = "about golang"
//	fmt.Println(book1)
//	fmt.Println("bbb")
//	fmt.Println(book1.String())
//	book2 := Books{
//		Id: 1002,
//		Title: "python",
//		Author: "Jack",
//		Subject: "about python",
//
//	}
//	fmt.Println(book2)
//
//	printBook(&book2)
//	fmt.Println(book2)
//	fmt.Println("aaa")
//	fmt.Println(book2.String())
//	fmt.Println(book2.GetTitle())
//
//
//	book3 := NewBook(1004,"Java","hsl","java in action")
//	fmt.Println(book3.GetTitle())
//}
