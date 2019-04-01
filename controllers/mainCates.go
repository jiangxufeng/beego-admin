package controllers

import (
	"myblog/models"
	"strconv"
	"strings"
	"time"
)

type MainCateController struct {
	BaseController
}

type MainReturnJson struct {
	Name string `json:"name"`
	Id int `json:"id"`
	Created time.Time `json:"created"`
	Sons []*models.SubCategory `json:"sons"`
}


func (m * MainCateController) Put() {
	mid, err := m.GetInt("mid")
	if err != nil {
		m.ajaxMsg(80001, "", "invalid params", nil)
	}
	main, err1 := models.GetMainCateById(mid)
	if err1 != nil{
		m.ajaxMsg(40404, "", "not found the maincategory", nil)
	}
	if err := m.ParseForm(main); err != nil{
		m.ajaxMsg(70001, nil, err.Error(), nil)
	}
	main.Update()
	m.ajaxMsg(0, main, "success to update the maincategory", nil)
}

// 删除
func (m *MainCateController) Delete() {
	ids := strings.Split(m.GetString("mid"), ",")

	for _, v := range ids {
		mid, _ := strconv.Atoi(v)
		main, err := models.GetMainCateById(mid)
		if err != nil{
			m.ajaxMsg(40404, "", "not found the maincategory", nil)
		}
		sons, _ := models.GetSons(mid)
		for _,v := range sons{
			v.Delete()
		}

		passages, _ := models.GetPassagesByMainCate(mid)
		for _, v := range passages{
			v.Delete()
		}

	    main.Delete()
	}
	m.ajaxMsg(0, "", "success to delete the MainCategory.", nil)

}

func (m *MainCateController) Post() {
	var main models.MainCategory
	if err1 := m.ParseForm(&main); err1 != nil{
		m.ajaxMsg(70001, nil, err1.Error(), nil)
	}
	pid, err := models.MainCateSave(&main)
	if err != nil {
		m.ajaxMsg(90001, "", "failed to save the maincategory.", nil)
	}
	m.ajaxMsg(0, pid, "success to save the maincategory.", nil)
}

func MainCateSearch(name string) ([]map[string]interface{}, int){
	var list []*models.MainCategory

	new(models.MainCategory).Query().Filter("Name__icontains", name).All(&list)
	var results []map[string]interface{}
	for _, v := range list{
		temp := make(map[string]interface{})
		_, sonnum := models.GetSons(v.Id)
		temp["id"] = v.Id
		temp["name"] = v.Name
		temp["Created"] = v.Created
		temp["sonnum"] = sonnum
		results = append(results, temp)
	}
	return results, len(results)

}
