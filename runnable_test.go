package frictional_test

import (
	"fmt"

	"github.com/alpacahq/alpacadecimal"
	"github.com/profe-ajedrez/frictional"
)

func udfs(s string) alpacadecimal.Decimal {
	d, _ := alpacadecimal.NewFromString(s)
	return d
}

// ExampleTaxHandlerFromUnitValue demonstrates the usage of the TaxHandlerFromUnitValue type.
// It creates a new Frictional instance with an initial value of 232.5 and a quantity of 3,
// then applies a 16% tax rate using the TaxHandlerFromUnitValue. The net value before tax,
// the brute value after tax, the total ratio, and the taxable amount are printed.
func ExampleTaxHandlerFromUnitValue() {
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
	// Brute value (after tax): 809.1
	// Total ratio: 16
	// Total amount: 111.6
	// Taxable amount: 697.5
}

func ExampleFromBrute() {
	// Create a new FromBrute instance with a default buffer
	bg := frictional.NewFromBruteDefault().WithBrute(udfs("1619.1"))

	// Create a SnapshotVisitor to capture the brute value
	brute := &frictional.SnapshotVisitor{}

	// Create a SnapshotVisitor to capture the net value
	net := &frictional.SnapshotVisitor{}

	// Create a SnapshotVisitor to capture the net value with discount
	netWD := &frictional.SnapshotVisitor{}

	// Create a UnitValue instance with a quantity of 3
	unitValue := frictional.NewUnitValue(udfs("3"))

	// Bind the visitors to the FromBrute instance
	bg.Bind(brute)
	bg.Bind(frictional.NewPercentualUnTax(udfs("16")))
	bg.Bind(net)
	bg.Bind(frictional.NewPercentualUnDiscount(udfs("0")))
	bg.Bind(netWD)
	bg.Bind(unitValue)
	bg.Bind(frictional.NewRound(12))

	netRounded := net.Get().Round(6)

	// Round the unitValue to 12 decimal places
	unitValue.Round(12)

	// Print the results
	fmt.Printf("Brute value: %v\nNet value: %v\nNet rounded: %v\nNet value with discount: %v\nUnit value: %v\nBuffer value: %v", brute.Get().String(), net.Get().String(), netRounded.String(), netWD.Get().String(), unitValue.Get().String(), bg.Buffer().String())
	// Output: Brute value: 1619.1
	// Net value: 1395.7758620689655172
	// Net rounded: 1395.775862
	// Net value with discount: 1395.7758620689655172
	// Unit value: 465.258620689655
	// Buffer value: 465.258620689655
}
