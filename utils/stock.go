package utils

import "fmt"

// StockStatus defines inventory condition.
type StockStatus string

const (
	StatusGood      StockStatus = "good"
	StatusDamaged   StockStatus = "damaged"
	StatusUnavailable StockStatus = "unavailable"
)

// CalcAvailableStock returns total available after reserved and damaged.
func CalcAvailableStock(total, reserved, damaged int64) int64 {
	available := total - reserved - damaged
	if available < 0 {
		return 0
	}
	return available
}

// UpdateStockIn returns new total + qty.
func UpdateStockIn(total, qty int64) (int64, error) {
	if qty < 0 {
		return total, fmt.Errorf("quantity in must be >= 0")
	}
	return total + qty, nil
}

// UpdateStockOut returns new total after out and error if insufficient.
func UpdateStockOut(total, qty int64) (int64, error) {
	if qty < 0 {
		return total, fmt.Errorf("quantity out must be >= 0")
	}
	if qty > total {
		return total, fmt.Errorf("insufficient stock")
	}
	return total - qty, nil
}

// RecordDamage returns new damaged value and new stock value.
func RecordDamage(total, damaged, qty int64) (int64, int64, error) {
	if qty < 0 {
		return total, damaged, fmt.Errorf("damage qty must be >=0")
	}
	if qty > total-damaged {
		return total, damaged, fmt.Errorf("damage qty exceeds good stock")
	}
	damaged += qty
	return total, damaged, nil
}

// ReorderPoint returns reorder point.
func ReorderPoint(dailyUsage, leadDays, safetyStock int64) int64 {
	return dailyUsage*leadDays + safetyStock
}
