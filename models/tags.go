package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	// ID
	Id int `json:"id"`
	// 标签名
	Name string `orm:"unique;index;size(16);column(标签名)" json:"name"`
	// 创建时间
	Created time.Time `orm:"auto_now_add;type(datetime);column(标签创建时间)" json:"created"`

}

// 表名
func (t *Tag) TableName() string{
	return TableName("tags")
}

// 插入数据
func (t *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(t); err != nil {
		return err
	}
	return nil
}

// 更新数据
func (t *Tag) Update(fields ...string) error{
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 查询数据
func (t *Tag) Read(fields ...string) error{
	if err := orm.NewOrm().Read(t, fields...); err != nil {
		return err
	}
	return nil
}

// 删除数据
func (t *Tag) Delete() error{
	if _, err := orm.NewOrm().Delete(t); err != nil{
		return err
	}
	return nil
}

// 查询数据
func (t *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(t)
}
