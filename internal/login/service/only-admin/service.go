package only_admin

import "users-api-server/internal/login/service"

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() service.ILoginService {
	return loginInformation{
		email:    "admin",
		password: "123",
	}
}
func (info loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
