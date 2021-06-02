package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name string `grom:"type:varchar(20);not null"`
	Telephone string `grom:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
			name = RandomString(10)
		}
		log.Println(name, password, telephone)

		//判断手机号是否存在
		if isTelePhoneExist(db,telephone) {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已经存在"})
			return
		}

		//创建用户
		newUser := User{
			Name:name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


//生成自定义字符
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnm")
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}


//数据库手机号检验
func isTelePhoneExist(db *gorm.DB,telephone string) bool {
	var user User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//初始化数据库 gorm
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "wanghanyong0718"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	log.Println(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failecd to connect detabase, err: "+err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}