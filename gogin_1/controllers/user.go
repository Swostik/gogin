package controllers

import (
	"zenniz/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user services.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	user.PersonalCode = uuid.New().String()
	newUser, err := c.userService.Create(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	ctx.JSON(200, newUser)

}
func (c *UserController) GetUser(ctx *gin.Context) {
	code := ctx.Param("personalCode")
	user, err := c.userService.Get(code)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "uer not found",
		})
		return
	}
	ctx.JSON(200, user)

}
func (c *UserController) UpdateUser(ctx *gin.Context) {
	code := ctx.Param("personalCode")
	var user services.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	for i, u := range services.Users {
		if u.PersonalCode == code {
			services.Users[i].FirstName = user.FirstName
			services.Users[i].LastName = user.LastName
			//services.Users[i].PersonalCode = user.PersonalCode
			ctx.JSON(200, gin.H{
				"error": false,
			})
			return

		}
	}
	updatedUser, err := c.userService.Update(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	ctx.JSON(200, updatedUser)

}
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("personalCode")

	for i, u := range services.Users {
		if u.PersonalCode == id {
			services.Users = append(services.Users[:i], services.Users[i+1:]...)
			ctx.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	if err := c.userService.Delete(id); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request",
		})
		return

	}

}
func (c *UserController) List(ctx *gin.Context) {
	ctx.JSON(200, services.Users)

}
