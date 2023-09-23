package manager

import "github.com/StephanieAgatha/Soraa-Go/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
	UserCredRepo() repository.UserCredential
}

type repoManager struct {
	im InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepo {
	//return constructor from repository user
	return repository.NewUserRepo(r.im.Connect())
}

func (r *repoManager) UserCredRepo() repository.UserCredential {
	return repository.NewUserCredentials(r.im.Connect())
}

func NewRepoManager(im InfraManager) RepoManager {
	return &repoManager{
		im: im,
	}
}
