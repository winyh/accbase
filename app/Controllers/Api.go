package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"accbase/app/Models"
	"reflect"
	"time"
)

type Admins struct {
	gorm.Model
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

func Ping(c *gin.Context) {
	e := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	fmt.Println(reflect.TypeOf(e))

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.

	var msg string

	if e.Enforce(sub, obj, act) == true {
		fmt.Println("进入了")
		msg = "进入了"
	} else {
		fmt.Println("拒绝了")
		msg = "无权访问"
	}

	c.JSON(200, gin.H{
		"message": msg,
	})
}

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context)  {
	var json  Params
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}

	hmacSampleSecret := []byte("my_secret_key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": json.Username,
		"password": json.Password, // 21219256@qq.com
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func AdminUserInfo(c *gin.Context)  {
	hmacSampleSecret := []byte("my_secret_key")
	tokenString := c.Request.Header.Get("token")
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-BRTwjN-sAlUjO-82qDrNHdMtGAwgWH05PrN49Ep_sU"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["nbf"])

		c.JSON(200, gin.H{
			"username": claims["username"],
			"password": claims["password"],
		})

	} else {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "获取失败！",
		})
	}
}

func AdminUserCreate(c *gin.Context) {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	id, err := json.Insert()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", id)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"id":id,
		"message":"创建成功",
	})
}

func AdminUserDestroy(c *gin.Context)  {
	var json  Models.Admins
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}

	json.Destroy()

	c.JSON(200, gin.H{
		"status": true,
		"message":"删除成功",
	})
}

func AdminUserUpdate(c *gin.Context) {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	id, err := json.Update(3)

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", id)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"id":id,
		"message":"更新成功",
	})
}

func AdminUserFindAll(c *gin.Context)  {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	result, err := json.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": result,
		"message":"查询成功",
	})
}
