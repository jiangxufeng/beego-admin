package controllers

import (
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
	"strings"
)

// Operations about Users
type UserController struct {
	BaseController
}

func (u *UserController) Create() {
	user := models.User{}
	if err1 := u.ParseForm(&user); err1 != nil{
		u.ajaxMsg(70001, nil, err1.Error(), nil)
	}
	// 检查用户名是否已经存在
	_, err := models.GetUserByName(user.Username)
	if err == nil{
		u.ajaxMsg(20001, nil, "The username exists.", nil)
	}

	if _, err := user.Save(); err != nil{
		u.ajaxMsg(80001, nil, err.Error(), nil)
	}
	u.ajaxMsg(0, user, "success to save the user.", nil)

}

func (u *UserController) Get() {
	claim, tkerr := u.ParseToken()
	if tkerr != nil {
		u.ajaxMsg(20007, "", tkerr.Error(), nil)
	}
	// 判断是否有权限
	if claim["admin"].(string) != "admin"{
		u.ajaxMsg(20007, "", "you don't have the permission.", nil)
	}
	// uid是否有效
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.ajaxMsg(70001, nil, "invalid param", nil)
	}
	user, err := models.GetUserById(uid)
	if err != nil {
		u.ajaxMsg(20004,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     nil, "not found the user", nil)
	} else {
		u.ajaxMsg(0, user, "success to get the user", nil)
	}

}

func (u *UserController) Put() {
	_, admin :=  u.Authentication()
	if !admin {
		u.Abort("403")
	}

	uid, err := u.GetInt(":uid")
	if err != nil {
		u.ajaxMsg(70001, nil, err.Error(), nil)
	}
	user, err := models.GetUserById(uid)
	if err != nil{
		u.ajaxMsg(20004, nil, "not found the user.", nil)
	}

	if err := u.ParseForm(user); err != nil{
		u.ajaxMsg(70001, nil, err.Error(), nil)
	}

	user.Update()
	u.ajaxMsg(0, user, "success to update the user.", nil)
}

func (u *UserController) Delete() {
	_, admin :=  u.Authentication()
	if !admin {
		u.Abort("403")
	}

	uid, err := u.GetInt(":uid")
	if err != nil {
		u.ajaxMsg(70001, nil, err.Error(), nil)
	}
	user, err := models.GetUserById(uid)
	if err != nil{
		u.ajaxMsg(20004, nil, "not found the user.", nil)
	}
	user.Delete()
	u.ajaxMsg(0, user, "success to delete the user.", nil)
}

func (u * UserController) LoginIndex() {
	u.Data["da"] = "da"

}

func (u *UserController) Login() {
	if u.Ctx.Input.IsGet(){
		u.TplName = "login.html"
	} else {
		username := u.GetString("username")
		password := u.GetString("password")
		if isadmin, err := Login(username, password); err == nil {
			u.SetSession("username", username)
			if isadmin{
				u.SetSession("admin", "admin")
			} else {
				u.SetSession("admin", "noadmin")
			}
			u.ajaxMsg(0, "", "success to login", nil)
		} else if err == PasswordError {
			u.ajaxMsg(20003, "", "password is wrong", nil)
		} else {
			u.ajaxMsg(20004, "", "the user does not exist", nil)
		}
	}
}

func (u *UserController) Logout() {
	u.DestroySession()
	u.Redirect("/admin", 302)
}


func (u *UserController) OpAll() {
	_, admin := u.Authentication()
	if !admin {
		u.Abort("403")
	}
    ids := strings.Split(u.GetString("ids"), ";")
    del, _ := u.GetInt("del")
    for _, v := range ids {
	    uid, _ := strconv.Atoi(v)
		user, err := models.GetUserById(uid)
		if err != nil{
			u.ajaxMsg(20004, nil, "not found the user.", nil)
		}
		if del == 1 {
		    user.Del = true
		} else {
		    user.Del = false
		}
		user.Update()
	}

	u.ajaxMsg(0, nil, "success to operate the users.", nil)
}

func Search(start, end, username string, del bool) ([]*models.User, int){
	var users []*models.User
	var query orm.QuerySeter
	query = new(models.User).Query().Filter("Username__icontains", username).Filter("Del", del)

	if start != "" {
		query = query.Filter("JoinTime__gte", start)
	}

	if end != "" {
		query = query.Filter("JoinTime__lte", end)
	}

    query.All(&users)
	total := len(users)
	return users, total
}