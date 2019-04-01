package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// type User models.User


type Passage struct {
	// ID
	Id int `json:"id"`
	// 文章名
	Title string `orm:"unique;size(32);index" json:"title"`
	// 头图链接
	HeadImg string `orm:"size(256);default()" json:"head_img"`
	// 内容
	Content string `orm:"type(text)" json:"content"`
	// 标签
	Tags string `orm:"size(100);" json:"tags"`
	// 子目录
	SubCategory int  `json:"sub_category"`
	// 主目录
	MainCategory int ` json:"main_category"`
	// 发布时间
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	// 浏览次数
	ViewNums int `orm:"default(0);" json:"view_nums"`
	// 点赞次数
	LikeNums int `orm:"default(0);" json:"like_nums"`
	// 是否是推荐文章
	IsPush bool `orm:"default(false);" json:"is_push"`
	// 是否发布
	IsPublish bool `orm:"default(false);" json:"is_publish"`
	// 所属用户
	User string `json:"user"`
	// 历史
	// History *History `orm:"rel(one);column(文章历史)"`
}

func (p *Passage) TableName() string{
	return TableName("passage")
}

func (p *Passage) Query() orm.QuerySeter{
	return orm.NewOrm().QueryTable(p)
}

func (p *Passage) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func (p *Passage) Delete() error{
	if _, err := orm.NewOrm().Delete(p); err != nil{
		return err
	}
	return nil
}


func PassageSave(p *Passage) (int64, error){
	id, err := orm.NewOrm().Insert(p)
	if err == nil{
		return id, nil
 	}
	return 0, err
}

func PassageGetById(id int) (*Passage, error){
	p := new(Passage)
	err := p.Query().Filter("Id", id).One(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 暂时没有分页和筛选
func PassageList()([]*Passage, int64){
	list := make([]*Passage, 0)
	total, _ := new(Passage).Query().Count()
	new(Passage).Query().OrderBy("-Id").All(&list)
	return list, total
}

// 根据主目录名获取文章
func GetPassagesByMainCate(id int) ([]*Passage, int64) {
	passages := make([]*Passage, 0)
	query := new(Passage).Query().Filter("MainCategory", id)
	total, _ := query.Count()
	query.OrderBy("-Id").All(&passages)
	return passages, total
}

// 根据子目录名获取文章
func GetPassagesBySubCate(id int) ([]*Passage, int64) {
	passages := make([]*Passage, 0)
	query := new(Passage).Query().Filter("SubCategory", id)
	total, _ := query.Count()
	query.OrderBy("-Id").All(&passages)
	return passages, total
}

func GetOtherPassages() map[string][]*Passage {
	// 获得推荐文章
	var recommends []*Passage
	new(Passage).Query().Filter("IsPush", true,).Filter("IsPublish", true).OrderBy("-Created").All(&recommends)

	// 获得点击数排行的文章
	var mostView []*Passage
	new(Passage).Query().Filter("IsPublish",true).OrderBy("-ViewNums").All(&mostView)

	// 获得轮播的文章
	var wheel []*Passage
	new(Passage).Query().Filter("IsPublish",true).All(&wheel)

	// fmt.Println(*(recommends[0].MainCategory))
	return map[string][]*Passage{
		"recommend_top": recommends[:1],
		"recommend_down": recommends[1:4],
		"most_view_top": mostView[:1],
		"most_view_down": mostView[1:5],
		"wheel": wheel[:3],
		"rightTop": wheel[4:6],
	}
}

func getPreAndNextPassages(passage Passage) map[string]Passage {
	var objList []Passage
	var pre, next Passage
	new(Passage).Query().Filter("IsPublish",true).OrderBy("Created").All(&objList)
    for i, v := range objList {
    	if v.Id == passage.Id {
    		if i != 0 {
    			pre = objList[i-1]
			} else {
				pre = Passage{}
			}
    		if i != len(objList) - 1 {
    			next = objList[i+1]
			} else {
				next = Passage{}
			}
    		break
		}
	}
    return map[string]Passage{
    	"pre": pre,
    	"next": next,
	}
}

func GetRelatedPassage(sub SubCategory, id int) (results []*Passage){
	new(Passage).Query().Exclude("Id", id).Filter("SubCategory", sub.Id).Filter("IsPublish", true).All(&results)
    return
}


