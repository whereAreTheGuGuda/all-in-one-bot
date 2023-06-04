package crypto

type priceResp struct {
	Price  string `json:"price"`
	Symbol string `json:"symbol"`
}

type Meme struct {
	Pairs []*Pair `json:"pairs"`
}

type Pair struct {
	URL         string       `json:"url"`
	ChainId     string       `json:"chainId"`
	PairAddress string       `json:"pairAddress"`
	PriceNative string       `json:"priceNative"`
	PriceUsd    string       `json:"priceUsd"`
	BaseToken   *BaseToken   `json:"baseToken"`
	PriceChange *PriceChange `json:"priceChange"`
}

type BaseToken struct {
	Addr   string `json:"address"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type PriceChange struct {
	M5  float64 `json:"m5"`
	H1  float64 `json:"h1"`
	H6  float64 `json:"h6"`
	H24 float64 `json:"h24"`
}

type MemeCheckerResp struct {
	Code    int64  `json:"code"`   
	Message string `json:"message"`
	MemeCheckers  map[string]*MemeChecker `json:"result"` 
}
type MemeChecker struct {
	AntiWhaleModifiable        string   `json:"anti_whale_modifiable"`       
	BuyTax                     string   `json:"buy_tax"`                     
	CanTakeBackOwnership       string   `json:"can_take_back_ownership"`     
	CannotBuy                  string   `json:"cannot_buy"`                  
	CannotSellAll              string   `json:"cannot_sell_all"`             
	CreatorAddress             string   `json:"creator_address"`             
	CreatorBalance             string   `json:"creator_balance"`             
	CreatorPercent             string   `json:"creator_percent"`             
	//Dex                        []Dex    `json:"dex"`                         
	ExternalCall               string   `json:"external_call"`               
	HiddenOwner                string   `json:"hidden_owner"`                
	HolderCount                string   `json:"holder_count"`                
	//Holders                    []Holder `json:"holders"`                     
	HoneypotWithSameCreator    string   `json:"honeypot_with_same_creator"`  
	IsAntiWhale                string   `json:"is_anti_whale"`               
	IsBlacklisted              string   `json:"is_blacklisted"`              
	IsHoneypot                 string   `json:"is_honeypot"`                 
	IsInDex                    string   `json:"is_in_dex"`                   
	IsMintable                 string   `json:"is_mintable"`                 
	IsOpenSource               string   `json:"is_open_source"`              
	IsProxy                    string   `json:"is_proxy"`                    
	IsWhitelisted              string   `json:"is_whitelisted"`              
	LpHolderCount              string   `json:"lp_holder_count"`             
	LpHolders                  []Holder `json:"lp_holders"`                  
	LpTotalSupply              string   `json:"lp_total_supply"`             
	OwnerAddress               string   `json:"owner_address"`               
	OwnerBalance               string   `json:"owner_balance"`               
	OwnerChangeBalance         string   `json:"owner_change_balance"`        
	OwnerPercent               string   `json:"owner_percent"`               
	PersonalSlippageModifiable string   `json:"personal_slippage_modifiable"`
	Selfdestruct               string   `json:"selfdestruct"`                
	SellTax                    string   `json:"sell_tax"`                    
	SlippageModifiable         string   `json:"slippage_modifiable"`         
	TokenName                  string   `json:"token_name"`                  
	TokenSymbol                string   `json:"token_symbol"`                
	TotalSupply                string   `json:"total_supply"`                
	TradingCooldown            string   `json:"trading_cooldown"`            
	TransferPausable           string   `json:"transfer_pausable"`           
}

type Dex struct {
	Name      string   `json:"name"`     
	Liquidity string `json:"liquidity"`
	Pair      string `json:"pair"`     
}

type Holder struct {
	Address      string         `json:"address"`      
	Tag          string           `json:"tag"`          
	IsContract   int64          `json:"is_contract"`  
	Balance      string         `json:"balance"`      
	Percent      string         `json:"percent"`      
	IsLocked     int64          `json:"is_locked"`    
	LockedDetail []LockedDetail `json:"locked_detail"`
}

type LockedDetail struct {
	Amount  string `json:"amount"`  
	EndTime string `json:"end_time"`
	OptTime string `json:"opt_time"`
}
