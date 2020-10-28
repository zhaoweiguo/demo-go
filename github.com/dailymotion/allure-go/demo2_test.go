package allure_go

import (
	"fmt"
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestStep2(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		allure.Step(allure.Description("Action"),
			allure.Action(func() {
				fmt.Println("This block of code is a step")
			}))
		PerformAction()
	}))
}

func PerformAction() {
	allure.Step(allure.Description("Action"),
		allure.Action(func() {
			fmt.Println("This block of code is also a step")
		}))
}
