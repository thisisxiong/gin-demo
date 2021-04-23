package Services

import "gin-demo/Models"

type User struct {
	Id int `json:"id"`
	Username string `json:"username" form:"username" binging:"required"`
	Password string `json:"password" form:"username" binging:"required"`
}

func (this *User) Insert() (id int, err error) {
	var userModel  Models.User
	userModel.Id = this.Id
	userModel.Username = this.Username
	userModel.Password = this.Password
	id, err = userModel.Insert()
	return
}
