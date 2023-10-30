// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	ID string       `json:"company_id"`
	CompanyName string `json:"company_name"` //gorm:"unique;not null"
	FoundedYear string `json:"founded_year"`
	Location    string `json:"location"`
	//	UserId      string `json:"user_id"`
	Jobs []Job `json:"jobs,omitempty" gorm:"foreignKey:CompanyID"`
}

type NewCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	FoundedYear string `json:"founded_year" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Jobs        []Job  `json:"jobs"`
}

type Job struct {
	ID string `json:"job_id"`
	Title           string `json:"title"`
	ExperienceLevel string `json:"experience_required"`
	CompanyID       uint   `json:"company_id"`
}
