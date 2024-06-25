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
	const maxScale = 12
	const scaleForNet = 6

	bg := frictional.NewFromBruteDefault().WithBrute(udfs("1619.1"))

	brute := &frictional.SnapshotVisitor{}
	net := &frictional.SnapshotVisitor{}
	netWD := &frictional.SnapshotVisitor{}
	unitValue := frictional.NewUnitValue(udfs("3"))

	bg.Bind(brute)
	bg.Bind(frictional.NewPercentualUnTax(udfs("16")))
	bg.Bind(net)
	bg.Bind(frictional.NewPercentualUnDiscount(udfs("0")))
	bg.Bind(netWD)
	bg.Bind(unitValue)
	bg.Bind(frictional.NewRound(maxScale))

	netRounded := net.Get().Round(scaleForNet)
	unitValue.Round(maxScale)

	fmt.Printf("Brute value: %v\nNet value: %v\nNet rounded: %v\nNet value with discount: %v\nUnit value: %v\nBuffer value: %v",
		brute.Get().String(), net.Get().String(), netRounded.String(), netWD.Get().String(), unitValue.Get().String(), bg.Value().String())
}
