package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MainCategory struct {
	// ID
	Id int `json:"id"`
	// 主目录名
	Name string `orm:"unique;index;size(8)" json:"name"`
	// 创建时间
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
}

func (m *MainCategory) TableName() string{
	return TableName("main_category")
}

func (m *MainCategory) Insert() error{
	if _, err := orm.NewOrm().Insert(m); err != nil{
		return err
	}
	return nil
}

func (m *MainCategory) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MainCategory) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil{
		return err
	}
	return nil
}

// 查询数据
func (m *MainCategory) Read(fields ...string) error{
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MainCategory) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func GetSons(Id int) (sons []*SubCategory){
	new(SubCategory).Query().Filter("father", Id).All(&sons)
	return
}

func MainCateList()([]*MainCategory, int64){
	list := make([]*MainCategory, 0)
	total, _ := new(MainCategory).Query().Count()
	new(MainCategory).Query().OrderBy("-Id").All(&list)
	return list, total
}

func MainCateSave(m *MainCategory) (int64, error){
	id, err := orm.NewOrm().Insert(m)
	if err == nil{
		return id, nil
	}
	return 0, err
}

func MainCateGetById(id int) (*MainCategory, error){
	m := new(MainCategory)
	err := m.Query().Filter("Id", id).One(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MainCateGetByName(name string) (*MainCategory, error){
	m := new(MainCategory)
	err := m.Query().Filter("name", name).One(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

