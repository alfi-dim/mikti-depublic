package service

type UserService interface {
	LoginUser(email string, password string) (map[string]interface{}, error)
	LoginAdmin(email string, password string) (map[string]interface{}, error)
}
