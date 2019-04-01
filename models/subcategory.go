package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SubCategory struct {
	// ID
	Id int `json:"id"`
	// 子目录名
	Name string `orm:"unique;index;size(8)" json:"name" form:"name"`
	// 创建时间
	Created time.Time `orm:"auto_now_add;type(datetime);" json:"created"`
	// 所属父目录
	Father int `json:"father" form:"main"`

}

func (s *SubCategory) TableName() string{
	return TableName("sub_category")
}

func (s *SubCategory) Insert() error{
	if _, err := orm.NewOrm().Insert(s); err != nil{
		return err
	}
	return nil
}

func (s *SubCategory) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

func (s *SubCategory) Delete() error {
	if _, err := orm.NewOrm().Delete(s); err != nil{
		return err
	}
	return nil
}

// 查询数据
func (s *SubCategory) Read(fields ...string) error{
	if err := orm.NewOrm().Read(s, fields...); err != nil {
		return err
	}
	return nil
}

func (s *SubCategory) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(s)
}

func SubCateList()([]map[string]interface{}, int64) {
	list := make([]*SubCategory, 0)
	total, _ := new(SubCategory).Query().Count()
	new(SubCategory).Query().OrderBy("-Id").All(&list)
	var results []map[string]interface{}
	for _, v := range list {
		temp := make(map[string]interface{})
		temp["Id"] = v.Id
		temp["Name"] = v.Name
		temp["Created"] = v.Created
		main, _ := GetMainCateById(v.Father)
		temp["Father"] = main.Name
		temp["PassageNums"], _ = new(Passage).Query().Filter("Subcategory", v.Id).Count()
		results = append(results, temp)
	}
	return results, total
}

func SubCateSave(m *SubCategory) (int64, error){
	id, err := orm.NewOrm().Insert(m)
	if err == nil{
		return id, nil
	}
	return 0, err
}

func GetSubCateById(id int) (*SubCategory, error){
	m := new(SubCategory)
	err := m.Query().Filter("Id", id).One(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetSubCateByName(name string) (*SubCategory, error){
	m := new(SubCategory)
	err := m.Query().Filter("name", name).One(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}


