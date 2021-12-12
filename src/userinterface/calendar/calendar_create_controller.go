package calendar

import (
	calendar_app "advent-calendar/application/calendar"
	"advent-calendar/domain/calendar"
	"errors"

	"github.com/gin-gonic/gin"
)

type CalendarCreateController struct {
	calendarCreateService calendar_app.CalendarCreateService
}

func NewCalendarCreateController(calendarCreateService calendar_app.CalendarCreateService) *CalendarCreateController {
	c := new(CalendarCreateController)
	c.calendarCreateService = calendarCreateService
	return c
}

func (controller *CalendarCreateController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		calendarName := c.PostForm("calendar_name")
		if len(calendarName) == 0 {
			c.Error(errors.New("calendar_name is required"))
			return
		}

		calendar, err := controller.calendarCreateService.Create(calendar.Name(calendarName))
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
