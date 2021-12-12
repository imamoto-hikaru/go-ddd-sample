package user

import (
	user_app "advent-calendar/application/user"
	"advent-calendar/domain/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserFindController struct {
	userFindService user_app.UserFindService
}

func NewUserFindController(userService user_app.UserFindService) *UserFindController {
	c := new(UserFindController)
	c.userFindService = userService
	return c
}

func (controller *UserFindController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Param("user_id")
		if len(v) == 0 {
			c.Error(errors.New("user_id is required"))
			return
		}

		id, err := user.CreateIdFrom(v)
		if err != nil {
			c.Error(err)
			return
		}

		user, err := controller.userFindService.Find(id)
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
