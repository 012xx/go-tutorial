package graph


			// This file will be automatically regenerated based on the schema, any resolver implementations
			// will be copied through when generating and any unknown code will be moved to the end.
			// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
"context"
"fmt"
"io"
"strconv"
"time"
"sync"
"errors"
"bytes"
gqlparser "github.com/vektah/gqlparser/v2"
"github.com/vektah/gqlparser/v2/ast"
"github.com/99designs/gqlgen/graphql"
"github.com/99designs/gqlgen/graphql/introspection"
"my_gql_server/graph/model"
"my_gql_server/internal")


















// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
	func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) ( *model.AddProjectV2ItemByIDPayload,  error){
		panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
	}

// Repository is the resolver for the repository field.
	func (r *queryResolver) Repository(ctx context.Context, name string, owner string) ( *model.Repository,  error){
		panic(fmt.Errorf("not implemented: Repository - repository"))
	}

// !!! WARNING !!!
The code below was going to be deleted when updating resolvers. It has been copied here so you have
one last chance to move it out of harms way if you want. There are two reasons this happens:
  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
    it when you're done.
  - You have helper methods in this file. Move them out to keep these resolver files clean.
	func (r *queryResolver) User(ctx context.Context, name string) ( *model.User,  error){
		return r.Srv.GetUserByName(ctx, name)
	}

// Node is the resolver for the node field.
	func (r *queryResolver) Node(ctx context.Context, id string) ( model.Node,  error){
		panic(fmt.Errorf("not implemented: Node - node"))
	}



// Mutation returns internal.MutationResolver implementation.
	func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }
// Query returns internal.QueryResolver implementation.
	func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }


type mutationResolver struct { *Resolver }
type queryResolver struct { *Resolver }



    // !!! WARNING !!!
    // The code below was going to be deleted when updating resolvers. It has been copied here so you have
    // one last chance to move it out of harms way if you want. There are two reasons this happens:
	//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
	//    it when you're done.
	//  - You have helper methods in this file. Move them out to keep these resolver files clean.
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

