package mstore

import (
	"fmt"
	"graphql/graph/model"
	"strconv"
)

// Declaring a struct named Service that includes a map that takes a string as key and pointer to Company as value
type Service struct {
	companyStore map[string]*model.Company
	jobStore     map[string]*model.Job
}

// Initializing the Service struct by creating empty maps for companyStore and jobStore
func NewService() Service {
	return Service{
		companyStore: make(map[string]*model.Company),
		jobStore:     make(map[string]*model.Job),
	}
}

// Method to add a company to the companyStore. Accepts a pointer to a company, performs a dummy database operation, then adds the company to the companyStore
func (s *Service) CreateCompany(company *model.Company) (*model.Company, error) {
	fmt.Println("doing heavy db operation specific stuff") // Displaying a message to signify performing database manipulation
	s.companyStore[strconv.Itoa(company.ID)] = company     // Adding the company to the map under the key of company.ID
	return company, nil                                    // Returning the added company and nil (indicating no error)
}

// Method to return all companies in the companyStore

// Method to find and return a company by its ID, or an error if i

// Method to add a job to the jobStore. Accepts a pointer to a job, performs a dummy database operation, then adds the job to the jobStore
func (s *Service) CreateJob(job *model.Job) (*model.Job, error) {
	fmt.Println("doing heavy db operation specific stuff") // Displaying a message to signify performing database manipulation
	s.jobStore[strconv.Itoa(job.ID)] = job                 // Adding the job to the map under the key of job.ID
	return job, nil                                        // Returning the added job and nil (indicating no error)
}
