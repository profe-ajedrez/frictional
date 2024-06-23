package frictional

import "testing"

func TestHandlerFromUnitValue(t *testing.T) {
	for i, tc := range testCaseTaxHandlerFromUnitValue {
		b, th, shouldFail, err := tc.tester()

		if !shouldFail && err != nil {
			t.Logf("[test %d FAILED] should not fail. %v", i, err)
			t.FailNow()
		}

		if shouldFail && err == nil {
			t.Logf("[test %d FAILED] should has been failed.", i)
			t.FailNow()
		}

		if shouldFail {
			continue
		}

		if !b.Buffer().Equal(tc.expected.Buffer()) {
			t.Logf("[test %d FAILED] buffer. Got %v  Expected %v", i, b.Buffer(), tc.expected.Buffer())
			t.FailNow()
		}

		if !th.totalRatio.Equal(tc.expected.totalRatio) {
			t.Logf("[test %d FAILED] total ratio. Got %v  Expected %v", i, th.totalRatio, tc.expected.totalRatio)
			t.FailNow()
		}

		if !th.totalAmount.Equal(tc.expected.totalAmount) {
			t.Logf("[test %d FAILED] total amount. Got %v  Expected %v", i, th.totalAmount, tc.expected.totalAmount)
			t.FailNow()
		}

		if !th.taxable.Equal(tc.expected.taxable) {
			t.Logf("[test %d FAILED] taxable. Got %v  Expected %v", i, th.taxable, tc.expected.taxable)
			t.FailNow()
		}
	}
}

func BenchmarkTaxHandlerFromUnitValue(b *testing.B) {
	entry := udfs("232.5")
	qty := udfs("3")
	tax := udfs("16")

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		bg := NewFromUnitValue(entry)
		th := NewTaxHandlerFromUnitValue()
		th.WithPercentualTax(tax)

		qtyEv := WithQTY(qty)

		bg.Bind(qtyEv)
		bg.Snapshot() // snapshot of the net value
		th.Do(bg)
		bg.Snapshot() // snapshot of the brute value
	}
}

var testCaseTaxHandlerFromUnitValue = []struct {
	tester   func() (Frictional, *TaxHandlerFromUnitValue, bool, error)
	expected struct {
		Frictional
		*TaxHandlerFromUnitValue
	}
}{
	{
		tester: func() (Frictional, *TaxHandlerFromUnitValue, bool, error) {
			entry := udfs("232.5")
			b := NewFromUnitValue(entry)

			b.Bind(WithQTY(udfs("3")))

			th := NewTaxHandlerFromUnitValue()
			th.WithPercentualTax(udfs("16"))
			net := SnapshotVisitor{}
			net.Do(b)
			th.Do(b)
			brute := SnapshotVisitor{}
			brute.Do(b)

			return b, th, false, nil
		},
		expected: struct {
			Frictional
			*TaxHandlerFromUnitValue
		}{
			Frictional: &FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("232.5").Mul(udfs("3")).Mul(udfs("1.16")),
				},
			},
			TaxHandlerFromUnitValue: &TaxHandlerFromUnitValue{
				taxHandler: &taxHandler{
					totalRatio:  udfs("16"),
					totalAmount: udfs("232.5").Mul(udfs("0.16")).Mul(udfs("3")),
					taxable:     udfs("232.5").Mul(udfs("3")),
				},
			},
		},
	},
	{
		tester: func() (Frictional, *TaxHandlerFromUnitValue, bool, error) {
			entry := udfs("100")
			b := NewFromUnitValue(entry)

			b.Bind(WithQTY(udfs("2")))

			th := NewTaxHandlerFromUnitValue()
			th.WithPercentualTax(udfs("20"))

			net := SnapshotVisitor{}
			net.Do(b)

			th.Do(b)

			brute := SnapshotVisitor{}
			brute.Do(b)

			return b, th, false, nil
		},
		expected: struct {
			Frictional
			*TaxHandlerFromUnitValue
		}{
			Frictional: &FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("100").Mul(udfs("2")).Mul(udfs("1.2")),
				},
			},
			TaxHandlerFromUnitValue: &TaxHandlerFromUnitValue{
				taxHandler: &taxHandler{
					totalRatio:  udfs("20"),
					totalAmount: udfs("100").Mul(udfs("0.2")).Mul(udfs("2")),
					taxable:     udfs("100").Mul(udfs("2")),
				},
			},
		},
	},
	{
		tester: func() (Frictional, *TaxHandlerFromUnitValue, bool, error) {
			entry := udfs("50")
			b := NewFromUnitValue(entry)

			b.Bind(WithQTY(udfs("4")))

			th := NewTaxHandlerFromUnitValue()
			th.WithPercentualTax(udfs("8"))

			net := SnapshotVisitor{}
			net.Do(b)

			th.Do(b)

			brute := SnapshotVisitor{}
			brute.Do(b)

			return b, th, false, nil
		},
		expected: struct {
			Frictional
			*TaxHandlerFromUnitValue
		}{
			Frictional: &FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("50").Mul(udfs("4")).Mul(udfs("1.08")),
				},
			},
			TaxHandlerFromUnitValue: &TaxHandlerFromUnitValue{
				taxHandler: &taxHandler{
					totalRatio:  udfs("8"),
					totalAmount: udfs("50").Mul(udfs("0.08")).Mul(udfs("4")),
					taxable:     udfs("50").Mul(udfs("4")),
				},
			},
		},
	},
	{
		tester: func() (Frictional, *TaxHandlerFromUnitValue, bool, error) {
			entry := udfs("75.25")
			b := NewFromUnitValue(entry)

			b.Bind(WithQTY(udfs("1.5")))

			th := NewTaxHandlerFromUnitValue()
			th.WithPercentualTax(udfs("12.5"))

			net := SnapshotVisitor{}
			net.Do(b)

			th.Do(b)

			brute := SnapshotVisitor{}
			brute.Do(b)

			return b, th, false, nil
		},
		expected: struct {
			Frictional
			*TaxHandlerFromUnitValue
		}{
			Frictional: &FromUnitValue{
				defaultFrictional: &defaultFrictional{
					buffer: udfs("75.25").Mul(udfs("1.5")).Mul(udfs("1.125")),
				},
			},
			TaxHandlerFromUnitValue: &TaxHandlerFromUnitValue{
				taxHandler: &taxHandler{
					totalRatio:  udfs("12.5"),
					totalAmount: udfs("75.25").Mul(udfs("0.125")).Mul(udfs("1.5")),
					taxable:     udfs("75.25").Mul(udfs("1.5")),
				},
			},
		},
	},
}
