package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	// id
	Id int `json:"id" form:"-"`
	// 昵称
	Nickname string `orm:"size(32)" json:"nickname" form:"nickname"`
	// 密码
	Password string `orm:"size(64);" json:"password" form:"password"`
	// 邮箱
	Username string `orm:"size(128);unique;" json:"username" form:"username"`
	// 性别
	Sex bool `orm:"default(0);" json:"sex" form:"sex"`
	// 状态
	Status bool `orm:"default(0);" json:"status" form:"status"`
	// 是否是管理员
	IsAdmin bool `orm:"default(0);" json:"is_admin" form:"is_admin"`
	// 加入时间
	JoinTime time.Time `orm:"auto_now_add;type(datetime);" json:"join_time" form:"-"`
	// 是否停用
    Del bool `orm:"default(0);" json:"del" form:"del"`
}

// 表名
func (u *User) TableName() string{
	return TableName("users")
}

// 保存新user
func (u *User) Save() (int64, error) {
	u.Password = Md5([]byte(u.Password))
	return orm.NewOrm().Insert(u)
}

// 更新数据
func (u *User) Update(fields ...string) error{
	user, _ := GetUserById(u.Id)
	if user.Password != u.Password{
		u.Password = Md5([]byte(u.Password))
	}

	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

// 删除数据
func (u *User) Delete() error{
	if _, err := orm.NewOrm().Delete(u); err != nil{
		return err
	}
	return nil
}

// 查询数据
func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

//　获取user
func GetUserById(id int)(*User, error){
	u := new(User)
	err := u.Query().Filter("Id", id).One(u)
	if err != nil{
		return nil, err
	}
	return u, nil
}

func GetUserByName(username string)(*User, error){
	u := new(User)
	err := u.Query().Filter("Username", username).One(u)
	if err != nil{
		return nil, err
	}
	return u, nil
}

// 暂时没有分页和筛选
func UserList(status bool)([]*User, int64){
	list := make([]*User, 0)
	if status {
	    total, _ := new(User).Query().Filter("Del", false).Count()
	    new(User).Query().Filter("Del", false).OrderBy("-JoinTime").All(&list)
		return list, total
	}

	total, _ := new(User).Query().Filter("Del", true).Count()
	new(User).Query().Filter("Del", true).OrderBy("-JoinTime").All(&list)
	return list, total

}