package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"why/ginessential/common"
	"why/ginessential/dto"
	"why/ginessential/model"
	"why/ginessential/response"
	"why/ginessential/util"
)

//注册接口控制器
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//没有名称 随机生成字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)

	//判断手机号是否存在
	if isTelePhoneExist(DB,telephone) {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}

	//创建用户
	//加密密码
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		response.Response(ctx,http.StatusUnprocessableEntity,500,nil,"加密错误")
		return
	}
	newUser := model.User{
		Name:name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	response.Success(ctx,nil,"注册成功")
}
//登录接口控制器
func Login(ctx *gin.Context)  {
	DB := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?",telephone).First(&user)
	if user.ID ==0 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		response.Response(ctx,http.StatusUnprocessableEntity,400,nil,"密码错误")
		return
	}

	//发放token
	token,err :=common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx,http.StatusUnprocessableEntity,500,nil,"系统异常")
		log.Printf("token generate error : %v",err)
		return
	}

	//返回结果
	response.Success(ctx,gin.H{"token":token},"登录成功")
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	// 后续dto 使用了类型断言
	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user": dto.ToUserDto(user.(model.
		User))}})
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