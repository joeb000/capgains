package calculator

import "fmt"

func (te *TaxEvent) PrintToLine() {
	fmt.Printf("Tax Event: %v\n", te)
}

func (r *Receipt) PrintReceipt() {
	fmt.Printf("*** New Sell Receipt ***\n")
	fmt.Printf("STG: %v | LTG: %v\n", r.ShortTermGains, r.LongTermGains)
	fmt.Printf("STL: %v | LTL: %v\n*******\n", r.ShortTermLosses, r.LongTermLosses)
}

func (s *State) PrintGainsAndLosses() {
	fmt.Printf("********** State **********\n")
	fmt.Printf("Short Term Gains:  %.2f\n", s.TotalShortGains)
	fmt.Printf("Long Term Gains:   %.2f\n", s.TotalLongGains)
	fmt.Printf("Short Term Losses: %.2f\n", s.TotalShortLosses)
	fmt.Printf("Long Term Losses:  %.2f\n", s.TotalLongLosses)
	fmt.Printf("***************************\n")

}
