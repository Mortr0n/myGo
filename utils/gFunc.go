package myGo

import "errors"

func RemoveSliceItem[T any](slice []T, itemNumber int) ([]T, error) {
	if itemNumber < 0 || itemNumber >= len(slice) {
		err := errors.New("index out of range")
		return nil, err
	}

	slice = append(slice[:itemNumber], slice[itemNumber+1:]...)
	return slice, nil
}
