package users

type User struct {
	Id          int64
	DisplayName string
	Username    string
	Key         string
	Salt        string
}
