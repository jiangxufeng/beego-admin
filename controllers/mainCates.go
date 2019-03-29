package controllers

import (
	"myblog/models"
	"encoding/json"
	"fmt"
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

// @Title Get All MainCategories
// @Description Get All MainCategories From The Database
// @Success 200 {object} MainCategory
// @Failure 400 Bad Request
// @router / [get]
func (m *MainCateController) GetAll(){
	mainCates, total := models.MainCateList()
	results := make(map[string]interface{})
	results["error"] = 0
	results["message"] = "success to get all passages"
	var temp []*MainReturnJson
	for _, v := range mainCates{
		sons := models.GetSons(v.Id)
		newp := &MainReturnJson{
			Id: v.Id,
			Name: v.Name,
			Created: v.Created,
			Sons: sons,
		}
		fmt.Println(newp)
		temp = append(temp, newp)
	}
	fmt.Println(temp)
	results["data"] = temp
	results["total"] = total
	m.Data["json"] = results
	m.ServeJSON()
}
//
// @Title Detail
// @Description Get The Detail Of A maincategory
// @Param	mid		path 	string	true	"The mid of the maincategory that you want to get"
// @Success 200 {error:0, data:MainCategory, message:"success to get the maincategory"}
// @Failure 404 Not Found
// @router /:mid [get]
func (m *MainCateController) Detail() {
	mid, err := m.GetInt(":mid")
	if err != nil {
		m.ajaxMsg(80001, "", "invalid params", nil)
	}

	mainCate, err := models.MainCateGetById(mid)
	if err != nil{
		m.ajaxMsg(40404, "", "not found the MainCategory.", nil)
	}
	sons := models.GetSons(mid)
	newp := MainReturnJson{
		Id: mid,
		Name: mainCate.Name,
		Created: mainCate.Created,
		Sons: sons,
	}
	m.ajaxMsg(0, newp, "success to get the MainCategory.", nil)

}

// @Title Update
// @Description update a maincategory
// @Param	mid		path 	string	true	"The pid of the maincategory that you want to get"
// @Param   body    body    models.MainCategory true "The field that you want to update"
// @Success 200 {error:0, data:MainCategory, message:"success to update the maincategory"}
// @Failure 404 Not Found
// @Failure 400 Bad Request
// @router /:mid [put]
func (m * MainCateController) Update() {
	mid, err := m.GetInt(":mid")
	if err != nil {
		m.ajaxMsg(80001, "", "invalid params", nil)
	}
	main, err1 := models.MainCateGetById(mid)
	if err1 != nil{
		m.ajaxMsg(40404, "", "not found the maincategory", nil)
	}
	json.Unmarshal(m.Ctx.Input.RequestBody, &main)
	main.Update()
	sons := models.GetSons(mid)
	for _,v := range sons{
		passages, _ := models.GetPassagesBySubCate(v.Name)
		for _, p := range passages{
			p.MainCategory = main.Name
			p.Update()
		}
	}
	m.ajaxMsg(0, main, "success to update the maincategory", nil)
}
//

// @Title Delete
// @Description delete a MainCategory
// @Param	mid		path 	string	true	"The mid of the MainCategory that you want to get"
// @Success 200 {error:0, data:"", message:"success to delete the passage"}
// @Failure 404 Not Found
// @router /:mid [delete]
func (m *MainCateController) Delete() {
	mid, err := m.GetInt(":mid")
	if err != nil {
		m.ajaxMsg(80001, "", "invalid params", nil)
	}

	main, err := models.MainCateGetById(mid)
	if err != nil{
		m.ajaxMsg(40404, "", "not found the passage", nil)
	}
	sons := models.GetSons(mid)
	for _,v := range sons{
		var passages []*models.Passage
		new(models.Passage).Query().Filter("Subcategory", v.Name).All(&passages)
		for _, p := range passages{
			p.MainCategory = ""
			p.Update()
		}
		v.Delete()
	}

	main.Delete()
	m.ajaxMsg(0, "", "success to get the MainCategory.", nil)

}

// @Title CreateMainCategory
// @Description create MainCategory
// @Param	body		body 	models.MainCategory	true		"body for MainCategory content"
// @Success 200 {error:0, data:MainCategory, message:"success to create the MainCategory"}
// @Failure 400 body is empty
// @router / [post]
func (m *MainCateController) Create() {
	var main models.MainCategory
	json.Unmarshal(m.Ctx.Input.RequestBody, &main)
	pid, err := models.MainCateSave(&main)
	if err != nil {
		m.ajaxMsg(90001, "", "failed to save the maincategory.", nil)
	}
	m.ajaxMsg(0, pid, "success to save the maincategory.", nil)
}
