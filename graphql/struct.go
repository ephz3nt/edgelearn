package main

import (
	"github.com/graphql-go/graphql"
)

var tutorials = populate()

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

func populate() []Tutorial {
	author := &Author{
		Name:      "Elliot Forbes",
		Tutorials: []int{1},
	}
	tutorial := Tutorial{
		//ID:     1,
		Title:  "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}
	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	return tutorials
}

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				// we'll use NewList to deal with an array
				// of int values
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)
var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				// here, we specify type as ahotherType
				// which we've already defined.
				// This is how we handle nested objects
				Type: authorType,
			},
			"comments": &graphql.Field{

				Type: graphql.NewList(commentType),
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        tutorialType,
				Description: "Create a new Tutorial",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tutorial := Tutorial{
						Title: params.Args["title"].(string),
					}
					tutorials = append(tutorials, tutorial)
					return tutorial, nil
				},
			},
		},
	},
)
