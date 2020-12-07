package calculator

import "testing"

func TestGains(t *testing.T) {
	asset := Asset{Name: "TEST"}
	asset.applyNewBuy(100, 2.50, 1506433889)    // buy 100 TEST @ $2.50 each
	r := asset.applyNewSell(30, 10, 1506433999) // sell 30 @ $10 each
	t.Logf("%v", r)

	if r.ShortTermGains != 225 {
		t.Errorf("Expected Gains = 225, Got: %v\n", r.ShortTermGains)
	}

	t.Logf("Asset: %v | Balance: %v \n", asset.Name, asset.Balance)
	t.Logf("%v", asset.CostBases)

	asset.applyNewBuy(20, 20, 1538057288) // buy 20 @ $20 each

	t.Logf("Asset: %v | Balance: %v \n", asset.Name, asset.Balance)
	t.Logf("%v", asset.CostBases)

	r = asset.applyNewSell(30, 40, 1538057399) // sell 30 @ $40 each
	t.Logf("%v", r)

	t.Logf("Asset: %v | Balance: %v \n", asset.Name, asset.Balance)
	t.Logf("%v", asset.CostBases)
}
