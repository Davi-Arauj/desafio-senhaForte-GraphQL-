package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/desafio-senhaForte-GraphQL/domain"
	"github.com/desafio-senhaForte-GraphQL/graph/model"
)

// Verify is the resolver for the verify field.
func (r *queryResolver) Verify(ctx context.Context, senhaInput *model.SenhaInput) (*model.Result, error) {
	var (
		senha       model.SenhaInput
		senhaResult model.Result
		err         error
	)

	senha.Password = senhaInput.Password
	senha.Rules = append(senha.Rules, senhaInput.Rules...)

	senhaResult.Verify, senhaResult.NoMatch, err = domain.ValidaSenha(senha)
	if err != nil {
		panic("erro ao validar senha")
	}

	return &senhaResult, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
