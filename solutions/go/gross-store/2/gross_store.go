package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	u := make(map[string]int)
    u["quarter_of_a_dozen"] = 3
    u["half_of_a_dozen"] =6
    u["dozen"] =12
    u["small_gross"] = 120
    u["gross"] =144
    u["great_gross"] = 1728
    return u
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _,e := units[unit]; !e {
        return false
    }
    bill[item] += units[unit]
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_,e := bill[item]
    _,e1 := units[unit]
    if !e || !e1 || bill[item] < units[unit] {
        return false
    } else if bill[item] == units[unit]{
        delete(bill, item)
        return true
    } 
    bill[item] -= units[unit]
    return true
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v,e := bill[item]
    return v,e
}
