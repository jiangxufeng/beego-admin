package controllers

import (
	"myblog/models"
	"strconv"
	"time"
)

type AdminController struct {
	BaseController
}

func (a *AdminController) Index() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "index.html"
}

func (a *AdminController) Welcome() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.Data["time"] = time.Now()
	a.TplName = "welcome.html"
}

func (a *AdminController) MemberList() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	if q := a.GetString("q"); q == "search" {
		start := a.GetString("start", "")
		end  := a.GetString("end", "")
		username := a.GetString("username", "")
		userList, total := Search(start, end, username, false)
		a.Data["userlist"] = userList
		a.Data["total"] = total
	} else {
		userList, total := models.UserList(true)
		a.Data["userlist"] = userList
		a.Data["total"] = total
	}
	a.TplName = "member-list.html"
}

func (a *AdminController) MemberDel() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	if q := a.GetString("q"); q == "search" {
		start := a.GetString("start", "")
		end  := a.GetString("end", "")
		username := a.GetString("username", "")
		userList, total := Search(start, end, username, true)
		a.Data["userlist"] = userList
		a.Data["total"] = total
	} else {
		userList, total := models.UserList(false)
		a.Data["userlist"] = userList
		a.Data["total"] = total
	}
	a.TplName = "member-del.html"
}

func (a *AdminController) MainList() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "main-list.html"
}

func (a *AdminController) TagList() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "cate.html"
}

func (a *AdminController) PassageList() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "passage-list.html"
}

func (a *AdminController) PassageDel() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "passage-del.html"
}

func (a *AdminController) SubList() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.Data["username"] = user.Username
	a.TplName = "sub-list.html"
}

func (a *AdminController) UserCreate() {
	user, err := a.Authentication()
	if err != nil || user.IsAdmin == false {
		a.Abort("403")
	}
	a.TplName = "member-add.html"
}

func (a *AdminController) ChangePassword() {
	admin, err := a.Authentication()
	if err != nil || admin.IsAdmin == false {
		a.Abort("403")
	}

	param := a.Ctx.Input.Params()["0"]
	uid, _ := strconv.Atoi(param)
	user, err := models.GetUserById(uid)
	a.Data["user"] = user
	a.TplName = "member-password.html"
}

func (a *AdminController) ChangeUserInfo() {
	admin, err := a.Authentication()
	if err != nil || admin.IsAdmin == false {
		a.Abort("403")
	}

	param := a.Ctx.Input.Params()["0"]
	uid, _ := strconv.Atoi(param)

	user, err := models.GetUserById(uid)
	a.Data["user"] = user
	a.TplName = "member-edit.html"
}
