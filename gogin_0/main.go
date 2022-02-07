package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PersonalCode string `json:"personalCode"`
}

var Users []User

func main() {
	r := gin.Default()
	r.GET("/users", GetAll)
	r.POST("/users", CreateUser)
	r.PUT("/users/:personalCode", EditUser)
	r.DELETE("/users/:personalCode", DeleteUSer)
	//fmt.Println("hello")
	r.Run()

}

func GetAll(c *gin.Context) {

	c.JSON(200, Users)

}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}
	reqBody.PersonalCode = uuid.New().String()
	Users = append(Users, reqBody)

	c.JSON(200, Users)

	//fmt.Println(reqBody)

}

func EditUser(c *gin.Context) {
	id := c.Param("personalCode")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}

	for i, u := range Users {
		if u.PersonalCode == id {
			Users[i].FirstName = reqBody.FirstName
			Users[i].LastName = reqBody.LastName

			c.JSON(200, gin.H{
				"error": false,
			})
			return

		}
	}
	c.JSON(400, gin.H{
		"error":   true,
		"message": "invalid personalCode",
	})

}

func DeleteUSer(c *gin.Context) {
	id := c.Param("personalCode")

	for i, u := range Users {
		if u.PersonalCode == id {
			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid code",
	})

}
