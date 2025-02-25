package currency

import (
	"sync"

	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)

// PairsManager manages asset pairs
type PairsManager struct {
	BypassConfigFormatUpgrades bool         `json:"bypassConfigFormatUpgrades"`
	RequestFormat              *PairFormat  `json:"requestFormat,omitempty"`
	ConfigFormat               *PairFormat  `json:"configFormat,omitempty"`
	UseGlobalFormat            bool         `json:"useGlobalFormat,omitempty"`
	LastUpdated                int64        `json:"lastUpdated,omitempty"`
	Pairs                      FullStore    `json:"pairs"`
	mutex                      sync.RWMutex `json:"-"`
}

// FullStore holds all supported asset types with the enabled and available
// pairs for an exchange.
type FullStore map[asset.Item]*PairStore

// PairStore stores a currency pair store
type PairStore struct {
	AssetEnabled  *bool       `json:"assetEnabled"`
	Enabled       Pairs       `json:"enabled"`
	Available     Pairs       `json:"available"`
	RequestFormat *PairFormat `json:"requestFormat,omitempty"`
	ConfigFormat  *PairFormat `json:"configFormat,omitempty"`
}

// PairFormat returns the pair format
type PairFormat struct {
	Uppercase bool   `json:"uppercase"`
	Delimiter string `json:"delimiter,omitempty"`
	Separator string `json:"separator,omitempty"`
	Index     string `json:"index,omitempty"`
}
