package di_test

import (
	"testing"

	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
)

type TestStruct struct {
	TestField string
}

func TestContainerWithCorrectType(t *testing.T) {
	str := "test"
	dummy := TestStruct{str}
	b := di.New().RegisterService(identificators.Test, dummy)
	c := b.Build()
	res, err := di.GetService[TestStruct](c, identificators.Test)
	if err != nil {
		t.Error("invalid type")
		return
	}

	if res.TestField != str {
		t.Errorf("%s %s", res.TestField, str)
		return
	}
}

func TestContainerWithWrongType(t *testing.T) {
	b := di.New().RegisterService(identificators.Test, "test")
	c := b.Build()
	_, err := di.GetService[TestStruct](c, identificators.Test)
	if err == nil {
		t.Error("Didn`t handle the type ok")
		return
	}

}
