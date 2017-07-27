package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"encoding/json"
"gopkg.in/mgo.v2"

)


func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

var db *mgo.Database

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	db = session.DB("straightothecode")
	defer session.Close()


	schema, err := graphql.Schema{(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType
	})}
	http.Handle("/graphql", handler(schema))

	http.ListenAndServe(":8080", nil)
}
