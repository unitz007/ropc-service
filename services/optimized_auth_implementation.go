package services

import (
	"ropc-service/repositories"
)

type OptimizedAuthenticator struct {
	userRepository repositories.UserRepository
}

func NewOptimizedAuthenticator(userRepository repositories.UserRepository) *OptimizedAuthenticator {
	return &OptimizedAuthenticator{
		userRepository: userRepository,
	}
}

//func (selfC OptimizedAuthenticator) Authenticate(user *model.User, client *model.Client) (*model.Token, error) {
//
//	_, err := selfC.userRepository.GetUserAndClient(user, client)
//
//	return nil, err
//}
