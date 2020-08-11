package controllers

import (
	"coffee_machine/src/utils"
	"fmt"

	"coffee_machine/src/config"
)



func DrinkRequests(drinks []string) (outputStatements []string) {
	// Function takes the drink request ehich can be parallel.
	// Ex - if 3 parallel requests allowed then
	// drinks - ["hot_milk", "hot_water", "green_tea"] or ["hot_milk"]

	//Check for n parallel request, if there are more than n parallel request
	if len(drinks) > config.CoffeeMachineInput.Machine.Outlets.CountN { //Check for n parallel request
		statement := fmt.Sprint("Only ", config.CoffeeMachineInput.Machine.Outlets.CountN, " parallel requests are allowed")
		outputStatements = append(outputStatements, statement)
		return
	}

	var missingIngredient, unavailableInsufficient, statement string
	for _, drink := range drinks {
		drinkRequirement, exist := config.CoffeeMachineInput.Machine.Beverages[drink]  // Taking Requirement of Drink
		if !exist { //Invalid drink Name
		    statement = "Invalid drink"
			outputStatements = append(outputStatements, statement)
			continue
		}
		totalItemsQuantity := config.CoffeeMachineInput.Machine.TotalItemsQuantity // Total Quantity Available
		//Checking drink requirement id fulfilled or not
		if missingIngredient, unavailableInsufficient = checkDrinkRequirement(drinkRequirement); unavailableInsufficient == "" {
			//Ingredient requirements are meet.
			// It will update quantity available of ingredient in machine
			config.CoffeeMachineInput.Machine.TotalItemsQuantity = getUpdatedItemsQuantiy(totalItemsQuantity, drinkRequirement)
			//Produce the Drink
			statement = fmt.Sprintf("%s is prepared", drink)
			//Calling function who will notify if any ingredient is now not available
		} else {
			// Case when ingredient is insufficient or Unavailable
			if missingIngredient == "hot_water" {
				missingIngredient = "item hot_water"
			}
			statement = fmt.Sprintf("%s cannot be prepared because %s is %s", drink, missingIngredient, unavailableInsufficient)
		}
		outputStatements = append(outputStatements, statement)
	}
	utils.NotifyEmptyIngredient(config.CoffeeMachineInput.Machine.TotalItemsQuantity)
	return
}

func Refill(ingredient string, amount int) {
	//Function for refilling the ingredient
	// It takes input ingredient name and amount in mil, we are adding
	// Ex- if I need to add 100 ml hot milk then ingredient = "hot_milk" and amount = 100
	ingredientQuantity := config.CoffeeMachineInput.Machine.TotalItemsQuantity
	switch ingredient {
	case "hot_milk":
		ingredientQuantity.HotMilk += amount
	case "hot_water":
		ingredientQuantity.HotWater += amount
	case "ginger_syrup":
		ingredientQuantity.GingerSyrup += amount
	case "green_mixture":
		ingredientQuantity.GreenMixture += amount
	case "sugar_syrup":
		ingredientQuantity.SugarSyrup += amount
	case "tea_leaves_syrup":
		ingredientQuantity.TeaLeavesSyrup += amount
	}
}
