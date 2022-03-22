package resolvers

import (
	"github.com/aso779/go-ddd-example/infrastructure/graph"
	"github.com/aso779/go-ddd-example/infrastructure/services"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type Resolver struct {
	metaContainer metadata.EntityMetaContainer
	services      services.ServiceContainer
}

func NewResolver(
	metaContainer metadata.EntityMetaContainer,
	services services.ServiceContainer,
) *Resolver {
	return &Resolver{
		metaContainer: metaContainer,
		services:      services,
	}
}

//Mutation constructor for main struct for GRAPHQL mutations
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{
		r,
	}
}

//Query Constructor for main struct for GRAPHQL queries
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

//Main struct for GRAPHQL mutations
type mutationResolver struct {
	*Resolver
}

// Main struct for GRAPHQL queries
type queryResolver struct {
	*Resolver
}
