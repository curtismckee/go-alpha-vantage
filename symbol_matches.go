package av

type SymbolMatches struct {
	BestMatches []*SymbolMatch `json:"bestMatches"`
}

type SymbolMatch struct {
	Symbol      string  `json:"1. symbol"`
	Name        string  `json:"2. name"`
	Type        string  `json:"3. type"`
	Region      string  `json:"4. region"`
	MarketOpen  string  `json:"5. marketOpen"`
	MarketClose string  `json:"6. marketClose"`
	Timezone    string  `json:"7. timezone"`
	Currency    string  `json:"8. currency"`
	MatchScore  float64 `json:"9. matchScore,string"`
}
