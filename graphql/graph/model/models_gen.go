// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	ID          string `json:"ID"`
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
	Jobs        []*Job `json:"jobs"`
}

type Job struct {
	ID                 string `json:"ID"`
	Title              string `json:"title"`
	ExperienceRequired string `json:"experience_required"`
	CompanyID          int    `json:"company_id"`
}

type NewCompany struct {
	CompanyName string `json:"company_name"`
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
}

type NewJob struct {
	Title              string `json:"title"`
	ExperienceRequired string `json:"experience_required"`
	CompanyID          int    `json:"company_id"`
}