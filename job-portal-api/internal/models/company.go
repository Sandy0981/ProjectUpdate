package models

import (
	"context"
)

// Define the function CreatInventory, which belongs to the struct 'Conn'.
// This function takes in 3 parameters: a context `ctx` of type `Context`, `ni` of type `NewInventory`, and `userId` of type `uint`.
// This function will return an `Inventory` and an `error`.

func (d *Conn) CreateCompany(ctx context.Context, ni NewCompany, userId int) (Company, error) {
	// Create a new 'Inventory' struct named 'inv'.
	// Initialize it with parameters from the 'NewInventory' struct and the `userId` passed to the function.
	cmp := Company{
		CompanyName: ni.CompanyName,
		FoundedYear: ni.FoundedYear,
		Location:    ni.Location,
		//Jobs:        ni.Jobs,
	}
	err := d.db.Create(&cmp).Error
	if err != nil {
		return Company{}, err
	}

	// Successfully created the record, return the user.
	return cmp, nil

}

func (s *Conn) CreateJob(ctx context.Context, job Job, userId string) (Job, error) {
	// Create a new 'Inventory' struct named 'inv'.
	// Initialize it with parameters from the 'NewInventory' struct and the `userId` passed to the function.
	result := s.db.Create(&job)
	if result.Error != nil {
		return Job{}, result.Error
	}
	return job, nil

}

func (s *Conn) ViewCompanyAll(ctx context.Context, companyId string) ([]Company, error) {
	var cmp = make([]Company, 0, 10)
	tx := s.db.Find(&cmp)
	err := tx.Find(&cmp).Error
	if err != nil {
		return nil, err
	}
	return cmp, nil
}

func (s *Conn) ViewCompany(ctx context.Context, companyID uint, UserId string) (Company, error) {
	var company Company
	result := s.db.Where("id = ?", companyID).Find(&company)
	if result.Error != nil {
		return Company{}, result.Error
	}
	return company, nil
}

func (s *Conn) ViewJobByCompId(ctx context.Context, companyID uint, UserId string) ([]Job, error) {
	var job []Job
	result := s.db.Where("company_id = ?", companyID).Find(&job)
	if result.Error != nil {
		return nil, result.Error
	}
	return job, nil
}

func (s *Conn) ViewJobByJobId(ctx context.Context, jobID uint, UserId string) ([]Job, error) {
	var job []Job
	result := s.db.Where("id = ?", jobID).Find(&job)
	if result.Error != nil {
		return nil, result.Error
	}
	return job, nil
}

func (s *Conn) ViewJob(ctx context.Context, userId string) ([]Job, error) {
	var cmp = make([]Job, 0, 10)
	tx := s.db.Find(&cmp)
	err := tx.Find(&cmp).Error
	if err != nil {
		return nil, err
	}
	return cmp, nil
}

func (s *Conn) ViewJobAll(ctx context.Context, companyId string) ([]Job, error) {
	var job = make([]Job, 0, 10)
	tx := s.db.Find(&job)
	err := tx.Find(&job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
