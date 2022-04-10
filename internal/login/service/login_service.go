package service

type ILoginService interface {
	LoginUser(email string, password string) bool
}
