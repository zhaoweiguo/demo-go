package allure_go

import (
	"fmt"
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestStep(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		fmt.Println("This block of code is a test")
	}))
}
