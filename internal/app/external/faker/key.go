package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-aws-key-pair/internal/errors"
)

type FakeKey interface {
	Name() string // Name Key fake
}

type fakeKey struct{}

func Key() FakeKey {
	return fakeKey{}
}

var (
	names = []string{"Optimus-Prime", "Wheeljack", "Bumblebee"}
)

func (n fakeKey) Name() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(names))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameUuid := fmt.Sprintf("%s-%s", names[num.Int64()], shortuuid.New())
	return nameUuid
}
