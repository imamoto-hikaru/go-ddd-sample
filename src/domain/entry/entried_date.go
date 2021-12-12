package entry

import (
	"errors"
	"strconv"
)

type EntriedDate uint

func (ed EntriedDate) IsEqual(other EntriedDate) bool {
	return ed == other
}

func CreateEntriedDateFrom(s string) (EntriedDate, error) {
	value, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return EntriedDate(0), err
	}
	if value < 1 || value > 25 {
		return EntriedDate(0), errors.New("invalid entried_date")
	}
	return EntriedDate(value), nil
}
