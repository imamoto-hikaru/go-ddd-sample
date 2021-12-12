package user

import (
	user_app "advent-calendar/application/user"
	"advent-calendar/domain/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserCreateController struct {
	userCreateService user_app.UserCreateService
}

func NewUserCreateController(userService user_app.UserCreateService) *UserCreateController {
	c := new(UserCreateController)
	c.userCreateService = userService
	return c
}

func (controller *UserCreateController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.PostForm("user_name")
		if len(userName) == 0 {
			c.Error(errors.New("user_name is required"))
			return
		}

		user, err := controller.userCreateService.Create(user.Name(userName))
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"user": gin.H{
				"id":   user.Id,
				"name": user.Name,
			},
		})
	}
}
