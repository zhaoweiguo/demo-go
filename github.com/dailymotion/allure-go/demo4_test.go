package allure_go

import (
	"github.com/dailymotion/allure-go"
	"testing"
)

func TestAllureSetupTeardown(t *testing.T) {
	allure.BeforeTest(t,
		allure.Description("setup"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Setup step 1"),
				allure.Action(func() {}))
		}))

	allure.Test(t,
		allure.Description("actual test"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Test step 1"),
				allure.Action(func() {}))
		}))

	allure.AfterTest(t,
		allure.Description("teardown"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Teardown step 1"),
				allure.Action(func() {}))
		}))
}
