package Proxy

type UserFinder interface {
	FindUser(id int32) (User, error)
}
