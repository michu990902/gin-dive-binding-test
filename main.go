package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/test", Test)

	r.Run()
}

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

type Group struct {
	Name string `json:"name"`
	// Users []User `json:"users" binding:"dive"`
	Users []User `json:"users"`
}

func Test(c *gin.Context){
	var err error
	var data Group

	err = c.BindJSON(&data)

	c.JSON(200, gin.H{
		"message": "test",
		"data": data,
		"err": err,
	})
}

/* test:
header:
Content-Type: application/json

body:
{
    "name": "Group1",
    "users": [
        {"name": "User1", "age": 1},
        {"name": "User2", "age": 2},
        {"name": "User3", "age": 3}
    ]
}
*/