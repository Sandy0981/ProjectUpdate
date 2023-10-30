package store

import (
	"errors"
	"fmt"
	"job-portal-api/graphql/graph/model"
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
func (s *Service) AddCompany(company *model.Company) (*model.Company, error) {
	fmt.Println("doing heavy db operation specific stuff") // Displaying a message to signify performing database manipulation
	s.companyStore[company.ID] = company                   // Adding the company to the map under the key of company.ID
	return company, nil                                    // Returning the added company and nil (indicating no error)
}

// Method to return all companies in the companyStore
func (s *Service) AllCompanies() []*model.Company {
	var companies []*model.Company     // Initialize an empty slice of pointers to companies
	for _, c := range s.companyStore { // Iterate over all entries in the companyStore map
		//c := c
		fmt.Println(c)                   // Display the company
		companies = append(companies, c) // Add the company to the companies slice
	}

	return companies // Return the slice containing all companies
}

// Method to find and return a company by its ID, or an error if it is not found
func (s *Service) FindCompanyByID(companyID string) (*model.Company, error) {
	c, ok := s.companyStore[companyID] // Trying to get the company by the provided companyID
	if !ok {                           // If the company is not found
		return nil, errors.New("company not found") // Return nil and an error saying the company was not found
	}
	return c, nil // Return the company found
}

// Method to add a job to the jobStore. Accepts a pointer to a job, performs a dummy database operation, then adds the job to the jobStore
func (s *Service) AddJob(job *model.Job) (*model.Job, error) {
	fmt.Println("doing heavy db operation specific stuff") // Displaying a message to signify performing database manipulation
	s.jobStore[job.ID] = job                               // Adding the job to the map under the key of job.ID
	return job, nil                                        // Returning the added job and nil (indicating no error)
}

// Method to return all jobs in the jobStore
func (s *Service) AllJobs() []*model.Job {
	var jobs []*model.Job          // Initialize an empty slice of pointers to jobs
	for _, j := range s.jobStore { // Iterate over all entries in the jobStore map
		//j := j
		fmt.Println(j)         // Display the job
		jobs = append(jobs, j) // Add the job to the jobs slice
	}

	return jobs // Return the slice containing all jobs
}

// Method to find and return a job by its ID, or an error if it is not found
func (s *Service) FindJobByID(jobID string) (*model.Job, error) {
	j, ok := s.jobStore[jobID] // Trying to get the job by the provided jobID
	if !ok {                   // If the job is not found
		return nil, errors.New("job not found") // Return nil and an error saying the job was not found
	}
	return j, nil // Return the job found
}
