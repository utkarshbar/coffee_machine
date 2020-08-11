package controllers

import (
	"coffee_machine/src/config"
	"coffee_machine/src/models"
)

func checkQuantity(ingredient string, reqAmount, totalAmountAvailable int) (unavailableInsufficient string) {
	//Function check the quantiy of ingredient
	// if it is sufficient then it return empty string
	if reqAmount <= totalAmountAvailable {
		return
	}
	// if ingredient available in NotAvailable list
	for _, notAvailableIngredient := range config.NotAvailableIngredients {
		if ingredient == notAvailableIngredient {
			return "not available"
		}
	}
	// if ingredient available amount is not sufficient  then it return not sufficient
	return "not sufficient"
}

func checkDrinkRequirement(drinkRequirement models.ItemQuantityStruct) (string, string) {
	//Function checks for ingredients available or not
	totalItemsQuantity := config.CoffeeMachineInput.Machine.TotalItemsQuantity
	var (
		unavailableInsufficient string
	)
	if unavailableInsufficient = checkQuantity("ginger_syrup", drinkRequirement.GingerSyrup, totalItemsQuantity.GingerSyrup); unavailableInsufficient != "" {
		return "ginger_syrup", unavailableInsufficient
	}
	if unavailableInsufficient = checkQuantity("green_mixture", drinkRequirement.GreenMixture, totalItemsQuantity.GreenMixture); unavailableInsufficient != "" {
		return "green_mixture", unavailableInsufficient
	}

	if unavailableInsufficient = checkQuantity("hot_milk", drinkRequirement.HotMilk, totalItemsQuantity.HotMilk); unavailableInsufficient != "" {
		return "hot_milk", unavailableInsufficient
	}
	if unavailableInsufficient = checkQuantity("hot_water", drinkRequirement.HotWater, totalItemsQuantity.HotWater); unavailableInsufficient != "" {
		return "hot_water", unavailableInsufficient
	}

	if unavailableInsufficient = checkQuantity("sugar_syrup", drinkRequirement.SugarSyrup, totalItemsQuantity.SugarSyrup); unavailableInsufficient != "" {
		return "sugar_syrup", unavailableInsufficient
	}
	if unavailableInsufficient = checkQuantity("tea_leaves_syrup", drinkRequirement.TeaLeavesSyrup, totalItemsQuantity.TeaLeavesSyrup); unavailableInsufficient != "" {
		return "tea_leaves_syrup", unavailableInsufficient
	}

	return "", ""
}

func getUpdatedItemsQuantiy(totalItemsQuantity, drinkRequirement models.ItemQuantityStruct) models.ItemQuantityStruct {
	//For Updating the Quantity of ingredient in the coffee machine
	return models.ItemQuantityStruct{
		totalItemsQuantity.HotWater - drinkRequirement.HotWater,
		totalItemsQuantity.HotMilk - drinkRequirement.HotMilk,
		totalItemsQuantity.GingerSyrup - drinkRequirement.GingerSyrup,
		totalItemsQuantity.SugarSyrup - drinkRequirement.SugarSyrup,
		totalItemsQuantity.TeaLeavesSyrup - drinkRequirement.TeaLeavesSyrup,
		totalItemsQuantity.GreenMixture - drinkRequirement.GreenMixture,
		config.NotAvailableIngredients,
	}
}
