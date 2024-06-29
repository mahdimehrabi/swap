package service

import (
	entity "bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	mock_logger "bbdk/mock/infrastructure"
	mock_user "bbdk/mock/repository"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		user         *entity.User
		ctx          context.Context
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: nil,
		},
		{
			name: "UserRepoError",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(err)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			err := service.CreateUser(test.user)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}
			loggerMock.EXPECT()
			userRepoMock.EXPECT()
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	user := entity.User{Name: "John Doe", Email: "john@example.com"}
	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		id           uint
		result       *entity.User
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(&user, nil)
				return userRepoMock
			},
			id:     1,
			result: &user,
			error:  nil,
		},
		{
			name: "not found",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(nil, userRepo.ErrNotFound)
				return userRepoMock
			},
			id:     2,
			result: nil,
			error:  userRepo.ErrNotFound,
		},
		{
			name: "repo error",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(nil, err)
				return userRepoMock
			},
			id:     3,
			result: nil,
			error:  err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			result, err := service.GetUserByID(test.id)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}

			if !gomock.Eq(result).Matches(test.result) {
				t.Errorf("result:%v is not equal to %v", result, test.result)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		user         *entity.User
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().UpdateUser(gomock.Any()).Return(nil)
				return userRepoMock
			},
			user:  &entity.User{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"},
			error: nil,
		},
		{
			name: "not found",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().UpdateUser(gomock.Any()).Return(userRepo.ErrNotFound)
				return userRepoMock
			},
			user:  &entity.User{ID: 2, Name: "John Doe", Email: "john@example.com", Password: "password123"},
			error: userRepo.ErrNotFound,
		},
		{
			name: "repo error",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().UpdateUser(gomock.Any()).Return(err)
				return userRepoMock
			},
			user:  &entity.User{ID: 3, Name: "John Doe", Email: "john@example.com", Password: "password123"},
			error: err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			err := service.UpdateUser(test.user)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		id           uint
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().DeleteUser(gomock.Any()).Return(nil)
				return userRepoMock
			},
			id:    1,
			error: nil,
		},
		{
			name: "not found",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().DeleteUser(gomock.Any()).Return(userRepo.ErrNotFound)
				return userRepoMock
			},
			id:    2,
			error: userRepo.ErrNotFound,
		},
		{
			name: "repo error",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().DeleteUser(gomock.Any()).Return(err)
				return userRepoMock
			},
			id:    3,
			error: err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			err := service.DeleteUser(test.id)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	users := []*entity.User{{Name: "John Doe", Email: "john@example.com"}}
	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		id           uint
		result       []*entity.User
		error        error
		page         int
		pageSize     int
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(users, nil)
				return userRepoMock
			},
			id:       1,
			result:   users,
			error:    nil,
			pageSize: 10,
			page:     2,
		},
		{
			name: "invalid page",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				return userRepoMock
			},
			id:       2,
			result:   nil,
			error:    ErrInvalidPage,
			page:     -2,
			pageSize: 0,
		},
		{
			name: "repo error",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, err)
				return userRepoMock
			},
			id:       3,
			result:   nil,
			error:    err,
			pageSize: 10,
			page:     2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			result, err := service.GetAllUsers(test.page, test.pageSize)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}

			if !gomock.Eq(result).Matches(test.result) {
				t.Errorf("result:%v is not equal to %v", result, test.result)
			}
		})
	}
}

func BenchmarkUserService_CreateUser(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().CreateUser(gomock.Any()).Times(b.N).Return(nil).AnyTimes()
	loggerMock := mock_logger.NewMockLogger(ctrl)
	service := NewUserService(userRepoMock, loggerMock)
	user := &entity.User{Name: "John Doe", Email: "john@example.com", Password: "password123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.CreateUser(user)
	}
}

func BenchmarkUserService_DeleteUser(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().DeleteUser(gomock.Any()).Return(nil).AnyTimes()
	loggerMock := mock_logger.NewMockLogger(ctrl)
	service := NewUserService(userRepoMock, loggerMock)
	id := uint(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.DeleteUser(id)
	}
}

func BenchmarkUserService_UpdateUser(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().UpdateUser(gomock.Any()).Return(nil).AnyTimes()
	loggerMock := mock_logger.NewMockLogger(ctrl)
	service := NewUserService(userRepoMock, loggerMock)
	user := &entity.User{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.UpdateUser(user)
	}
}

func BenchmarkUserService_GetUserByID(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().GetUserByID(gomock.Any()).Return(&entity.User{Name: "John Doe", Email: "john@example.com"}, nil).AnyTimes()
	loggerMock := mock_logger.NewMockLogger(ctrl)
	service := NewUserService(userRepoMock, loggerMock)
	id := uint(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUserByID(id)
	}
}

func BenchmarkUserService_GetAllUsers(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	loggerMock := mock_logger.NewMockLogger(ctrl)
	service := NewUserService(userRepoMock, loggerMock)

	// Prepare mock data
	users := make([]*entity.User, 1000)
	for i := 0; i < 1000; i++ {
		users[i] = &entity.User{Name: fmt.Sprintf("User%d", i+1)}
	}

	// Define benchmark parameters
	page := 1
	pageSize := 10
	offset := (page - 1) * pageSize
	limit := pageSize

	// Mock repository call
	userRepoMock.EXPECT().GetAll(offset, limit).Return(users, nil).AnyTimes()

	b.ResetTimer()

	// Run benchmark
	for i := 0; i < b.N; i++ {
		_, err := service.GetAllUsers(page, pageSize)
		if err != nil {
			b.Fatalf("Failed to fetch users: %v", err)
		}
	}
}
