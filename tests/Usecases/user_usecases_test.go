package tests

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	usecases "github.com/Tamiru-Alemnew/task-manager/Usecases"
	"github.com/Tamiru-Alemnew/task-manager/tests/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userUsecaseSuite struct{
	suite.Suite
	userRepository *mocks.UserRepository
	paswordService *mocks.PasswordService
	userUsecase usecases.UserUsecase
	testingServer *httptest.Server
}

func (suite *userUsecaseSuite) SetupSuite(){
	suite.userRepository = new(mocks.UserRepository)
	suite.userUsecase = usecases.UserUsecase{
		UserRepo: suite.userRepository,
	}
}

func(suite *userUsecaseSuite) SetupTest(){
	suite.userRepository = new(mocks.UserRepository)
	suite.userUsecase.UserRepo = suite.userRepository
}

func (suite *userUsecaseSuite) TearDownSuite(){
	defer suite.testingServer.Close()
}

func (suite *userUsecaseSuite) SignUp(){
	user := &domain.User{
		Username : "tamtam",
		Password : "123456",
	}
	suite.userRepository.On("FindByUsername" , mock.Anything , user.Username).Return(nil , nil)
	suite.paswordService.On("HashPassword" , user.Password).Return("hashed password", nil)
	suite.userRepository.On("Create" , mock.Anything , user).Return(nil)

	_ , err := suite.userUsecase.SignUp(context.TODO() , user)

	suite.NoError(err , "no error")
	suite.userRepository.AssertCalled(suite.T() , "FindByUsername" , mock.Anything , user.Username)
	suite.paswordService.AssertCalled(suite.T() , "HashPassword" , user.Password)
	suite.userRepository.AssertCalled(suite.T() , "Create" , mock.Anything , user)

}

func (suite *userUsecaseSuite) SignUp_UsernameTaken(){
	user := &domain.User{
		Username : "tamtam",
		Password : "123456",
	}

	suite.userRepository.On("FindByUsername" , mock.Anything , user.Username).Return(user , nil)
	_ , err := suite.userUsecase.SignUp(context.TODO() , user)

	suite.Error(err , "error")
	suite.userRepository.AssertCalled(suite.T() , "FindByUsername" , mock.Anything , user.Username)
}


func (suite *userUsecaseSuite) Login(){
	user := &domain.User{
		Username : "tamtam",
		Password : "123456",
	}
	suite.userRepository.On("FindByUsername" , mock.Anything , user.Username).Return(user , nil)
	suite.paswordService.On("ComparePassword" , user.Password , user.Password).Return(nil)
	suite.userRepository.On("Promote" , mock.Anything , user.ID).Return(nil)

	_ , _ , err := suite.userUsecase.Login(context.TODO() , user.Username , user.Password)

	suite.NoError(err , "no error")
	suite.userRepository.AssertCalled(suite.T() , "FindByUsername" , mock.Anything , user.Username)
	suite.paswordService.AssertCalled(suite.T() , "ComparePassword" , user.Password , user.Password)
	suite.userRepository.AssertCalled(suite.T() , "Promote" , mock.Anything , user.ID)
}

func (suite *userUsecaseSuite) Login_InvalidUsername(){
	user := &domain.User{
		Username : "tamtam",
		Password : "123456",
	}
	suite.userRepository.On("FindByUsername" , mock.Anything , user.Username).Return(nil , nil)

	_ , _ , err := suite.userUsecase.Login(context.TODO() , user.Username , user.Password)

	suite.Error(err , "error")
	suite.userRepository.AssertCalled(suite.T() , "FindByUsername" , mock.Anything , user.Username)
}

func (suite *userUsecaseSuite) Login_InvalidPassword(){
	user := &domain.User{
		Username : "tamtam",
		Password : "123456",
	}

	suite.userRepository.On("FindByUsername" , mock.Anything , user.Username).Return(user , nil)
	suite.paswordService.On("ComparePassword" , user.Password , user.Password).Return(errors.New("error"))

	_ , _ , err := suite.userUsecase.Login(context.TODO() , user.Username , user.Password)

	suite.Error(err , "error")

	suite.userRepository.AssertCalled(suite.T() , "FindByUsername" , mock.Anything , user.Username)
	suite.paswordService.AssertCalled(suite.T() , "ComparePassword" , user.Password , user.Password)
}




func (suite *userUsecaseSuite) Promote(){
	userId := 1

	suite.userRepository.On("Promote" , mock.Anything , userId).Return(nil)

	err := suite.userUsecase.Promote(context.TODO() , userId)

	suite.NoError(err , "no error")
	suite.userRepository.AssertCalled(suite.T() , "Promote" , mock.Anything , userId)
}



func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}