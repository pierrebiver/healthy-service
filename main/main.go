package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
)
func main() {

	http.ListenAndServe(":8080", nil)
}