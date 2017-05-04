package controllers

import (
	"dkvgo-admin/forms"
	"dkvgo-admin/models"
	"dkvgo-admin/services"
	"dkvgo-admin/utils"

	"github.com/astaxie/beego/validation"
)

type UsersController struct {
	BaseController
}

func (this *UsersController) Get() {
	page, err := this.GetInt("page", 1)
	this.CheckError(err)
	pageSize, err := this.GetInt("size", 10)
	this.CheckError(err)
	var users []*models.User
	qs := services.UserService.GetUserList(page, pageSize)
	field := this.GetString("field")
	keyword := this.GetString("keyword")
	if field != "" && keyword != "" {
		if field == "Username" || field == "Email" {
			qs = qs.Filter(field+"__contains", keyword)
		}
	}
	_, err = qs.OrderBy("-UpdateAt").All(&users)
	this.CheckError(err)
	pager, err := services.UserService.GetPage(page, pageSize)
	this.DataJsonResponseWithPage(users, pager)
}

func (this *UsersController) Post() {
	loginUser := this.LoginUser()
	if !loginUser.IsAdmin() {
		this.ShowErrorMsg("没有权限")
	}
	form := forms.UserCreateForm{
		Username: this.GetString("Username"),
		Email:    this.GetString("Email"),
		Password: this.GetString("Password"),
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		this.CheckError(err)
	}
	if !ok {
		this.ErrorJsonResponse(utils.ErrorsPack(valid.Errors), nil)
	}
	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}
	err = services.UserService.Create(&user)
	if err != nil {
		this.CheckError(err)
	}
	this.DataJsonResponse(user, "user")
}
