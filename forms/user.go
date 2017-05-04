package forms

import (
	"github.com/astaxie/beego/validation"

	"dkvgo-admin/services"
)

type UserCreateForm struct {
	Username string `valid:"Required;"`
	Email    string `valid:"Required;Email"`
	Password string `valid:"Required;MinSize(6);MaxSize(20)"`
}

func (f *UserCreateForm) Valid(v *validation.Validation) {
	user, _ := services.UserService.GetUserByUsername(f.Username)
	if user != nil {
		v.SetError("Username", "用户名已经被占用")
	}
	user, _ = services.UserService.GetUserByEmail(f.Email)
	if user != nil {
		v.SetError("Email", "邮箱已经被占用")
	}
}
