package frictional

import "testing"

func TestFromBrute(t *testing.T) {
	for i, tc := range testCasesFromBrute {
		brute, net, netWD, unitValue, bg := tc.tester()

		if !brute.buffer.Equal(tc.expected.brute.buffer) {
			t.Logf("[test case %d] got brute %v expected %v", i, brute.buffer, tc.expected.brute.buffer)
			t.FailNow()
		}

		if !net.buffer.Equal(tc.expected.net.buffer) {
			t.Logf("[test case %d] got net %v expected %v", i, net.buffer, tc.expected.net.buffer)
			t.FailNow()
		}

		if !netWD.buffer.Equal(tc.expected.netWD.buffer) {
			t.Logf("[test case %d] got netWD %v expected %v", i, netWD.buffer, tc.expected.netWD.buffer)
			t.FailNow()
		}

		if !unitValue.unitValue.Equal(tc.expected.unitValue.unitValue) {
			t.Logf("[test case %d] got unitValue %v expected %v", i, unitValue.unitValue, tc.expected.unitValue.unitValue)
			t.FailNow()
		}

		if !bg.value.Equal(tc.expected.bg.value) {
			t.Logf("[test case %d] got buffer %v expected %v", i, bg.value, tc.expected.bg.value)
			t.FailNow()
		}
	}
}

var testCasesFromBrute = []struct {
	tester   func() (brute *SnapshotVisitor, net *SnapshotVisitor, netWD *SnapshotVisitor, unitValue *UnitValue, bg *FromBrute)
	expected struct {
		brute     *SnapshotVisitor
		net       *SnapshotVisitor
		netWD     *SnapshotVisitor
		unitValue *UnitValue
		bg        *FromBrute
	}
}{
	{
		tester: func() (brute *SnapshotVisitor, net *SnapshotVisitor, netWD *SnapshotVisitor, unitValue *UnitValue, bg *FromBrute) {

			brute = NewSnapshot()
			net = NewSnapshot()
			netWD = NewSnapshot()

			bg = NewFromBruteDefault().WithBrute(udfs("1619.1"))

			disc := NewPercentualUnDiscount(udfs("0"))
			tx := NewPercentualUnTax(udfs("16"))

			unitValue = NewUnitValue(udfs("3"))

			bg.Bind(brute)
			bg.Bind(tx)
			bg.Bind(net)
			bg.Bind(disc)
			bg.Bind(netWD)
			bg.Bind(unitValue)
			bg.Bind(NewRound(12))

			unitValue.unitValue = unitValue.unitValue.Round(12)

			return brute, net, netWD, unitValue, bg
		},
		expected: struct {
			brute     *SnapshotVisitor
			net       *SnapshotVisitor
			netWD     *SnapshotVisitor
			unitValue *UnitValue
			bg        *FromBrute
		}{
			brute: &SnapshotVisitor{
				buffer: udfs("1619.1"),
			},
			net: &SnapshotVisitor{
				buffer: udfs("1395.7758620689655172"),
			},
			netWD: &SnapshotVisitor{
				buffer: udfs("1395.7758620689655172"),
			},
			unitValue: &UnitValue{
				qty:       udfs("3"),
				unitValue: udfs("465.258620689655"),
			},
			bg: &FromBrute{
				DefaultFrictional: &DefaultFrictional{
					value: udfs("465.258620689655"),
				},
			},
		},
	},
	{
		tester: func() (brute *SnapshotVisitor, net *SnapshotVisitor, netWD *SnapshotVisitor, unitValue *UnitValue, bg *FromBrute) {
			brute = &SnapshotVisitor{}
			net = &SnapshotVisitor{}
			netWD = &SnapshotVisitor{}

			bg = NewFromBruteDefault().WithBrute(udfs("1000"))

			disc := NewPercentualUnDiscount(udfs("10"))
			tx := NewPercentualUnTax(udfs("20"))

			unitValue = NewUnitValue(udfs("5"))

			bg.Bind(brute)
			bg.Bind(tx)
			bg.Bind(net)
			bg.Bind(disc)
			bg.Bind(netWD)
			bg.Bind(unitValue)
			bg.Bind(NewRound(2))

			unitValue.unitValue = unitValue.unitValue.Round(2)

			return brute, net, netWD, unitValue, bg
		},
		expected: struct {
			brute     *SnapshotVisitor
			net       *SnapshotVisitor
			netWD     *SnapshotVisitor
			unitValue *UnitValue
			bg        *FromBrute
		}{
			brute: &SnapshotVisitor{
				buffer: udfs("1000"),
			},
			net: &SnapshotVisitor{
				buffer: udfs("833.3333333333333333"),
			},
			netWD: &SnapshotVisitor{
				buffer: udfs("925.92592592592593"),
			},
			unitValue: &UnitValue{
				qty:       udfs("5"),
				unitValue: udfs("185.19"),
			},
			bg: &FromBrute{
				DefaultFrictional: &DefaultFrictional{
					value: udfs("185.19"),
				},
			},
		},
	},
}

func BenchmarkFromBrute(b *testing.B) {
	brute := &SnapshotVisitor{}
	net := &SnapshotVisitor{}
	netWD := &SnapshotVisitor{}
	tree := udfs("3")
	sixteennineteenone := udfs("1619.1")
	zero := udfs("0")
	sixteen := udfs("16")

	unitValue := NewUnitValue(tree)
	bg := NewFromBruteDefault().WithBrute(sixteennineteenone)
	disc := NewPercentualUnDiscount(zero)
	tx := NewPercentualUnTax(sixteen)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bg.Bind(brute)
		bg.Bind(tx)
		bg.Bind(net)
		bg.Bind(disc)
		bg.Bind(netWD)
		bg.Bind(unitValue)

		b.StopTimer()
		bg.Restore(brute.Get().Copy())
		b.StartTimer()
	}
}
