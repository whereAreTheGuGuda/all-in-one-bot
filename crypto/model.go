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
	DexId       string       `json:"dexId"`
	PairAddress string       `json:"pairAddress"`
	PriceNative string       `json:"priceNative"`
	PriceUsd    string       `json:"priceUsd"`
	BaseToken   *BaseToken   `json:"baseToken"`
	PriceChange *PriceChange `json:"priceChange"`
	Volume      *Volume      `json:"volume"`
	Txns        *Txn         `json:"txns"`
	Lp          *Liquidity   `json:"liquidity"`
	Fdv         int64        `json:"fdv"`
	CreateAt    int64        `json:"pairCreatedAt"`
	CreateTime  string
}

type Liquidity struct {
	Usd float64 `json:"usd"`
}

type Volume struct {
	M5  float64 `json:"m5"`
	H1  float64 `json:"h1"`
	H6  float64 `json:"h6"`
	H24 float64 `json:"h24"`
}

type Txn struct {
	M5  *BS `json:"m5"`
	H1  *BS `json:"h1"`
	H6  *BS `json:"h6"`
	H24 *BS `json:"h24"`
}

type BS struct {
	B int `json:"buys"`
	S int `json:"sells"`
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
	Code         int64                   `json:"code"`
	Message      string                  `json:"message"`
	MemeCheckers map[string]*MemeChecker `json:"result"`
}
type MemeChecker struct {
	AntiWhaleModifiable  string `json:"anti_whale_modifiable"`
	BuyTax               string `json:"buy_tax"`
	CanTakeBackOwnership string `json:"can_take_back_ownership"`
	//CannotBuy                  string   `json:"cannot_buy"`
	CannotSellAll  string `json:"cannot_sell_all"`
	CreatorAddress string `json:"creator_address"`
	CreatorBalance string `json:"creator_balance"`
	CreatorPercent string `json:"creator_percent"`
	//Dex                        []Dex    `json:"dex"`
	ExternalCall string `json:"external_call"`
	HiddenOwner  string `json:"hidden_owner"`
	HolderCount  string `json:"holder_count"`
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
	LpLockedTotal              float64
}

type Dex struct {
	Name      string `json:"name"`
	Liquidity string `json:"liquidity"`
	Pair      string `json:"pair"`
}

type Holder struct {
	Address    string `json:"address"`
	Tag        string `json:"tag"`
	IsContract int64  `json:"is_contract"`
	Balance    string `json:"balance"`
	Percent    string `json:"percent"`
	IsLocked   int64  `json:"is_locked"`
	//LockedDetail []LockedDetail `json:"locked_detail"`
}

type LockedDetail struct {
	Amount  string `json:"amount"`
	EndTime string `json:"end_time"`
	OptTime string `json:"opt_time"`
}

// Generated by https://quicktype.io

type TokenTxResp struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Result  []TokenTx `json:"result"`
}

type TokenTx struct {
	TokenName       string `json:"tokenName"`
	TokenSymbol     string `json:"tokenSymbol"`
	ContractAddress string `json:"contractAddress"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	Decimal         string `json:"tokenDecimal"`
	TimeStamp       string `json:"timeStamp"`
	Block           string `json:"blockNumber"`
}

type HoneypotResp struct {
	Honeypot Honeypot `json:"honeypotResult"`
}

type Honeypot struct {
	Is bool `json:"isHoneypot"`
}

type ContractCreationResp struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Result  []ContractCreation `json:"result"`
}

type ContractCreation struct {
	Address string `json:"contractAddress"`
	Creator string `json:"contractCreator"`
}

type GasOracleResp struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Result  GasOracle `json:"result"`
}

type GasOracle struct {
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
}

// DextoolsResp Start
type DextoolsResp struct {
	Code string  `json:"code"`
	Data []Datum `json:"data"`
}

type Datum struct {
	ID            DatumID       `json:"id"`
	CreationBlock int64         `json:"creationBlock"`
	CreationTime  string        `json:"creationTime"`
	DextScore     DextScore     `json:"dextScore"`
	Metrics       DatumMetrics  `json:"metrics"`
	Name          string        `json:"name"`
	NameRef       string        `json:"nameRef"`
	Symbol        string        `json:"symbol"`
	SymbolRef     string        `json:"symbolRef"`
	Type          string        `json:"type"`
	Locks         []interface{} `json:"locks"`
	Votes         Votes         `json:"votes"`
	Token         Token         `json:"token"`
	Price         float64       `json:"price"`
	Price5M       Price         `json:"price5m"`
	Price1H       Price         `json:"price1h"`
	Price6H       Price         `json:"price6h"`
	Price24H      Price         `json:"price24h"`
}

type DextScore struct {
	Information  int64 `json:"information"`
	Holders      int64 `json:"holders"`
	Pool         int64 `json:"pool"`
	Transactions int64 `json:"transactions"`
	Creation     int64 `json:"creation"`
	Total        int64 `json:"total"`
}

type DatumID struct {
	Chain    string `json:"chain"`
	Exchange string `json:"exchange"`
	Pair     string `json:"pair"`
	Token    string `json:"token"`
	TokenRef string `json:"tokenRef"`
}

type DatumMetrics struct {
	Liquidity          float64 `json:"liquidity"`
	InitialLiquidity   float64 `json:"initialLiquidity"`
	LiquidityUpdatedAt string  `json:"liquidityUpdatedAt"`
	Reserve            float64 `json:"reserve"`
	ReserveRef         float64 `json:"reserveRef"`
}

type Price struct {
	Volume     float64 `json:"volume"`
	Swaps      int64   `json:"swaps"`
	Price      float64 `json:"price"`
	PriceChain float64 `json:"priceChain"`
	Buys       int64   `json:"buys"`
	Sells      int64   `json:"sells"`
}

type Token struct {
	ID            TokenID           `json:"id"`
	Audit         Audit             `json:"audit"`
	Decimals      int64             `json:"decimals"`
	Info          Info              `json:"info"`
	Links         map[string]string `json:"links"`
	Logo          string            `json:"logo"`
	Metrics       TokenMetrics      `json:"metrics"`
	Name          string            `json:"name"`
	Symbol        string            `json:"symbol"`
	TotalSupply   string            `json:"totalSupply"`
	ReprPair      ReprPair          `json:"reprPair"`
	CreationBlock int64             `json:"creationBlock"`
	CreationTime  string            `json:"creationTime"`
}

type Audit struct {
	IsContractRenounced bool   `json:"is_contract_renounced"`
	CodeVerified        bool   `json:"codeVerified"`
	Date                string `json:"date"`
	LockTransactions    bool   `json:"lockTransactions"`
	Mint                bool   `json:"mint"`
	Provider            string `json:"provider"`
	Proxy               bool   `json:"proxy"`
	Status              string `json:"status"`
	UnlimitedFees       bool   `json:"unlimitedFees"`
	Version             int64  `json:"version"`
}

type TokenID struct {
	Chain string `json:"chain"`
	Token string `json:"token"`
}

type Info struct {
	Cmc           string `json:"cmc"`
	Coingecko     string `json:"coingecko"`
	Description   string `json:"description"`
	Dextools      bool   `json:"dextools"`
	Email         string `json:"email"`
	ExtraInfo     string `json:"extraInfo"`
	NftCollection string `json:"nftCollection"`
	Ventures      bool   `json:"ventures"`
}

type TokenMetrics struct {
	MaxSupply        float64     `json:"maxSupply"`
	TotalSupply      float64     `json:"totalSupply"`
	Holders          int64       `json:"holders"`
	HoldersUpdatedAt string      `json:"holdersUpdatedAt"`
	TxCount          int64       `json:"txCount"`
	Mcap             interface{} `json:"mcap"`
	Fdv              float64     `json:"fdv"`
}

type ReprPair struct {
	ID        DatumID `json:"id"`
	UpdatedAt string  `json:"updatedAt"`
}

type Votes struct {
	Warning   int64 `json:"_warning"`
	Downvotes int64 `json:"downvotes"`
	Upvotes   int64 `json:"upvotes"`
}

// DextoolsResp End
