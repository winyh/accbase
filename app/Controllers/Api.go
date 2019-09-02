package Controllers

import (
	"accbase/app/Models"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Admins struct {
	Models.Admins
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

func UserInfo(c *gin.Context)  {
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

func AdminUserInfo(c *gin.Context)  {
	var json  Models.Admins
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	admin, err := json.FindOne(id)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": admin,
		"message":"查询成功",
	})
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
	var msg string
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := json.Destroy(id)

	if err != nil{
		msg = "删除失败"
	} else {
		msg = "删除成功"
	}

	c.JSON(200, gin.H{
		"status": true,
		"message": msg,
	})
}

func AdminUserUpdate(c *gin.Context) {
	var json  Models.Admins
	var msg string
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	c.BindJSON(&json)

	err := json.Update(id)

	if err != nil{
		msg = "更新失败"
	} else {
		msg = "更新成功"
	}

	if err != nil {
		fmt.Printf("database error %v", err)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"message": msg,
	})
}

func AdminUserFindAll(c *gin.Context)  {
	var json  Models.Admins

	admin, err := json.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": admin,
		"message":"查询成功",
	})
}
