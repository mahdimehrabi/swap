package crypto_compare

type Raw struct {
	Market                string  `json:"MARKET"`
	FromSymbol            string  `json:"FROMSYMBOL"`
	ToSymbol              string  `json:"TOSYMBOL"`
	Flags                 int     `json:"FLAGS"`
	Price                 float64 `json:"PRICE"`
	LastUpdate            int64   `json:"LASTUPDATE"`
	LastVolume            float64 `json:"LASTVOLUME"`
	LastVolumeTo          float64 `json:"LASTVOLUMETO"`
	LastTradeID           string  `json:"LASTTRADEID"`
	Volume24Hour          float64 `json:"VOLUME24HOUR"`
	Volume24HourTo        float64 `json:"VOLUME24HOURTO"`
	Open24Hour            float64 `json:"OPEN24HOUR"`
	High24Hour            float64 `json:"HIGH24HOUR"`
	Low24Hour             float64 `json:"LOW24HOUR"`
	LastMarket            string  `json:"LASTMARKET"`
	TopTierVolume24Hour   float64 `json:"TOPTIERVOLUME24HOUR"`
	TopTierVolume24HourTo float64 `json:"TOPTIERVOLUME24HOURTO"`
	Change24Hour          float64 `json:"CHANGE24HOUR"`
	ChangePct24Hour       float64 `json:"CHANGEPCT24HOUR"`
	ChangeDay             float64 `json:"CHANGEDAY"`
	ChangePctDay          float64 `json:"CHANGEPCTDAY"`
	ChangeHour            float64 `json:"CHANGEHOUR"`
	ChangePctHour         float64 `json:"CHANGEPCTHOUR"`
}

type Response struct {
	Raw *Raw `json:"RAW"`
}
