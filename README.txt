Command for running the test Cases:
    1. export PROJECT_PATH = $(pwd)
    2. go test src/controllers/coffee_machine_test.go


 Assumptions:
 I. Only 5 kind of Drinks are Allowed
    1.hot_tea 2.hot_coffee 3.black_tea 4.green_tea 5.hot_water
    Input is going to be these 5 drinks only
 II. Any Drink can have requirement of these ingredients only
    HotWater, HotMilk, GingerSyrup, SugarSyrup, TeaLeavesSyrup, GreenMixtrue


Project Directory Structure:

src folders contain 4 sub folders:

1. config: It consist of Machine input config given by Dunzo Team, code for loading it and global variable
           CoffeeMachineInput(which contains machine info)
2. controllers:
    a) It consist of function for Drink Request (function: DrinkRequests filename: coffee_machine.go)
    b) It consist of function for refilling the machine (function: refill filename: coffee_machine.go)
    c) It consist of function for checking the ingredient requirement and updating quantity
       (functions: checkDrinkRequirement, getUpdatedItemsQuantiy, filename: coffee_machine_helper.go)
    d) It consist of unit test cases for the above two functions (filename: coffee_machine_test.go)

3. models: It consist of structs for storing the machine infos
4. utils:  It consist of logic for notifying in case of ingredients are not available(function NotifyEmptyIngredient filename: utils.go)

