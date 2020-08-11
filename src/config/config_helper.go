package config

import (
	"coffee_machine/src/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"coffee_machine/src/models"
)

// Global variable and contains complete machine Info
var (
	CoffeeMachineInput models.CoffeeInputStruct
	NotAvailableIngredients []string
)

func Init() {
	filePath := os.Getenv("PROJECT_PATH")+"/src/config/coffee_input.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("InitConf", "Error in taking input Json: ", err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &CoffeeMachineInput)
	if err != nil {
		log.Println("InitConf", "Error in taking input Json: ", err)
		os.Exit(1)
	}
	NotAvailableIngredients = utils.NotifyEmptyIngredient(CoffeeMachineInput.Machine.TotalItemsQuantity)
}