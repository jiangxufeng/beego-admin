package controllers

import (
	"myblog/models"
	"encoding/json"
	"time"
)

type PassageController struct {
	BaseController
}

//func MainCateName(Id int) string{
//	main := models.MainCategory{Id: Id}
//	if err := main.Read(); err != nil{
//		return err.Error()
//	} else {
//		return main.Name
//	}
//}
//
//func SubCateName(Id int) string{
//	main := models.SubCategory{Id: Id}
//	if err := main.Read(); err != nil{
//		return err.Error()
//	} else {
//		return main.Name
//	}
//}
//
//func UserName(Id int) string{
//	user, err := models.GetUserById(Id)
//	if err != nil{
//		return err.Error()
//	}
//	return user.Username
//}

type PassageJson struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	HeadImg      string    `json:"headImg"`
	Content      string    `json:"content"`
	SubCategory  string    `json:"sub_category"`
	MainCategory string    `json:"main_category"`
	User         string    `json:"user"`
	ViewNums     int       `json:"view_nums"`
	LikeNums     int       `json:"like_nums"`
	Tags         []string  `json:"tags"`
	Created      time.Time `json:"created"`
	IsPush       bool      `json:"is_push"`
	IsPublish    bool      `json:"is_publish"`
}

// @Title GetAllPassages
// @Description Get All Passages From The Database
// @Param	main	query 	string	false	"The name of maincategory that you want to get"
// @Param	sub	    query 	string	false	"The name of subcategory that you want to get"
// @Success 200 {object} Passage
// @Failure 400 Bad Request
// @router / [get]
func (p *PassageController) GetAll(){
	var (
		passages []*models.Passage
		total int64
	)
	main, _ := p.GetInt("main", 0)
	sub, _ := p.GetInt("sub", 0)
	if sub != 0 {
		passages, total = models.GetPassagesBySubCate(sub)
	} else if main != 0 {
		passages, total = models.GetPassagesByMainCate(main)
	} else {
		passages, total = models.PassageList()
	}
	results := make(map[string]interface{})
	results["error"] = 0
	results["message"] = "success to get all passages"
    //var temp []*models.Passage
	//for _, v := range passages{
	//	temp = append(temp, v)
	//}
	results["data"] = passages
	results["total"] = total
	p.Data["json"] = results
	p.ServeJSON()
}

// @Title Detail
// @Description Get The Detail Of A Passage
// @Param	pid		path 	string	true	"The pid of the passage that you want to get"
// @Success 200 {error:0, data:Passage, message:"success to get the passage"}
// @Failure 404 Not Found
// @router /:pid [get]
func (p *PassageController) Detail() {
	pid, err := p.GetInt(":pid")
	if err != nil {
		p.ajaxMsg(80001, "", "invalid params", nil)
	}
	passage, err := models.PassageGetById(pid)
	if err != nil{
		p.ajaxMsg(40404, "", "not found the passage", nil)
	}
	p.ajaxMsg(0, passage, "success to get the passage.", nil)

}

// @Title Update
// @Description update a passage
// @Param	pid		path 	string	true	"The pid of the passage that you want to get"
// @Param   body    body    models.Passage true "The field that you want to update"
// @Success 200 {error:0, data:Passage, message:"success to update the passage"}
// @Failure 404 Not Found
// @Failure 400 Bad Request
// @router /:pid [put]
func (p *PassageController) Update() {
	pid, err := p.GetInt(":pid")
	if err != nil {
		p.ajaxMsg(80001, "", "invalid params", nil)
	}
	passage, err1 := models.PassageGetById(pid)
	if err1 != nil{
		p.ajaxMsg(40404, "", "not found the passage", nil)
	}
	json.Unmarshal(p.Ctx.Input.RequestBody, &passage)
	passage.Update()
	p.ajaxMsg(0, passage, "success to update the passage", nil)
}
//

// @Title Delete
// @Description delete a passage
// @Param	pid		path 	string	true	"The pid of the passage that you want to get"
// @Success 200 {error:0, data:Passage, message:"success to delete the passage"}
// @Failure 404 Not Found
// @router /:pid [delete]
func (p *PassageController) Delete() {
	pid, err := p.GetInt(":pid")
	if err != nil {
		p.ajaxMsg(80001, "", "invalid params", nil)
	}

	passage, err := models.PassageGetById(pid)
	if err != nil{
		p.ajaxMsg(40404, "", "not found the passage", nil)
	}
	passage.Delete()
	p.ajaxMsg(0, passage, "success to get the passage.", nil)

}

// @Title CreatePassage
// @Description create Passage
// @Param	body		body 	models.Passage	true		"body for passage content"
// @Success 200 {error:0, data:Passage, message:"success to delete the passage"}
// @Failure 400 body is empty
// @router / [post]
func (p *PassageController) Create() {
	var passage models.Passage
	json.Unmarshal(p.Ctx.Input.RequestBody, &passage)
	pid, err := models.PassageSave(&passage)
	if err != nil {
		p.ajaxMsg(90001, "", "failed to save the passage.", nil)
	}
	p.ajaxMsg(0, pid, "success to save the passage.", nil)
}

