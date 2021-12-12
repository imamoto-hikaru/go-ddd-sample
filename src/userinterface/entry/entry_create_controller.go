package entry

import (
	entry_app "advent-calendar/application/entry"
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/entry"
	"advent-calendar/domain/user"
	"errors"

	"github.com/gin-gonic/gin"
)

type EntryCreateController struct {
	entryCreateService entry_app.EntryCreateService
}

func NewEntryCreateController(entryCreateService entry_app.EntryCreateService) *EntryCreateController {
	c := new(EntryCreateController)
	c.entryCreateService = entryCreateService
	return c
}

func (controller *EntryCreateController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdForm := c.PostForm("user_id")
		calendarIdForm := c.PostForm("calendar_id")
		entriedDateForm := c.PostForm("entried_date")
		if len(userIdForm) == 0 || len(calendarIdForm) == 0 || len(entriedDateForm) == 0 {
			if len(userIdForm) == 0 {
				c.Error(errors.New("user_id is required"))
			}
			if len(calendarIdForm) == 0 {
				c.Error(errors.New("calendar_id is required"))
			}
			if len(entriedDateForm) == 0 {
				c.Error(errors.New("entried_date is required"))
			}
			return
		}

		userId, errUserId := user.CreateIdFrom(userIdForm)
		if errUserId != nil {
			c.Error(errUserId)
			return
		}
		calendarId, errCalendarId := calendar.CreateIdFrom(calendarIdForm)
		if errCalendarId != nil {
			c.Error(errCalendarId)
			return
		}
		entriedDate, errEntriedDate := entry.CreateEntriedDateFrom(entriedDateForm)
		if errEntriedDate != nil {
			c.Error(errEntriedDate)
			return
		}

		err := controller.entryCreateService.Create(
			entry.EntrySheet{
				UserId:      userId,
				CalendarId:  calendarId,
				EntriedDate: entriedDate,
			},
		)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}
