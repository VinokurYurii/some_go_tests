package shuffler

import (
	"fmt"
	"math/rand"
	"reflect"
)

func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice {
		fmt.Println("Shuffle: provided interface not a slice type")
		return
	}

	lenght := rv.Len()
	swap := reflect.Swapper(slice)
	rand.Shuffle(lenght, func(i, j int) {
		swap(i, j)
	})
}
