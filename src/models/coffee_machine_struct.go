package models

type CoffeeInputStruct struct {
	Machine MachineStruct `json:"machine"`
}
type MachineStruct struct {
	Outlets OutletStruct `json:"outlets"`
	TotalItemsQuantity ItemQuantityStruct`json:"total_items_quantity"`
	Beverages    map[string]ItemQuantityStruct `json:"beverages"`
}

type OutletStruct struct {
	CountN int `json:"count_n"`
}

type ItemQuantityStruct struct {
	HotWater       int `json:"hot_water"`
	HotMilk        int `json:"hot_milk"`
	GingerSyrup    int `json:"ginger_syrup"`
	SugarSyrup     int `json:"sugar_syrup"`
	TeaLeavesSyrup int `json:"tea_leaves_syrup"`
	GreenMixture   int `json:"green_mixture"`
	NotAvailable    []string `json:"not_available"`
}

