package graph

import "job-portal-api/graphql/graph/model"

type Storer interface {
	func (s *Service) AddCompany(company *model.Company) (*model.Company, error)
    func (s *Service) AllCompanies() []*model.Company
    func (s *Service) FindCompanyByID(companyID string) (*model.Company, error)
    func (s *Service) AllJobs() []*model.Job
    func (s *Service) FindJobByID(jobID string) (*model.Job, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}
