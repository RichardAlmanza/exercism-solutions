package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// A more explicit way is: if _, exists := units[unit]; !exists {
	// but I ignore it because the Units map doesn't have any item
	// with a value of 0... yeah, the explicit way is the safest
	if units[unit] == 0 {
		return false
	}

	bill[item] += units[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if units[unit] == 0 {
		return false
	}

	// This is a simple test, so it is not too important to write it in
	// the safest way, but in a project, I can't give such a luxury of laziness.
	// Yeah… the laziness of writing, but look how long my comments are.
	//
	// In this case, there is no unit with the value of 0, and in this function,
	// every result with the value 0 will remove the item, then at this moment of
	// the function, it means the item doesn't exist
	if bill[item] == 0 {
		return false
	}

	var newAmount int = bill[item] - units[unit]

	if newAmount < 0 {
		return false
	}

	bill[item] = newAmount

	if newAmount == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemAmount, itemExistence := bill[item]
	return itemAmount, itemExistence
}
