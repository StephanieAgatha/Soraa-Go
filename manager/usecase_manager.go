package manager

import "github.com/StephanieAgatha/Soraa-Go/usecase"

type UsecaseManager interface {
	//all usecase object goes here
	UserUsecase() usecase.UserUsecase
	UserCredUsecase() usecase.UserCredentialUsecase
}

type usecaseManager struct {
	rm RepoManager
}

func (u usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.rm.UserRepo())
}

func (u usecaseManager) UserCredUsecase() usecase.UserCredentialUsecase {
	return usecase.NewUserCredentialUsecase(u.rm.UserCredRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		rm: rm,
	}
}
