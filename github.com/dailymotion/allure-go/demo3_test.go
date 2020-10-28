package allure_go

import (
	"github.com/dailymotion/allure-go"
	"io/ioutil"
	"log"
	"testing"
)

func TestParameter(t *testing.T) {
	allure.Test(t,
		allure.Description("Test that has parameters"),
		allure.Parameter("testParameter", "test"),
		allure.Action(func() {
			allure.Step(allure.Description("Step with parameters"),
				allure.Action(func() {}),
				allure.Parameter("stepParameter", "step"))
		}))
}

func TestParameterized(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {
			allure.Test(t,
				allure.Description("Test with parameters"),
				allure.Parameter("counter", i),
				allure.Action(func() {}))
		})
	}
}

func TestImageAttachmentToStep(t *testing.T) {
	img, err := ioutil.ReadFile("./image.png")
	if err != nil {
		log.Panic(err)
	}
	allure.Test(t,
		allure.Description("testing image attachment"),
		allure.Action(func() {
			allure.Step(allure.Description("adding an image attachment"),
				allure.Action(func() {
					err := allure.AddAttachment("image", allure.ImagePng, img)
					if err != nil {
						log.Println(err)
					}
				}))
		}))
}
