package entry

import "advent-calendar/domain/user"

type UserEntryList struct {
	Id        user.Id
	UserName  user.Name
	EntryList []Entry
}
