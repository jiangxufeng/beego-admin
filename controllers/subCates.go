package controllers

import (
	"encoding/json"
	"myblog/models"
)

type SubCateController struct {
	BaseController
}

// @Title Get All SubCategories
// @Description Get All SubCategories From The Database
// @Success 200 {object} models.SubCategory
// @Failure 400 Bad Request
// @router / [get]
func (s *SubCateController) GetAll(){
	subCates, total := models.SubCateList()
	results := make(map[string]interface{})
	results["error"] = 0
	results["message"] = "success to get all passages"
	var temp []*models.SubCategory
	for _, v := range subCates{
		temp = append(temp, v)
	}
	results["data"] = temp
	results["total"] = total
	s.Data["json"] = results
	s.ServeJSON()
}
//
// @Title Detail
// @Description Get The Detail Of A SubCategory
// @Param	sid		path 	string	true	"The mid of the SubCategory that you want to get"
// @Success 200 {error:0, data:SubCategory, message:"success to get the SubCategory"}
// @Failure 404 Not Found
// @router /:sid [get]
func (s *SubCateController) Detail() {
	sid, err := s.GetInt(":sid")
	if err != nil {
		s.ajaxMsg(80001, "", "invalid params", nil)
	}

	subCate, err := models.SubCateGetById(sid)
	if err != nil{
		s.ajaxMsg(60404, "", "not found the SubCategory.", nil)
	}
	s.ajaxMsg(0, subCate, "success to get the SubCategory.", nil)

}

// @Title Update
// @Description update a SubCategory
// @Param	sid		path 	string	true	"The pid of the SubCategory that you want to get"
// @Param   body    body    models.SubCategory true "The field that you want to update"
// @Success 200 {error:0, data:SubCategory, message:"success to update the SubCategory"}
// @Failure 404 Not Found
// @Failure 400 Bad Request
// @router /:sid [put]
func (s * SubCateController) Update() {
	sid, err := s.GetInt(":sid")
	if err != nil {
		s.ajaxMsg(80001, "", "invalid params", nil)
	}
	sub, err1 := models.SubCateGetById(sid)
	if err1 != nil{
		s.ajaxMsg(60404, "", "not found the SubCategory", nil)
	}
	passages, _ := models.GetPassagesBySubCate(sub.Name)
	json.Unmarshal(s.Ctx.Input.RequestBody, &sub)
	for _, v := range passages{
		v.SubCategory = sub.Name
		v.Update()
	}
	sub.Update()
	s.ajaxMsg(0, sub, "success to update the SubCategory", nil)
}
//

// @Title Delete
// @Description delete a SubCategory
// @Param	sid		path 	string	true	"The pid of the SubCategory that you want to get"
// @Success 200 {error:0, data:"", message:"success to delete the passage"}
// @Failure 404 Not Found
// @router /:sid [delete]
func (s *SubCateController) Delete() {
	sid, err := s.GetInt(":sid")
	if err != nil {
		s.ajaxMsg(80001, "", "invalid params", nil)
	}

	sub, err := models.SubCateGetById(sid)
	if err != nil{
		s.ajaxMsg(60404, "", "not found the SubCategory", nil)
	}
	passages, _ := models.GetPassagesBySubCate(sub.Name)
	for _, v := range passages{
		v.SubCategory = ""
		v.Update()
	}
	sub.Delete()
	s.ajaxMsg(0, "", "success to get the SubCategory.", nil)

}

// @Title CreateSubCategory
// @Description create SubCategory
// @Param	body		body 	models.SubCategory	true		"body for SubCategory content"
// @Success 200 {error:0, data:SubCategory, message:"success to create the SubCategory"}
// @Failure 400 body is empty
// @router / [post]
func (s *SubCateController) Create() {
	var sub models.SubCategory
	json.Unmarshal(s.Ctx.Input.RequestBody, &sub)
	sid, err := models.SubCateSave(&sub)
	if err != nil {
		s.ajaxMsg(90001, "", "failed to save the SubCategory.", nil)
	}
	s.ajaxMsg(0, sid, "success to save the SubCategory.", nil)
}