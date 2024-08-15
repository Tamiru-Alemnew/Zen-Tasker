package tests

import (
	"testing"


	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	"github.com/Tamiru-Alemnew/task-manager/Infrastructures"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JwtServiceSuite struct {
	suite.Suite
	service domain.JWTService
	secret  string
}

func (suite *JwtServiceSuite) SetupTest() {
	suite.secret = "test_secret"
	service := infrastructure.NewJWTService(suite.secret)
	suite.service = service
}

// Test GenerateToken

func (suite *JwtServiceSuite) TestGenerateToken_Success() {
	token, err := suite.service.GenerateToken(1 , "testuser", "admin")
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(suite.secret), nil
	})

	assert.NoError(suite.T(), err)
	assert.True(suite.T(), parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "testuser", claims["username"])
	assert.Equal(suite.T(), "admin", claims["role"])
}


func (suite *JwtServiceSuite) ParseToken_Success() {
	token, err := suite.service.GenerateToken(1 , "testUser" , "admin")


	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)

	claims, err := suite.service.ParseToken(token)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "testUser", claims.Username)
	assert.Equal(suite.T(), "admin", claims.Role)
}

func (suite *JwtServiceSuite) ParseToken_InvalidToken() {
	invalidToken := "invalid.token.string"

	_, err := suite.service.ParseToken(invalidToken)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "invalid character '\\u008a' looking for beginning of value", err.Error())
}


func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}