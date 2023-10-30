package store

import "graphql/graph/model"

type Storer interface {
	CreateCompany(company *model.Company) (*model.Company, error)
	CreateJob(job *model.Job) (*model.Job, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}
