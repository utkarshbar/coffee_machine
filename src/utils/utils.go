package utils

import (
	"fmt"

	"coffee_machine/src/models"
)

func NotifyEmptyIngredient(totalItemsQuantity models.ItemQuantityStruct) (notAvailable []string){
	//Checks if all the ingredients are sufficient or not
	if totalItemsQuantity.HotMilk == 0 {
		fmt.Println("hot_milk is not available")
		notAvailable = append(notAvailable, "hot_milk")
	}
	if totalItemsQuantity.HotWater == 0 {
		fmt.Println("hot water is not available")
		notAvailable = append(notAvailable, "hot_water")
	}
	if totalItemsQuantity.GreenMixture == 0 {
		fmt.Println("Green Mixture is not available")
		notAvailable = append(notAvailable, "green_mixture")
	}
	if totalItemsQuantity.TeaLeavesSyrup == 0 {
		fmt.Println("Tea Leaves Syrup is not available")
		notAvailable = append(notAvailable, "test_leaves_syrup")
	}
	if totalItemsQuantity.SugarSyrup == 0 {
		fmt.Println("Sugar Syrup is not available")
		notAvailable = append(notAvailable, "sugar_syrup")
	}
	if totalItemsQuantity.GingerSyrup == 0 {
		fmt.Println("Ginger Syrup is not available")
		notAvailable = append(notAvailable, "ginger_syrup")
	}
	return
}
