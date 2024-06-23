package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestDiscount(t *testing.T) {
	for i, tc := range testCasesDiscounts {
		b, d, shouldFail, whatErrorShouldBe, err := tc.tester()

		if !shouldFail && err != nil {
			t.Logf("[FAIL test case %d] %v", i, err)
			t.FailNow()
		}

		if shouldFail && err == nil {
			t.Logf("[FAIL test case %d %s] error expected: %v", i, tc.name, whatErrorShouldBe)
			t.FailNow()
		}

		if !shouldFail && !b.buffer.Equal(tc.expecteds.buffer) {
			t.Logf("[FAIL test case %d %s] got frictional.buffer %v. Expected %v", i, tc.name, b.buffer, tc.expecteds.buffer)
			t.FailNow()
		}

		if !shouldFail && !d.ratio.Equal(tc.expecteds.ratio) {
			t.Logf("[FAIL test case %d %s] got discount.ratio %v. Expected %v", i, tc.name, d.ratio, tc.expecteds.ratio)
			t.FailNow()
		}

		if !shouldFail && !d.amount.Equal(tc.expecteds.amount) {
			t.Logf("[FAIL test case %d] got discount.amount %v. Expected %v", i, d.amount, tc.expecteds.amount)
			t.FailNow()
		}
	}
}

func BenchmarkDiscount(b *testing.B) {
	entry, _ := alpacadecimal.NewFromString("100.123")
	ratio := unsafeDecFromStr("10")
	ratio2 := unsafeDecFromStr("15")
	qty := unsafeDecFromStr("10")

	discountOverEntryValue := NewPercentualDiscount(ratio)
	discountConsideringQty := NewPercentualDiscount(ratio2)

	for i := 0; i <= b.N; i++ {
		bg := NewFromUnitValue(entry)

		bg.Bind(discountOverEntryValue)
		bg.Bind(WithQTY(qty))
		bg.Bind(discountConsideringQty)
	}
}

// testCasesDiscounts test cases list
var testCasesDiscounts = []struct {
	name string
	// tester is a function that implements test cases
	// should return the *frictional instance constructed by the case, the one from Discount,
	// a bool indicating whether the case should end with an error,
	// the possible error or, failing that, nil
	// and a string explaining why it should end in an error if this should happen
	tester func() (*FromUnitValue, Discount, bool, string, error)

	// expecteds contiene un struct con los datos que la función tester debería producir
	expecteds struct {
		// FromUnitValue debe contener los valores que el *FromUnitValue devuelto por tester debería contener
		FromUnitValue

		// Discount debe contener los valores que el Discount devuelto por tester debería contener
		Discount
	}
}{
	{
		name: "Discount over entry unit value",
		tester: func() (*FromUnitValue, Discount, bool, string, error) {
			// Se crea un entry unit value de tipo decimal
			entry, _ := alpacadecimal.NewFromString("3.453561112")

			// Se crea un frictional de scala 12 con el entry value
			b := NewFromUnitValue(entry)

			// Se define un valor de descuento porcentual de 10%
			ratio := alpacadecimal.NewFromInt32(10)

			// Se crea el Evaluer con el ratio de descuento indicado
			d1 := NewPercentualDiscount(ratio)

			// Se bindea para evaluación al Evaluer d1, que es un descuento porcentual
			b.Bind(d1)

			// En este punto, b.buffer debería contener el valor de b.entryValue - 10% del valor de d1.ratio
			// y d1.amount debería contener el valor equivalente al 10% de b.entryValue
			return b, d1.Discount, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Discount
		}{
			FromUnitValue: FromUnitValue{
				defaultFrictional: &defaultFrictional{
					// expected buffer should be 90% of entryValue, because discount is 10%
					buffer: unsafeDecFromStr("3.453561112").Mul(unsafeDecFromStr("0.9")),
				},
			},
			Discount: Discount{
				ratio: unsafeDecFromStr("10"),
				// expected amount should be 10% of entryValue
				amount: unsafeDecFromStr("3.453561112").Mul(unsafeDecFromStr("0.1")),
			},
		},
	},
	{
		name: "Amount Discount over entry unit value",
		tester: func() (*FromUnitValue, Discount, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("3.453561112")
			b := NewFromUnitValue(entry)
			amount := unsafeDecFromStr("1.834566333")
			d1 := NewAmountDiscount(amount)
			b.Bind(d1)

			return b, d1.Discount, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Discount
		}{
			FromUnitValue: FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("3.453561112").Sub(udfs("1.834566333")),
				},
			},
			Discount: Discount{
				ratio:  Hundred().Mul(udfs("1.834566333")).Div(udfs("3.453561112")),
				amount: unsafeDecFromStr("1.834566333"),
			},
		},
	},
	{
		name: "Combo Discount percentual over entry unit value and other considering quantity",
		tester: func() (*FromUnitValue, Discount, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("100.123")
			b := NewFromUnitValue(entry)

			ratio := unsafeDecFromStr("10")
			discountOverEntryValue := NewPercentualDiscount(ratio)

			b.Bind(discountOverEntryValue)

			ratio = unsafeDecFromStr("15")
			discountConsideringQty := NewPercentualDiscount(ratio)

			qty := unsafeDecFromStr("10")

			b.Bind(WithQTY(qty))
			b.Bind(discountConsideringQty)

			totalDiscountApplied := Discount{
				ratio:  udfs("1"),
				amount: udfs("100"),
			}

			return b, totalDiscountApplied, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Discount
		}{
			FromUnitValue: FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("100.123").Mul(udfs("0.9")).Mul(udfs("10")).Mul(udfs("0.85")),
				},
			},
			Discount: Discount{
				ratio:  udfs("1"),
				amount: udfs("100"),
			},
		},
	},
}

func udfs(d string) alpacadecimal.Decimal {
	return unsafeDecFromStr(d)
}

func unsafeDecFromStr(d string) alpacadecimal.Decimal {
	dec, _ := alpacadecimal.NewFromString(d)
	return dec
}
