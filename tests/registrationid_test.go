package tests

import (
	"fmt"
	"testing"

	"github.com/arugaz/libsignal/util/keyhelper"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
