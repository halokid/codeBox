package main 

import (
  "fmt"
)

type Books struct {
  title string
  author string
  subject string
  book_id int
}

func main() {
  var book1 Books
  var book2 Books
  
  book1.title = "go language"
  book1.author = "test"
  book1.subject = "go tech"
  book1.book_id = 65517
  
  book2.title = "php language"
  book2.author = "test2"
  book2.subject = "php tech"
  book2.book_id = 232671
  
  fmt.Println("book1 title: ", book1.title)
  fmt.Println("book1 author: ", book1.author)
  fmt.Println("book1 subject: ", book1.subject)
  fmt.Println("book1 book_id: ", book1.book_id)
  
  
  fmt.Println("book2 title: ", book2.title)
  fmt.Println("book2 author: ", book2.author)
  fmt.Println("book2 subject: ", book2.subject)
  fmt.Println("book2 book_id: ", book2.book_id)
  
  ptrStruct(&book1)
}


func ptrStruct(book *Books) {
  fmt.Println("-----------------------")
  fmt.Println("book title: ", book.title)
  fmt.Println("book author: ", book.author)
  fmt.Println("book subject: ", book.subject)
  fmt.Println("book book_id: ", book.book_id)
  fmt.Println("-----------------------")
}













