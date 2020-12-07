package calculator

import "testing"

func TestState(t *testing.T) {
	myState := State{}
	assets := make(map[string]*Asset)
	myState.Assets = assets

	t1 := TaxEvent{"XXX", 5, "USD", 20, 4, 1, 1546300800}  // Bought 5 XXX for $20 ($4 each) on 1/1/2019
	t2 := TaxEvent{"XXX", 2, "USD", 60, 30, 1, 1559779200} // Bought 2 XXX for $60 ($30 each) on 6/6/2019
	t3 := TaxEvent{"USD", 80, "XXX", 4, 1, 20, 1580515200} // Sold 4 XXX for $80 ($20 each) on 2/1/2020

	myState.applyNewTaxEvent(t1)
	myState.applyNewTaxEvent(t2)
	myState.applyNewTaxEvent(t3)

	// expect STL of $20 and LTG of $32
	if myState.TotalShortLosses != 20 {
		t.Errorf("Short Term Loss Expected: 20 | Got: %v", myState.TotalShortLosses)
	}
	if myState.TotalLongGains != 32 {
		t.Errorf("Long Term Gain Expected: 32 | Got: %v", myState.TotalLongGains)
	}
}
