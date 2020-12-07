package calculator

const YEAR = 31622400

func (s *State) applyNewTaxEvent(t TaxEvent) {
	if t.BoughtAsset == "USD" {
		a := s.getAsset(t.SoldAsset)
		r := a.applyNewSell(t.SoldQty, (t.BoughtQty / t.SoldQty), t.Timestamp)
		s.applyReceipt(r)
	} else if t.SoldAsset == "USD" {
		a := s.getAsset(t.BoughtAsset)
		a.applyNewBuy(t.BoughtQty, (t.SoldQty / t.BoughtQty), t.Timestamp)
	} else {
		a1 := s.getAsset(t.BoughtAsset)
		a1.applyNewBuy(t.BoughtQty, t.BoughtPriceUSD, t.Timestamp)
		a2 := s.getAsset(t.SoldAsset)
		r := a2.applyNewSell(t.SoldQty, t.SoldPriceUSD, t.Timestamp)
		s.applyReceipt(r)
	}
}

func (s *State) getAsset(name string) *Asset {
	if s.Assets[name] == nil {
		a := Asset{Name: name}
		s.Assets[name] = &a
		return &a
	} else {
		return s.Assets[name]
	}
}

func (s *State) applyReceipt(r Receipt) {
	s.TotalShortGains += r.ShortTermGains
	s.TotalLongGains += r.LongTermGains
	s.TotalShortLosses += r.ShortTermLosses
	s.TotalLongLosses += r.LongTermLosses
}

func (a *Asset) applyNewBuy(amt float32, price float32, date uint32) {
	a.Balance += amt
	cb := CostBasis{amt, price, date}
	a.CostBases = append(a.CostBases, cb)
}

func (a *Asset) applyNewSell(amt float32, price float32, date uint32) Receipt {
	a.Balance -= amt
	r := Receipt{}
	for amt > 0 {
		if len(a.CostBases) > 0 {
			lastCB := a.CostBases[len(a.CostBases)-1]
			a.CostBases = a.CostBases[:len(a.CostBases)-1]
			dateDiff := (date - lastCB.Date)
			isShort := YEAR > dateDiff
			// fmt.Printf("last CB: %v | Date diff: %v | Short: %v\n", lastCB, dateDiff, isShort)
			if amt >= lastCB.Amount {
				r.calculateGains(price, lastCB.USDPrice, lastCB.Amount, isShort)
				amt -= lastCB.Amount
			} else if lastCB.Amount > amt {
				r.calculateGains(price, lastCB.USDPrice, amt, isShort)
				lastCB.Amount -= amt
				amt = 0
				a.CostBases = append(a.CostBases, lastCB)
			}
		}
	}
	// r.PrintReceipt()
	return r
}

func (r *Receipt) calculateGains(price, lastPrice, amount float32, isShort bool) {
	if price >= lastPrice {
		gain := (price * amount) - (lastPrice * amount)
		if isShort {
			r.ShortTermGains += gain
		} else {
			r.LongTermGains += gain
		}
	}
	if price < lastPrice {
		loss := (lastPrice * amount) - (price * amount)
		if isShort {
			r.ShortTermLosses += loss
		} else {
			r.LongTermLosses += loss
		}
	}
}
