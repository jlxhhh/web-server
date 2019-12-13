package serializer

import (
	"web-server/model"
)



type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"Password"`
}

func BuildUser(user model.User) User{
	return User{
		ID: user.ID,
		Username:user.Username,
		Password: user.Password,
	}
}
func BuildUserResponse(user model.User) Response{
	return Response{
		Code:  200,
		Data:  BuildUser(user),
		Msg:   "",
		Error: "",
	}
}
