package helpers

import "reflect"

func DeleteOfSliceR[T any](slice []T, element T) []T {
	//get index of element
	for index, value := range slice {
		if reflect.DeepEqual(value, element) {
			// Delete the element by appending all the elements before and after the index.
			return append((slice)[:index], (slice)[index+1:]...)
			break
		}
	}
	return slice
}

func DeleteOfSlice[T any](slice *[]T, element T) {
	//get index of element
	for index, value := range *slice {
		if reflect.DeepEqual(value, element) {
			// Delete the element by appending all the elements before and after the index.
			*slice = append((*slice)[:index], (*slice)[index+1:]...)
			break
		}
	}
}
