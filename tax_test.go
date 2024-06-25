package frictional

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
)

func TestTax(t *testing.T) {
	for i, tc := range testCasesTaxes {
		b, d, shouldFail, whatErrorShouldBe, err := tc.tester()

		if !shouldFail && err != nil {
			t.Logf("[FAIL test case %d] %v", i, err)
			t.FailNow()
		}

		if shouldFail && err == nil {
			t.Logf("[FAIL test case %d %s] error expected: %v", i, tc.name, whatErrorShouldBe)
			t.FailNow()
		}

		if !shouldFail && !b.value.Equal(tc.expecteds.value) {
			t.Logf("[FAIL test case %d %s] got frictional.buffer %v. Expected %v", i, tc.name, b.value, tc.expecteds.value)
			t.FailNow()
		}

		if !shouldFail && !d.ratio.Equal(tc.expecteds.ratio) {
			t.Logf("[FAIL test case %d %s] got tax.ratio %v. Expected %v", i, tc.name, d.ratio, tc.expecteds.ratio)
			t.FailNow()
		}

		if !shouldFail && !d.amount.Equal(tc.expecteds.amount) {
			t.Logf("[FAIL test case %d] got tax.amount %v. Expected %v", i, d.amount, tc.expecteds.amount)
			t.FailNow()
		}

		if !shouldFail && !d.taxable.Equal(tc.expecteds.taxable) {
			t.Logf("[FAIL test case %d] got tax.taxable %v. Expected %v", i, d.taxable, tc.expecteds.taxable)
			t.FailNow()
		}
	}
}

func BenchmarkCreateTaxe(b *testing.B) {
	entry, _ := alpacadecimal.NewFromString("100.123")
	ratio := unsafeDecFromStr("10")

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		bg := NewFromUnitValue(entry)
		taxOverEntryValue := NewPercTax(ratio)
		bg.Bind(taxOverEntryValue)
	}
}

func BenchmarkTax(b *testing.B) {
	entry, _ := alpacadecimal.NewFromString("100.123")
	ratio := unsafeDecFromStr("10")
	ratio2 := unsafeDecFromStr("15")
	qty := unsafeDecFromStr("10")

	taxOverEntryValue := NewPercTax(ratio)
	taxConsideringQty := NewPercTax(ratio2)

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		bg := NewFromUnitValue(entry)

		bg.Bind(taxOverEntryValue)
		bg.Bind(WithQTY(qty))
		bg.Bind(taxConsideringQty)
	}
}

// testCasesTaxes test cases list for taxes
var testCasesTaxes = []struct {
	name string
	// tester is a function that implements test cases
	// should return the *frictional instance constructed by the case, the one from Tax,
	// a bool indicating whether the case should end with an error,
	// the possible error or, failing that, nil
	// and a string explaining why it should end in an error if this should happen
	tester func() (*FromUnitValue, Tax, bool, string, error)

	// expecteds contiene un struct con los datos que la función tester debería producir
	expecteds struct {
		// FromUnitValue debe contener los valores que el *FromUnitValue devuelto por tester debería contener
		FromUnitValue

		// Tax debe contener los valores que el Discount devuelto por tester debería contener
		Tax
	}
}{
	{
		name: "Tax over entry unit value",
		tester: func() (*FromUnitValue, Tax, bool, string, error) {
			// Se crea un entry unit value de tipo decimal
			entry, _ := alpacadecimal.NewFromString("17.3475345")

			// Se crea un frictional de scala 12 con el entry value
			b := NewFromUnitValue(entry)

			// Se define un valor de impuesto porcentual de 10%
			ratio := alpacadecimal.NewFromInt32(10)

			// Se crea el Evaluer con el ratio de impuesto indicado
			t1 := NewPercTax(ratio)

			// Se bindea para evaluación al Evaluer d1, que es un descuento porcentual
			b.Bind(t1)

			// En este punto, b.buffer debería contener el valor de b.entryValue - 10% del valor de d1.ratio
			// y d1.amount debería contener el valor equivalente al 10% de b.entryValue
			return b, t1.Tax, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Tax
		}{
			FromUnitValue: FromUnitValue{
				DefaultFrictional: &DefaultFrictional{
					// expected buffer should be 90% of entryValue, because discount is 10%
					value: udfs("17.3475345").Mul(udfs("1.1")),
				},
			},
			Tax: Tax{
				ratio: udfs("10"),
				// expected amount should be 10% of entryValue
				amount: udfs("17.3475345").Mul(udfs("0.1")),

				taxable: udfs("17.3475345"),
			},
		},
	},
	{
		name: "Amount tax over entry unit value",
		tester: func() (*FromUnitValue, Tax, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("100")
			b := NewFromUnitValue(entry)
			amount := udfs("9.8")
			t1 := NewAmountTax(amount)
			b.Bind(t1)

			return b, t1.Tax, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Tax
		}{
			FromUnitValue: FromUnitValue{
				DefaultFrictional: &DefaultFrictional{
					value: udfs("109.8"),
				},
			},
			Tax: Tax{
				ratio:   udfs("9.8"),
				amount:  udfs("9.8"),
				taxable: udfs("100"),
			},
		},
	},
	{
		name: "Combo Tax percentual over entry unit value and other considering quantity",
		tester: func() (*FromUnitValue, Tax, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("100.123")
			b := NewFromUnitValue(entry)

			ratio := unsafeDecFromStr("10")
			taxOverEntryValue := NewPercTax(ratio)

			b.Bind(taxOverEntryValue)

			ratio = unsafeDecFromStr("15")
			taxConsideringQty := NewPercTax(ratio)

			qty := unsafeDecFromStr("10")

			b.Bind(WithQTY(qty))
			b.Bind(taxConsideringQty)

			// Nos interesa validar el valor de b.buffer, pues no se deberían mezclar
			// impuestos calculados en distintos pasos
			totalTaxApplied := Tax{
				ratio:   udfs("1"),
				amount:  udfs("100"),
				taxable: udfs("1"),
			}

			return b, totalTaxApplied, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Tax
		}{
			FromUnitValue: FromUnitValue{
				DefaultFrictional: &DefaultFrictional{
					value: unsafeDecFromStr("1266.55595"),
				},
			},
			Tax: Tax{
				ratio:   udfs("1"),
				amount:  udfs("100"),
				taxable: udfs("1"),
			},
		},
	},
	{
		name: "Tax over entry unit value with quantity",
		tester: func() (*FromUnitValue, Tax, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("17.3475345")
			b := NewFromUnitValue(entry)
			ratio := alpacadecimal.NewFromInt32(10)
			t1 := NewPercTax(ratio)

			qty := udfs("5")
			b.Bind(WithQTY(qty))
			b.Bind(t1)

			return b, t1.Tax, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Tax
		}{
			FromUnitValue: FromUnitValue{
				DefaultFrictional: &DefaultFrictional{
					value: udfs("17.3475345").Mul(udfs("5")).Mul(udfs("1.1")),
				},
			},
			Tax: Tax{
				ratio:   udfs("10"),
				amount:  udfs("17.3475345").Mul(udfs("0.1")).Mul(udfs("5")),
				taxable: udfs("17.3475345").Mul(udfs("5")),
			},
		},
	},
	{
		name: "Tax over entry unit value with multiple taxes",
		tester: func() (*FromUnitValue, Tax, bool, string, error) {
			entry, _ := alpacadecimal.NewFromString("100")
			b := NewFromUnitValue(entry)
			ratio1 := udfs("10")
			t1 := NewPercTax(ratio1)

			ratio2 := udfs("5")
			t2 := NewPercTax(ratio2)

			b.Bind(t1)
			// this tax will be applied over the previous tax t1
			b.Bind(t2)

			return b, t2.Tax, false, "", nil
		},
		expecteds: struct {
			FromUnitValue
			Tax
		}{
			FromUnitValue: FromUnitValue{
				DefaultFrictional: &DefaultFrictional{
					value: udfs("100").Mul(udfs("1.1")).Mul(udfs("1.05")),
				},
			},
			Tax: Tax{
				ratio:   udfs("5"),
				amount:  udfs("100").Mul(udfs("1.1")).Mul(udfs("0.05")),
				taxable: udfs("100").Mul(udfs("1.1")),
			},
		},
	},
}
