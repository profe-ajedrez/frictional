package main

import (
	"fmt"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
)

// udfs stands for unsafe decimal from string.
// helps to have a decimal value from string ignoring errors.
func udfs(s string) alpacadecimal.Decimal {
	d, _ := alpacadecimal.NewFromString(s)
	return d
}

func main() {
	// Create a new FromUnitValue instance with an initial value of 232.5
	entry := udfs("232.5")
	b := frictional.NewFromUnitValue(entry)

	// Bind a visitor to apply a quantity of 3
	b.Bind(frictional.WithQTY(udfs("3")))

	// Create a new TaxHandlerFromUnitValue instance
	th := frictional.NewTaxHandlerFromUnitValue()
	th.WithPercentualTax(udfs("16")) // Set the tax rate to 16%

	// Snapshot the net value before applying tax
	net := frictional.SnapshotVisitor{}
	net.Do(b)

	// Apply the tax handler to the Frictional instance
	th.Do(b)

	// Snapshot the brute value after applying tax
	brute := frictional.SnapshotVisitor{}
	brute.Do(b)

	// Print the results
	fmt.Printf("Net value (before tax): %v\nBrute value (after tax): %v\nTotal ratio: %v\nTotal amount: %v\nTaxable amount: %v", net.Get().String(), brute.Get().String(), th.TotalRatio().String(), th.TotalAmount().String(), th.Taxable().String())
	// Output: Net value (before tax): 697.5
	//Brute value (after tax): 809.1
	//Total ratio: 16
	//Total amount: 111.6
	//Taxable amount: 697.5
}
