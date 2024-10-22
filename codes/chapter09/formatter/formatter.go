package print

import "fmt"

// Format returns a formatted string
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
