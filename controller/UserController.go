package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"why/ginessential/common"
	"why/ginessential/model"
	"why/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}
	//没有名称 随机生成字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)

	//判断手机号是否存在
	if isTelePhoneExist(DB,telephone) {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已经存在"})
		return
	}

	//创建用户
	newUser := model.User{
		Name:name,
		Telephone: telephone,
		Password: password,
	}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

//数据库手机号检验
func isTelePhoneExist(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}