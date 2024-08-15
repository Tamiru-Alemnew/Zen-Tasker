package usecases

import (
	"context"
	"errors"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
)

type UserUsecase struct {
	UserRepo        domain.UserRepository
	passwordService domain.PasswordService
	jwtService      domain.JWTService
}

func NewUserUsecase(userRepo domain.UserRepository, passwordService domain.PasswordService, jwtService domain.JWTService) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (uc *UserUsecase) SignUp(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	// Check if the username is already taken
	existingUser, err := uc.UserRepo.FindByUsername(ctx, newUser.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Hash the user's password
	hashedPassword, err := uc.passwordService.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashedPassword

	// Assign role based on the number of existing users
	allUsers, err := uc.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(allUsers) == 0 {
		newUser.Role = "admin"
	} else {
		newUser.Role = "user"
	}

	// Save the new user to the repository
	err = uc.UserRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc *UserUsecase) Login(ctx context.Context, username string, password string) (*domain.User , string, error) {
	// Find the user by username
	user, err := uc.UserRepo.FindByUsername(ctx, username)
	if err != nil {
		return  nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("invalid username or password")
	}

	// Verify the password
	err = uc.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return nil, "", errors.New("invalid username or password")
	}

	
	token, err := uc.jwtService.GenerateToken(user.ID , user.Username, user.Role)
	if err != nil {
		return nil , "", err
	}

	return user , token, nil
}

func (uc *UserUsecase)Promote(ctx context.Context , id int ) error{
	err := uc.UserRepo.Promote(ctx, id)
	return err
}