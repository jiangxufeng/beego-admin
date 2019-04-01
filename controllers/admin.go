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
	username, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.Data["username"] = username
	a.TplName = "index.html"
}

func (a *AdminController) Welcome() {
	username, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.Data["username"] = username
	a.Data["time"] = time.Now()
	a.TplName = "welcome.html"
}

func (a *AdminController) MemberList() {
	_, admin := a.Authentication()
	if !admin {
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
	_, admin := a.Authentication()
	if !admin {
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
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}

	if q := a.GetString("q"); q == "search" {
		name := a.GetString("name", "")
		mainList, total := MainCateSearch(name)
		a.Data["mainlist"] = mainList
		a.Data["total"] = total
	} else {
		mainList, total := models.MainCateList()
		a.Data["mainlist"] = mainList
		a.Data["total"] = total
	}
	a.TplName = "main-list.html"
}

func (a *AdminController) TagList() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.TplName = "cate.html"
}

func (a *AdminController) PassageList() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.TplName = "passage-list.html"
}

func (a *AdminController) PassageDel() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.TplName = "passage-del.html"
}

func (a *AdminController) SubList() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	if q := a.GetString("q"); q == "search" {
		name := a.GetString("name", "")
		subList, total := SubCateSearch(name)
		a.Data["sublist"] = subList
		a.Data["total"] = total
	} else {
		subList, total := models.SubCateList()
		a.Data["sublist"] = subList
		a.Data["total"] = total
	}
	a.TplName = "sub-list.html"
}

func (a *AdminController) UserCreate() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.TplName = "member-add.html"
}

func (a *AdminController) ChangePassword() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}

	param := a.Ctx.Input.Params()["0"]
	uid, _ := strconv.Atoi(param)
	user, _ := models.GetUserById(uid)
	a.Data["user"] = user
	a.TplName = "member-password.html"
}

func (a *AdminController) ChangeUserInfo() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}

	param := a.Ctx.Input.Params()["0"]
	uid, _ := strconv.Atoi(param)

	user, _ := models.GetUserById(uid)
	a.Data["user"] = user
	a.TplName = "member-edit.html"
}

func (a *AdminController) ChangeMainCateInfo() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	param := a.Ctx.Input.Params()["0"]
	mid, _ := strconv.Atoi(param)

	main, _ := models.GetMainCateById(mid)
	a.Data["main"] = main
	a.TplName = "main-edit.html"
}

func (a *AdminController) MainCateCreate() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	a.TplName = "main-add.html"
}

func (a *AdminController) SubCateCreate() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}

	mainList, _ := models.MainCateList()
	a.Data["mainList"] = mainList
	a.TplName = "sub-add.html"
}

func (a *AdminController) ChangeSubCateInfo() {
	_, admin := a.Authentication()
	if !admin {
		a.Abort("403")
	}
	param := a.Ctx.Input.Params()["0"]
	sid, _ := strconv.Atoi(param)

	sub, _ := models.GetSubCateById(sid)
	mainList, _ := models.MainCateList()
	a.Data["sub"] = sub
	a.Data["mainList"] = mainList
	a.TplName = "sub-edit.html"
}