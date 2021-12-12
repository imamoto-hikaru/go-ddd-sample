package entry

import "advent-calendar/domain/calendar"

type CalendarEntryList struct {
	Id           calendar.Id
	CalendarName calendar.Name
	EntryList    []Entry
}

func (cel CalendarEntryList) CanAddEntryTo(entriedDate EntriedDate) bool {
	for _, e := range cel.EntryList {
		if e.EntriedDate.IsEqual(entriedDate) {
			return false
		}
	}
	return true
}
