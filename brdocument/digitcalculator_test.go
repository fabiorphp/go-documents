package brdocument

import (
	"testing"
)

func TestDigitCalculator(t *testing.T) {
	t.Run("CPF 068432731", func(t *testing.T) {
		dc := NewDigitCalculator("068432731")
		dc.WithMultipliersInterval(2, 11)
		dc.UseComplementaryInsteadOfModule()
		dc.ReplaceWhen(0, []int{10, 11})
		dc.WithModule(Module11)
		firstDigit := dc.Calculate()
		dc.AddDigit(firstDigit)
		secondDigit := dc.Calculate()

		if firstDigit != 7 || secondDigit != 3 {
			t.Errorf("Digit not is 73")
		}
	})

}
