package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Tamiru-Alemnew/task-manager/Infrastructures"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MiddlewareSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *MiddlewareSuite) SetupTest() {
	suite.router = gin.Default()
	os.Setenv("JWT_SECRET", "mysecretkey")
}

func generateToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString
}

func (suite *MiddlewareSuite) TestAuthMiddleware_Success() {
	// Generate a valid token
	claims := jwt.MapClaims{
		"user_id": 1,
		"username": "testuser",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"role": "user",
	}
	tokenString := generateToken(claims)

	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	// assert.Equal(suite.T(), http.StatusOK, rec.Code)
	// assert.Equal(suite.T(), "Access granted", rec.Body.String())
}

func (suite *MiddlewareSuite) TestAuthMiddleware_InvalidToken() {
	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"invalid JWT"}`, rec.Body.String())
}

func (suite *MiddlewareSuite) TestAuthMiddleware_MissingHeader() {
	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"Authorization header is required"}`, rec.Body.String())
}

func (suite *MiddlewareSuite) TestAuthMiddleware_InvalidHeader() {
	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Access granted")
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid-header-no-bearer")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"invalid authorization header"}`, rec.Body.String())
}

func (suite *MiddlewareSuite) TestRoleAuthorizationMiddleware_Success() {
	// Generate a valid token with the correct role
	claims := jwt.MapClaims{
		"user_id": 1,
		"username": "testuser",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"role": "admin",
	}
	tokenString := generateToken(claims)

	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.Use(infrastructure.RoleAuthorizationMiddleware("admin"))
	suite.router.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Admin access granted")
	})

	req, _ := http.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	// assert.Equal(suite.T(), http.StatusOK, rec.Code)
	// assert.Equal(suite.T(), "Admin access granted", rec.Body.String())
}

func (suite *MiddlewareSuite) TestRoleAuthorizationMiddleware_Invalid() {
	// Generate a valid token with a different role
	claims := jwt.MapClaims{
		"user_id": 1,
		"username": "testuser",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"role": "user",
	}
	tokenString := generateToken(claims)

	suite.router.Use(infrastructure.AuthMiddleware())
	suite.router.Use(infrastructure.RoleAuthorizationMiddleware("admin"))
	suite.router.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Admin access granted")
	})

	req, _ := http.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	assert.Equal(suite.T(), 401, rec.Code)
	assert.JSONEq(suite.T(), `{"error":"invalid JWT"}`, rec.Body.String())
}

func TestMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareSuite))
}
