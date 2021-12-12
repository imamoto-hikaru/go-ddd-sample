//go:build wireinject

package main

import (
	"advent-calendar/application"
	"advent-calendar/infrastructure"
	"advent-calendar/userinterface/calendar"
	"advent-calendar/userinterface/entry"
	"advent-calendar/userinterface/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserCreateController(db *gorm.DB) *user.UserCreateController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		user.NewUserCreateController,
	)
	return new(user.UserCreateController)
}

func InitializeUserFindController(db *gorm.DB) *user.UserFindController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		user.NewUserFindController,
	)
	return new(user.UserFindController)
}

func InitializeCalendarCreateController(db *gorm.DB) *calendar.CalendarCreateController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		calendar.NewCalendarCreateController,
	)
	return new(calendar.CalendarCreateController)
}

func InitializeCalendarFindController(db *gorm.DB) *calendar.CalendarFindController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		calendar.NewCalendarFindController,
	)
	return new(calendar.CalendarFindController)
}

func InitializeEntryCreateController(db *gorm.DB) *entry.EntryCreateController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		entry.NewEntryCreateController,
	)
	return new(entry.EntryCreateController)
}

func InitializeCalendarEntryListFindController(db *gorm.DB) *entry.CalendarEntryListFindController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		entry.NewCalendarEntryListFindController,
	)
	return new(entry.CalendarEntryListFindController)
}

func InitializeUserEntryListFindController(db *gorm.DB) *entry.UserEntryListFindController {
	wire.Build(
		infrastructure.RepositoryProviderSet,
		application.ServiceProviderSet,
		entry.NewUserEntryListFindController,
	)
	return new(entry.UserEntryListFindController)
}
