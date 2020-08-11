package controllers_test

import (
	"coffee_machine/src/config"
	"coffee_machine/src/controllers"
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TestInput struct {
	Inputs []RequestStruct `json:"inputs"`
	Outputs []ResponseStruct `json:"outputs"`
}

type RequestStruct struct {
	 Requests [][]string `json:"requests"`
}

type ResponseStruct struct {
	Responses [][]string `json:"responses"`
}

var inputJson TestInput

func init() {
	//loading input json for machine
	config.Init()

	//loading test case json
	filePath := os.Getenv("PROJECT_PATH")+"/src/controllers/drink_request_test.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Test Cases:", "Error in taking input Test  Json: ", err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &inputJson)
	if err != nil {
		log.Println("Test Cases:", "Error in taking input Test  Json: ", err)
		os.Exit(1)
	}
}

func TestDrinkRequests(t *testing.T) {
	//Test Case 1:
	// requests:[["hot_coffee", "hot_tea", "green_tea"], ["black_tea"]]
	// responses:"responses": [
	//        [ "hot_tea is prepared",  "hot_coffee is prepared", "green_tea cannot be prepared because green_mixture is not available" ],
	//        [ "black_tea cannot be prepared because item hot_water is not sufficient"]
	//      ]
	for index, request := range inputJson.Inputs[0].Requests {
		assert.Equal(t, controllers.DrinkRequests(request), inputJson.Outputs[0].Responses[index],  fmt.Sprintf("Test Case 1: Issue in this Request: %+v\n", request))
	}

	//Taking Coffee Machine to original State
	config.Init()
	//Test case 2:
	//"requests": [ [ "hot_tea", "black_tea", "green_tea" ], ["hot_coffee"] ]
	// "responses": [
	// [ "hot_tea is prepared", "black_tea is prepared", "green_tea cannot be prepared because green_mixture is not available" ],
	//        [  "hot_coffee cannot be prepared because item hot_water is not sufficient" ]
	//      ]
	for index, request := range inputJson.Inputs[1].Requests {
		assert.Equal(t, controllers.DrinkRequests(request), inputJson.Outputs[1].Responses[index],  fmt.Sprintf("Test Case 2: Issue in this Request: %+v\n", request))
	}

	//Taking Coffee Machine to original State
	config.Init()
	//Test Case 3:
	//"requests": [ ["hot_tea","black_tea"], ["green_tea", "hot_coffee"]  ]
	//"responses":[
	// ["hot_tea is prepared",  "black_tea is prepared" ],
	// [  "green_tea cannot be prepared because green_mixture is not available", "hot_coffee cannot be prepared because item hot_water is not sufficient" ]
	//]
	for index, request := range inputJson.Inputs[2].Requests {
		assert.Equal(t, controllers.DrinkRequests(request), inputJson.Outputs[2].Responses[index],  fmt.Sprintf("Test Case 3: Issue in this Request: %+v\n", request))
	}


	//Taking Coffee Machine to original State
	config.Init()
	//Test Case 4:
	//"requests": [ ["hot_coffXXX"],["black_tea"], ["hot_tea"]  ]
	//"responses": [
	//			["Invalid drink"],
	//			["black_tea is prepared"],
	//			["hot_tea is prepared"]
	//]
	for index, request := range inputJson.Inputs[3].Requests {
		assert.Equal(t, controllers.DrinkRequests(request), inputJson.Outputs[3].Responses[index],  fmt.Sprintf("Test Case 4: Issue in this Request: %+v\n", request))
	}

	//Taking Coffee Machine to original State
	config.Init()
	//Test Case 5:
	//"requests": [ ["hot_coffee","black_tea", "hot_tea", "green_tea"]  ]
	//"responses": [
	//			["Only 3 parallel requests are allowed"]
	//]
	for index, request := range inputJson.Inputs[4].Requests {
		assert.Equal(t, controllers.DrinkRequests(request), inputJson.Outputs[4].Responses[index],  fmt.Sprintf("Test Case 4: Issue in this Request: %+v\n", request))
	}
}

