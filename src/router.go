package main

import (
	"advent-calendar/userinterface"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type endPoint struct {
	HttpMethod     string
	Path           string
	CtlInitializer func(db *gorm.DB) userinterface.Controller
}

var routings = []endPoint{
	{
		HttpMethod:     http.MethodPost,
		Path:           "/user/create",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeUserCreateController(db) },
	},
	{
		HttpMethod:     http.MethodGet,
		Path:           "/user/:user_id",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeUserFindController(db) },
	},
	{
		HttpMethod:     http.MethodGet,
		Path:           "/calendar/:calendar_id",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeCalendarFindController(db) },
	},
	{
		HttpMethod:     http.MethodPost,
		Path:           "/calendar/create",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeCalendarCreateController(db) },
	},
	{
		HttpMethod:     http.MethodGet,
		Path:           "/entries/calendar/:calendar_id",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeCalendarEntryListFindController(db) },
	},
	{
		HttpMethod:     http.MethodGet,
		Path:           "/entries/user/:user_id",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeUserEntryListFindController(db) },
	},
	{
		HttpMethod:     http.MethodPost,
		Path:           "/entry/create",
		CtlInitializer: func(db *gorm.DB) userinterface.Controller { return InitializeEntryCreateController(db) },
	},
}

func CreateRouter() *gin.Engine {
	r := gin.Default()
	for _, ep := range routings {
		ctlInitializer := ep.CtlInitializer
		handlerFunc := func(c *gin.Context) {
			db := c.MustGet("db").(*gorm.DB)
			(ctlInitializer(db).Handler())(c)
		}
		r.Handle(ep.HttpMethod, ep.Path, ApiTransactionManager, handlerFunc)
	}

	return r
}
