package entry

import (
	entry_app "advent-calendar/application/entry"
	"advent-calendar/domain/calendar"
	"errors"

	"github.com/gin-gonic/gin"
)

type CalendarEntryListFindController struct {
	calendarEntryListFindService entry_app.CalendarEntryListFindService
}

func NewCalendarEntryListFindController(calendarEntryListFindService entry_app.CalendarEntryListFindService) *CalendarEntryListFindController {
	c := new(CalendarEntryListFindController)
	c.calendarEntryListFindService = calendarEntryListFindService
	return c
}

func (controller *CalendarEntryListFindController) Handler() gin.HandlerFunc {
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

		calendarEntryList, err := controller.calendarEntryListFindService.Find(calendarId)
		if err != nil {
			c.Error(err)
			return
		}

		entries := []gin.H{}
		for _, e := range calendarEntryList.EntryList {
			entries = append(entries, gin.H{
				"user_id":      e.User.Id,
				"user_name":    e.User.Name,
				"entried_date": e.EntriedDate,
			})
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"detail": gin.H{
				"calendar_id":   calendarEntryList.Id,
				"calendar_name": calendarEntryList.CalendarName,
				"entries":       entries,
			},
		})
	}
}
