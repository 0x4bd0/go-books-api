package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Book struct {
	Id  string `json:"id"`
	Name string `json:"name"`
	Author  string `json:"author"`
	BookType string `json:"bookType"`
	Price  int `json:"price"`
}

type bookHandler struct {
	sync.Mutex
	store map[string]Book;
}


func (b *bookHandler) books(writer http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET" :
			b.get(writer, request)
			return
	case "POST" :
			b.get(writer, request)
			return
	default :
	        writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Action not alowed"))
			return
	}
}

func (b *bookHandler) post(writer http.ResponseWriter, request *http.Request){

}

func (b *bookHandler) get(writer http.ResponseWriter, request *http.Request){

	books := make([]Book, len(b.store))

	b.Lock()
	i := 0
	for _, book := range b.store {
		books[i] = book
		i++
	}
	b.Unlock()

	jsonBytes, err  := json.Marshal(books)
	
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	    writer.Write([]byte (err.Error()))
	}

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
}

func newBookHandler() *bookHandler {
	return &bookHandler{
		store : map[string]Book{
			"id1" : Book{
					Id  : "id1",
					Name :"Harry Potter and the Philosopher's Stone",
					Author : "J. K. Rowling",
					BookType : "fantasy",
					Price  : 19,
			},
		},
	}
}		

func main () {
	bookHandler := newBookHandler()

	http.HandleFunc("/books",bookHandler.books)

	err := http.ListenAndServe(":1234",nil)

	if(err != nil){
		panic(err)
	}

}