package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our tutorial array for the matching id
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							// return our tutorial
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"tutorial_db": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial By ID with SQL",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					db, err := sql.Open("sqlite3", "./tutorial.db")
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					var tutorial Tutorial
					err = db.QueryRow("SELECT ID, Title FROM tutorials where ID = ?", id).Scan(&tutorial.ID, &tutorial.Title)
					if err != nil {
						log.Fatal(err)
					}
					return tutorial, nil
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
		"list_db": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List with SQL",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := sql.Open("sqlite3", "./tutorial.db")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()

				// perform a db.Query insert
				var tutorials []Tutorial
				results, err := db.Query("SELECT * FROM tutorials")
				if err != nil {
					fmt.Println(err)
				}
				for results.Next() {
					var tutorial Tutorial
					err = results.Scan(&tutorial.ID, &tutorial.Title)
					if err != nil {
						fmt.Println(err)
					}
					log.Println(tutorial)
					tutorials = append(tutorials, tutorial)
				}
				return tutorials, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	//schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// define mutation schema
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	// Query
	query := `
   {
        list_db {
            id
            title
        }
		tutorial_db(id:2){
	id
title
}
    }
`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
