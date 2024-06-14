package helpers

import "errors"

func ValidateDates(dates []string) error {
	for _, date := range dates {
		_, err := ParseDate(date)
		if err != nil {
			return errors.New("dates are not valid")
		}
	}
	return nil
}
