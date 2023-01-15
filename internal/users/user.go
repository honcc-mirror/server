package users

type UserId int64
type User struct {
	Id             UserId
	DisplayName    string
	Username       string
	HashedPassword string
}
