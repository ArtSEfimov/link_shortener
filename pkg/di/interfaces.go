package di

import "http_server/internal/user"

type StatisticRepositoryInterface interface {
	AddClick(uint)
}
type UserRepositoryInterface interface {
	Create(*user.User) (*user.User, error)
	FindByEmail(string) (*user.User, error)
}
