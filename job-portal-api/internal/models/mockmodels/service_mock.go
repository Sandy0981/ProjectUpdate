package mockmodels

import (
	"context"

	models "job-portal-api/internal/models"
	reflect "reflect"

	v5 "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockService) Authenticate(ctx context.Context, email, password string) (v5.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, email, password)
	ret0, _ := ret[0].(v5.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockServiceMockRecorder) Authenticate(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockService)(nil).Authenticate), ctx, email, password)
}

// AutoMigrate mocks base method.
func (m *MockService) AutoMigrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoMigrate")
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockServiceMockRecorder) AutoMigrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockService)(nil).AutoMigrate))
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, nu)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, nu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, nu)
}

// Add mock implementations for other methods as needed...

func (m *MockService) CreateCompany(ctx context.Context, ni models.NewCompany, companyId int) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, ni, companyId)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) CreateCompany(ctx, ni, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockService)(nil).CreateCompany), ctx, ni, companyId)
}

func (m *MockService) ViewCompanyAll(ctx context.Context, companyId string) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanyAll", ctx, companyId)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewCompanyAll(ctx, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanyAll", reflect.TypeOf((*MockService)(nil).ViewCompanyAll), ctx, companyId)
}

func (m *MockService) ViewCompany(ctx context.Context, companyById uint, userId string) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompany", ctx, companyById, userId)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewCompany(ctx, companyById, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompany", reflect.TypeOf((*MockService)(nil).ViewCompany), ctx, companyById, userId)
}

func (m *MockService) CreateJob(ctx context.Context, newJob models.Job, userId string) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, newJob, userId)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) CreateJob(ctx, newJob, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockService)(nil).CreateJob), ctx, newJob, userId)
}

func (m *MockService) ViewJobAll(ctx context.Context, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllJob", ctx, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewJobAll(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllJob", reflect.TypeOf((*MockService)(nil).ViewJobAll), ctx, userId)
}
func (m *MockService) ViewJob(ctx context.Context, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllJob", ctx, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewJob(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllJob", reflect.TypeOf((*MockService)(nil).ViewJob), ctx, userId)
}

func (m *MockService) ViewJobByCompId(ctx context.Context, companyId uint, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobByCompId", ctx, companyId, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewJobByCompId(ctx, companyId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobByCompId", reflect.TypeOf((*MockService)(nil).ViewJobByCompId), ctx, companyId, userId)
}

func (m *MockService) ViewJobByJobId(ctx context.Context, jobById uint, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobsByID", ctx, jobById, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewJobByJobId(ctx, jobById, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobsByID", reflect.TypeOf((*MockService)(nil).ViewJobByJobId), ctx, jobById, userId)
}

func (m *MockService) Login(ctx context.Context, email, password string) (v5.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(v5.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), ctx, email, password)
}
