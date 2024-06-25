package main

import (
	"fmt"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
)

func udfs(s string) alpacadecimal.Decimal {
	d, _ := alpacadecimal.NewFromString(s)
	return d
}

func main() {
	// Create a new FromUnitValue instance with an initial value of 100
	entry := udfs("100")
	b := frictional.NewFromUnitValue(entry)

	// Create an UnbufferedPercTax instance with a tax ratio of 10%
	taxRatio := udfs("10")
	upt := frictional.NewUnbufferedPercTax(taxRatio)

	// Get the value from the Frictional instance
	value := b.Value()

	// Apply the UnbufferedPercTax to the Frictional instance's value
	upt.Do(b)

	// Print the results
	fmt.Printf("Initial value: %v\n", value)
	fmt.Printf("Value after applying tax: %v\n", b.Value())
	fmt.Printf("Tax ratio: %v\n", upt.Ratio())
	fmt.Printf("Tax amount: %v\n", upt.Amount())
	fmt.Printf("Taxable amount: %v\n", upt.Taxable())

	// Output:
	//	Initial value: 100
	//	Value after applying tax: 110
	//	Tax ratio: 10
	//	Tax amount: 10
	//	Taxable amount: 100
}
