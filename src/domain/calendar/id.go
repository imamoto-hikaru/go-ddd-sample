package calendar

import "strconv"

type Id uint

func CreateIdFrom(s string) (Id, error) {
	value, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return Id(0), err
	}
	return Id(value), nil
}
