package handlers

import (
	"fmt"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"net/http"
)

// Define a function called API that takes an argument a of type *auth.Auth
// and returns a pointer to a gin.Engine

func API(a *auth.Auth, c *models.Conn) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	r := gin.New()

	// Attempt to create new middleware with authentication
	// Here, *auth.Auth passed as a parameter will be used to set up the middleware
	m, err := middleware.NewMid(a)
	ms := services.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}

	// If there is an error in setting up the middleware, panic and stop the application
	// then log the error message
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// Attach middleware's Log function and Gin's Recovery middleware to our application
	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
	r.Use(m.Log(), gin.Recovery())

	// Define a route at path "/check"
	// If it receives a GET request, it will use the m.Authenticate(check) function.
	r.GET("/check", m.Authenticate(check))
	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/addcompany", m.Authenticate(h.CreateCompany))
	r.GET("/viewcompanyall", m.Authenticate(h.ViewCompanyAll))
	r.GET("/viewcompany/:companyID", m.Authenticate(h.ViewCompany))
	r.POST("/createjob/:companyID/jobs", m.Authenticate(h.CreateJob))
	r.GET("/viewjob/:companyID/jobs", m.Authenticate(h.ViewJobByCompId))
	r.GET("/viewjobbyid/:jobID/jobs", m.Authenticate(h.ViewJobByJobId))
	r.GET("/viewjoball", m.Authenticate(h.ViewJobAll))

	// Return the prepared Gin engine
	return r
}

func check(c *gin.Context) {
	//handle panic using recovery function when happening in separate goroutine
	//go func() {
	//	panic("some kind of panic")
	//}()
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
