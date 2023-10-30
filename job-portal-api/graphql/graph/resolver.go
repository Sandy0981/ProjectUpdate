package graph

import "job-portal-api/graphql/graph/store"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	S store.Store
}
