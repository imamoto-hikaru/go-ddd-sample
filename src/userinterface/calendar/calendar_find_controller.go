package calendar

import (
	calendar_app "advent-calendar/application/calendar"
	"advent-calendar/domain/calendar"
	"errors"

	"github.com/gin-gonic/gin"
)

type CalendarFindController struct {
	calendarFindService calendar_app.CalendarFindService
}

func NewCalendarFindController(calendarFindService calendar_app.CalendarFindService) *CalendarFindController {
	c := new(CalendarFindController)
	c.calendarFindService = calendarFindService
	return c
}

func (controller *CalendarFindController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Param("calendar_id")
		if len(v) == 0 {
			c.Error(errors.New("calendar_id is required"))
			return
		}

		calendarId, err := calendar.CreateIdFrom(v)
		if err != nil {
			c.Error(err)
			return
		}

		calendar, err := controller.calendarFindService.Find(calendarId)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"calendar": gin.H{
				"id":   calendar.Id,
				"name": calendar.Name,
			},
		})
	}
}
