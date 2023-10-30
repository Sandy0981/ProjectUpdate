// resolver.go

package graph

import (
	"context"
	"job-portal-api/graphql/graph/model"
)

// CreateCompany is the resolver for the createCompany field.
func (r *queryResolver) AddCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	company := &model.Company{
		CompanyName: input.CompanyName,
		FoundedYear: input.FoundedYear,
		Location:    input.Location,
		// Assuming there's no association with User in this context.
	}

	return r.S.AddCompany(company)
}

// AllCompanies is the resolver for the allCompanies field.
func (r *queryResolver) AllCompanies(ctx context.Context) ([]*model.Company, error) {
	return r.S.AllCompanies(), nil
}

// FindCompanyByID is the resolver for the findCompanyById field.
func (r *queryResolver) FindCompanyByID(ctx context.Context, companyID string) (*model.Company, error) {
	return r.S.FindCompanyByID(companyID)
}

// jobResolver struct for implementing JobResolver interface
type jobResolver struct{ *Resolver }

// CreateJob is the resolver for the createJob field.
func (r *queryResolver) CreateJob(ctx context.Context, input model.Job) (*model.Job, error) {
	job := &model.Job{
		ID:              input.ID,
		Title:           input.Title,
		ExperienceLevel: input.ExperienceLevel,
		CompanyID:       input.CompanyID,
	}

	return r.S.AddJob(job)
}

// AllJobs is the resolver for the allJobs field.
func (r *queryResolver) AllJobs(ctx context.Context) ([]*model.Job, error) {
	return r.S.AllJobs(), nil
}

// FindJobByID is the resolver for the findJobById field.
func (r *queryResolver) FindJobByID(ctx context.Context, jobID string) (*model.Job, error) {
	return r.S.FindJobByID(jobID)
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
