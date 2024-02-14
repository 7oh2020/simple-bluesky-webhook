package validation

import (
	"fmt"
)

// textを検証します。
func ValidateText(text string) error {
	if text == "" {
		return fmt.Errorf("can not be blank")
	}
	if len([]rune(text)) > 300 {
		return fmt.Errorf("length must be less than or equal to 300")
	}
	return nil
}
