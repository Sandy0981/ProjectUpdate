package services

import (
	"context"
	"job-portal-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels

type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
		error)
	CreateCompany(ctx context.Context, nu models.NewCompany, companyId int) (models.Company, error)
	CreateJob(ctx context.Context, nu models.Job, UserId string) (models.Job, error)
	ViewCompanyAll(ctx context.Context, companyId string) ([]models.Company, error)
	ViewJobAll(ctx context.Context, companyId string) ([]models.Job, error)
	ViewJob(ctx context.Context, companyId string) ([]models.Job, error)
	ViewCompany(ctx context.Context, companyID uint, userId string) (models.Company, error)
	ViewJobByCompId(ctx context.Context, companyID uint, userId string) ([]models.Job, error)
	ViewJobByJobId(ctx context.Context, jobById uint, userId string) ([]models.Job, error)
	AutoMigrate() error
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
