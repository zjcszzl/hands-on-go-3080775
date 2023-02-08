// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	author
	title string
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(newBook book) {
	_, ok := l[newBook.author.name]
	if !ok {
		l[newBook.author.name] = make([]book, 0)
	}
	l[newBook.author.name] = append(l[newBook.author.name], newBook)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	l := library{}

	// add 2 books to the library by the same author
	l.addBook(book{author: author{name: "Tom"}, title: "Book 1"})
	l.addBook(book{author: author{name: "Tom"}, title: "Book 2"})

	// add 1 book to the library by a different author
	l.addBook(book{author: author{name: "Jerry"}, title: "Book 3"})

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	books := l.lookupByAuthor("Tom")

	// print out the first book's title and its author's name
	fmt.Println(books[0].title, books[0].author.name)

}
