package controllers

import (
	"myblog/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type BaseController struct {
	beego.Controller
}


func (b *BaseController) ajaxMsg(code int, data interface{}, msg interface{}, other map[string]interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["message"] = msg
	if other != nil{
		for k, v := range other{
			out[k] = v
		}
	}
	b.Data["json"] = out
	b.ServeJSON()
	b.StopRun()
}

func (b *BaseController) Authentication() (user *models.User, err error) {
	username := b.GetSession("username")
	if username != nil {
		if user, err := models.GetUserByName(username.(string)); err != nil{
			return nil, err
		} else {
			return user, nil
		}
	}
	return nil, errors.New("no session")
}

// token验证
func (b *BaseController) ParseToken() (jwt.MapClaims, error){
	tokenString := b.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", tokenString)

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("suummmmer"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, InputError
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, TokenExpiredError
			} else {
				// Couldn't handle this token
				return nil, InputError
			}
		} else {
			// Couldn't handle this token
			return nil, InputError
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil, TokenInvalidError
	}
	beego.Debug("Token:", token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, TokenInvalidError
	}
	return claims, nil

}


func fatal(err error) {
	if err != nil {
		beego.Error(err)
	}
}

// 登录
func Login(username, password string) (int, error){
	user := models.User{
		Username: username,
	}
	if err := orm.NewOrm().Read(&user,"username"); err == nil {
		if user.Password == models.Md5([]byte(password)){
			return user.Id, nil
		} else {
            return 0, PasswordError
		}
	} else {
		return 0, UserDoesNotExistError
	}
}

// 生成Token
func CreateToken(username string, uid int) (string, error){
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["username"] = username
	if username == "suummmmer"{
	    claims["admin"] = "admin"
	} else{
		claims["admin"] = "noadmin"
	}
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix() //20天有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("suummmmer"))
	if err != nil {
		return err.Error(), err
	}
	return tokenString, nil
}

// 分页
func Paginator(page, num, total int) (paginatorMap map[string]interface{}) {
	return
}


// 错误处理
// 输入错误，不合法
var InputError = errors.New("invalid token format")

// 密码错误
var PasswordError = errors.New("password is wrong")

// 用户不存在
var UserDoesNotExistError = errors.New("the username does not exist")

// token过期
var TokenExpiredError = errors.New("the token has expired")

// token无效
var TokenInvalidError = errors.New("the token is invalid")


