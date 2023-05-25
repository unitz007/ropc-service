package services

//func Test_OptimizedAuthentication(t *testing.T) {
//
//	mockUserRepository := new(mocks.UserRepository)
//
//	mockUserRepository.On("GetUserAndClient", wrongTestUser).Return(nil, errors.New("invalid user credentials"))
//
//	optimizedAuthenticator := NewOptimizedAuthenticator(mockUserRepository)
//
//	token, err := optimizedAuthenticator.Authenticate(&wrongTestUser, &wrongClient)
//	assert.EqualError(t, err, "invalid credentials")
//	assert.Nil(t, token)
//	mockUserRepository.AssertCalled(t, "GetUserClient", wrongTestUser)
//
//}
