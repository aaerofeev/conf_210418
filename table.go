package main

import (
	"fmt"
)

// Форматирует единицы площади
func SquareFormatter(square int) string {
	if square <= 0 {
		return "-"
	}
	
	return fmt.Sprintf("%d м²", square)
}
