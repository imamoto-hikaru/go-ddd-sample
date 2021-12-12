//go:build wireinject

package infrastructure

import (
	calendar_domain "advent-calendar/domain/calendar"
	entry_domain "advent-calendar/domain/entry"
	user_domain "advent-calendar/domain/user"
	"advent-calendar/infrastructure/calendar"
	"advent-calendar/infrastructure/entry"
	"advent-calendar/infrastructure/user"

	"github.com/google/wire"
)

var RepositoryProviderSet = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user_domain.UserRepository), new(*user.UserRepository)),
	calendar.NewCalendarRepository,
	wire.Bind(new(calendar_domain.CalendarRepository), new(*calendar.CalendarRepository)),
	entry.NewEntryRepository,
	wire.Bind(new(entry_domain.EntryRepository), new(*entry.EntryRepository)),
)
