package calculator

// type State map[string]*Asset

type State struct {
	TotalShortGains  float32
	TotalLongGains   float32
	TotalShortLosses float32
	TotalLongLosses  float32
	Assets           map[string]*Asset
}

type Receipt struct {
	ShortTermGains  float32
	ShortTermLosses float32
	LongTermGains   float32
	LongTermLosses  float32
}

type TaxEvent struct {
	BoughtAsset    string
	BoughtQty      float32
	SoldAsset      string
	SoldQty        float32
	BoughtPriceUSD float32
	SoldPriceUSD   float32
	Timestamp      uint32
}

type Asset struct {
	Name      string
	Balance   float32
	CostBases []CostBasis
}

type CostBasis struct {
	Amount   float32
	USDPrice float32
	Date     uint32 //unix time
}
