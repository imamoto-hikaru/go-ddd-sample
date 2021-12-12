//go:build wireinject

package application

import (
	"advent-calendar/application/calendar"
	"advent-calendar/application/entry"
	"advent-calendar/application/user"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	user.NewUserCreateService,
	user.NewUserFindService,
	calendar.NewCalendarCreateService,
	calendar.NewCalendarFindService,
	entry.NewEntryCreateService,
	entry.NewCalendarEntryListFindService,
	entry.NewUserEntryListFindService,
)
