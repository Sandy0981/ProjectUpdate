package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"job-portal-api/internal/auth"
	middlewares "job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"job-portal-api/internal/models/mockmodels"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_ViewCompanyAll(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockCompany := []models.Company{{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		CompanyName: "TEKsystem",
		FoundedYear: "2016",
		Location:    "USA",
	}}
	// Define the list of test cases
	testCases := []struct {
		name              string                          // Name of the test case
		expectedStatus    int                             // Expected status of the response
		expectedResponse  string                          // Expected response body
		expectedInventory []models.Company                // Expected user after signup
		mockService       func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `{"companies list":[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"company_name":"TEKsystem","founded_year":"2016","location":"USA"}]}`,
			// Function for mocking service.
			// This simulates CreateInventory service and its return value.
			mockService: func(m *mockmodels.MockService) {

				//var f float64 = 60
				m.EXPECT().ViewCompanyAll(gomock.Any(), gomock.Any()).Times(1).
					Return(mockCompany, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/viewcompanyall", h.ViewCompanyAll)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/viewcompanyall", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_ViewJobAll(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// MockJob struct initialization
	mockJobs := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:           "Software Engineer",
			ExperienceLevel: "Senior",
			CompanyID:       1,
		},
		// Add more mock jobs as needed.
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `{"job list":[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","experience_required":"Senior","company_id":1}]}`,
			// Function for mocking service.
			// This simulates ViewJobAll service and its return value.
			mockService: func(m *mockmodels.MockService) {

				m.EXPECT().ViewJobAll(gomock.Any(), gomock.Any()).Times(1).
					Return(mockJobs, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/viewjoball", h.ViewJobAll)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/viewjoball", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_ViewJobByCompId(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// MockJob struct initialization
	mockJobs := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:           "Software Engineer",
			ExperienceLevel: "Senior",
			CompanyID:       1,
		},
		// Add more mock jobs as needed.
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","experience_required":"Senior","company_id":1}]`,
			// Function for mocking service.
			// This simulates ViewJobByCompId service and its return value.
			mockService: func(m *mockmodels.MockService) {

				m.EXPECT().ViewJobByCompId(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(mockJobs, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/viewcompany/:companyID", h.ViewJobByCompId)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/viewcompany/1", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_ViewJobByJobId(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// MockJob struct initialization
	mockJobs := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:           "Software Engineer",
			ExperienceLevel: "Senior",
			CompanyID:       1,
		},
		// Add more mock jobs as needed.
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","experience_required":"Senior","company_id":1}]`,
			// Function for mocking service.
			// This simulates ViewJobByJobId service and its return value.
			mockService: func(m *mockmodels.MockService) {

				m.EXPECT().ViewJobByJobId(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(mockJobs, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/viewjobbyid/:jobID/jobs", h.ViewJobByJobId)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/viewjobbyid/1/jobs", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_ViewCompany(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// MockCompany struct initialization
	mockCompany := models.Company{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		CompanyName: "TEKsystem",
		FoundedYear: "2016",
		Location:    "USA",
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"company_name":"TEKsystem","founded_year":"2016","location":"USA"}`,
			// Function for mocking service.
			// This simulates ViewCompany service and its return value.
			mockService: func(m *mockmodels.MockService) {
				m.EXPECT().ViewCompany(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(mockCompany, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/viewcompany/:companyID", h.ViewCompany)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/viewcompany/1", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_CreateCompany(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// Define the input data for creating a company
	companyData := models.Company{
		CompanyName: "NewCompany",
		FoundedYear: "2022",
		Location:    "USA",
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:           "OK",
			expectedStatus: 200,
			// You can adjust the expected response based on your application's actual response format.
			expectedResponse: `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"NewCompany","founded_year":"2022","location":"USA"}`,
			// Function for mocking service.
			// This simulates CreateCompany service and its return value.
			mockService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateCompany(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(companyData, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/addcompany", h.CreateCompany)

			// Serialize companyData to JSON and create a request body
			reqBody, _ := json.Marshal(companyData)

			// Create a new HTTP POST request to "/createcompany".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/addcompany", bytes.NewReader(reqBody))
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_CreateJob(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// Define the input data for creating a job
	jobData := models.Job{
		Title:           "Software Engineer",
		ExperienceLevel: "Senior",
		CompanyID:       1,
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                          // Name of the test case
		expectedStatus   int                             // Expected status of the response
		expectedResponse string                          // Expected response body
		mockService      func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:           "OK",
			expectedStatus: 201,
			// You can adjust the expected response based on your application's actual response format.
			expectedResponse: `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"Software Engineer","experience_required":"Senior","company_id":1}`,
			// Function for mocking service.
			// This simulates CreateJob service and its return value.
			mockService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateJob(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(jobData, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.

			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/createjob/:companyID/jobs", h.CreateJob)

			// Serialize jobData to JSON and create a request body
			reqBody, _ := json.Marshal(jobData)

			// Create a new HTTP POST request to "/createjob".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/createjob/1/jobs", bytes.NewReader(reqBody))
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
