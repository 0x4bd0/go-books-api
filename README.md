# go-books-api

> A simple REST API built to try Go lang

## Requirements

* No third-party packages/dependencies

## Endpoints

* [x] `GET /books` returns list of books as JSON
* [x] `GET /books/{id}` returns details of specific book as JSON
* [x] `POST /books` to add book as JSON
* [x] `POST /books` returns status 415 if content is not `application/json`
* [x] `PUT /books/{id}` to update a  book
* [x] `DELETE /books/{id}` to delete a  book


### Data Types

A book object should look like this:
```json
{
  "id": "someid",
  "name": "name of the book",
  "author": "author of the book",
  "bookType": "the type of the book",
  "price": 99,
}
```

### Where is the data stored ?

There is no database, using the machine memory.


### What do i think about Go lang ?

This lang rocks !

