package entry

import (
	entry_app "advent-calendar/application/entry"
	"advent-calendar/domain/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserEntryListFindController struct {
	userEntryListFindService entry_app.UserEntryListFindService
}

func NewUserEntryListFindController(userEntryListFindService entry_app.UserEntryListFindService) *UserEntryListFindController {
	c := new(UserEntryListFindController)
	c.userEntryListFindService = userEntryListFindService
	return c
}

func (controller *UserEntryListFindController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Param("user_id")
		if len(v) == 0 {
			c.Error(errors.New("user_id is required"))
			return
		}

		userId, err := user.CreateIdFrom(v)
		if err != nil {
			c.Error(err)
			return
		}

		userEntryList, err := controller.userEntryListFindService.Find(userId)
		if err != nil {
			c.Error(err)
			return
		}

		entries := []gin.H{}
		for _, e := range userEntryList.EntryList {
			entries = append(entries, gin.H{
				"calendar_id":   e.Calendar.Id,
				"calendar_name": e.Calendar.Name,
				"entried_date":  e.EntriedDate,
			})
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"detail": gin.H{
				"user_id":   userEntryList.Id,
				"user_name": userEntryList.UserName,
				"entries":   entries,
			},
		})
	}
}
