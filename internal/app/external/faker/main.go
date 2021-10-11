package faker

import (
	"reflect"

	fakerTag "github.com/bxcodec/faker/v3"
)

func Generator() {
	_ = fakerTag.AddProvider("keyNameFaker", func(v reflect.Value) (interface{}, error) {
		return Key().Name(), nil
	})
}
