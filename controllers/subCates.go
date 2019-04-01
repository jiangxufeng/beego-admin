package controllers

import (
	"myblog/models"
	"strconv"
	"strings"
)

type SubCateController struct {
	BaseController
}


func (s * SubCateController) Put() {
	sid, err := s.GetInt("sid")
	if err != nil {
		s.ajaxMsg(80001, "", "invalid params", nil)
	}
	sub, err1 := models.GetSubCateById(sid)
	if err1 != nil{
		s.ajaxMsg(60404, "", "not found the SubCategory", nil)
	}
	if err := s.ParseForm(sub); err != nil{
		s.ajaxMsg(70001, nil, err.Error(), nil)
	}

	sub.Update()
	s.ajaxMsg(0, sub, "success to update the SubCategory", nil)
}

func (s *SubCateController) Delete() {
	ids := strings.Split(s.GetString("sid"), ",")

	for _, v := range ids {
		sid, _ := strconv.Atoi(v)
		sub, err := models.GetSubCateById(sid)
		if err != nil {
			s.ajaxMsg(60404, "", "not found the SubCategory", nil)
		}
		sub.Delete()
	}
	s.ajaxMsg(0, "", "success to delete the SubCategory.", nil)
}

func (s *SubCateController) Post() {
	var sub models.SubCategory
	if err1 := s.ParseForm(&sub); err1 != nil{
		s.ajaxMsg(70001, nil, err1.Error(), nil)
	}
	sid, err := models.SubCateSave(&sub)
	if err != nil {
		s.ajaxMsg(90001, "", "failed to save the SubCategory.", nil)
	}
	s.ajaxMsg(0, sid, "success to save the SubCategory.", nil)
}

func SubCateSearch(name string) ([]map[string]interface{}, int){
	var list []*models.SubCategory
	new(models.SubCategory).Query().Filter("Name__icontains", name).All(&list)
	var results []map[string]interface{}
	for _, v := range list{
		temp := make(map[string]interface{})
		temp["Id"] = v.Id
		temp["Name"] = v.Name
		temp["Created"] = v.Created
		main, _ := models.GetMainCateById(v.Father)
		temp["Father"] = main.Name
		temp["PassageNums"], _ = new(models.Passage).Query().Filter("Subcategory", v.Id).Count()
		results = append(results, temp)
	}
	return results, len(results)

}