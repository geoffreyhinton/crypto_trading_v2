package models

import (
	"time"

	"gorm.io/gorm"
)

// Base CryptoAddress model
type CryptoAddress struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null;index"`
	Address     string `json:"address" gorm:"unique;not null"`
	PublicKey   string `json:"public_key,omitempty"`
	PrivateKey  string `json:"-" gorm:"column:private_key"` // Never serialize to JSON
	Network     string `json:"network" gorm:"not null"`     // mainnet, testnet, regtest
	CryptoType  string `json:"crypto_type" gorm:"not null"` // bitcoin, ethereum, etc.
	Label       string `json:"label,omitempty"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	Balance     string `json:"balance" gorm:"type:decimal(28,18);default:0"` // Use string for precision
	LastSyncAt  *time.Time `json:"last_sync_at,omitempty"`
	
	// Relationships
	User         User                `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Deposits     []CryptoDeposit     `json:"deposits,omitempty" gorm:"foreignKey:AddressID"`
	Withdrawals  []CryptoWithdrawal  `json:"withdrawals,omitempty" gorm:"foreignKey:FromAddressID"`
	Transactions []CryptoTransaction `json:"transactions,omitempty" gorm:"foreignKey:AddressID"`
}

// BitcoinAddress extends CryptoAddress
type BitcoinAddress struct {
	CryptoAddress
	AddressType   string `json:"address_type"` // P2PKH, P2SH, P2WPKH, P2WSH, P2TR
	ScriptPubKey  string `json:"script_pub_key,omitempty"`
	RedeemScript  string `json:"redeem_script,omitempty"`
	WitnessScript string `json:"witness_script,omitempty"`
	Derivation    string `json:"derivation,omitempty"` // HD wallet derivation path
}

// EthereumAddress extends CryptoAddress
type EthereumAddress struct {
	CryptoAddress
	Nonce        uint64 `json:"nonce" gorm:"default:0"`
	GasPrice     string `json:"gas_price,omitempty" gorm:"type:decimal(28,0)"`
	GasLimit     uint64 `json:"gas_limit" gorm:"default:21000"`
	ChainID      uint64 `json:"chain_id" gorm:"default:1"` // 1 for mainnet, 3 for ropsten, etc.
	IsContract   bool   `json:"is_contract" gorm:"default:false"`
	ContractABI  string `json:"contract_abi,omitempty" gorm:"type:text"`
	TokenBalance string `json:"token_balance,omitempty" gorm:"type:decimal(28,18);default:0"`
}

// Base CryptoTransaction model
type CryptoTransaction struct {
	gorm.Model
	AddressID        uint      `json:"address_id" gorm:"not null;index"`
	TxHash           string    `json:"tx_hash" gorm:"unique;not null"`
	BlockHash        string    `json:"block_hash,omitempty"`
	BlockHeight      uint64    `json:"block_height,omitempty"`
	BlockTime        *time.Time `json:"block_time,omitempty"`
	FromAddress      string    `json:"from_address"`
	ToAddress        string    `json:"to_address"`
	Amount           string    `json:"amount" gorm:"type:decimal(28,18);not null"`
	Fee              string    `json:"fee" gorm:"type:decimal(28,18);default:0"`
	Status           string    `json:"status" gorm:"default:pending"` // pending, confirmed, failed
	Confirmations    uint      `json:"confirmations" gorm:"default:0"`
	Network          string    `json:"network" gorm:"not null"`
	CryptoType       string    `json:"crypto_type" gorm:"not null"`
	Direction        string    `json:"direction" gorm:"not null"` // incoming, outgoing, internal
	RawTransaction   string    `json:"raw_transaction,omitempty" gorm:"type:text"`
	Memo             string    `json:"memo,omitempty"`
	
	// Relationships
	Address CryptoAddress `json:"address,omitempty" gorm:"foreignKey:AddressID"`
}

// Base CryptoDeposit model
type CryptoDeposit struct {
	gorm.Model
	AddressID       uint      `json:"address_id" gorm:"not null;index"`
	TxHash          string    `json:"tx_hash" gorm:"unique;not null"`
	FromAddress     string    `json:"from_address" gorm:"not null"`
	Amount          string    `json:"amount" gorm:"type:decimal(28,18);not null"`
	Confirmations   uint      `json:"confirmations" gorm:"default:0"`
	RequiredConfirms uint     `json:"required_confirms" gorm:"default:6"`
	Status          string    `json:"status" gorm:"default:pending"` // pending, confirmed, credited
	BlockHeight     uint64    `json:"block_height,omitempty"`
	BlockTime       *time.Time `json:"block_time,omitempty"`
	CreditedAt      *time.Time `json:"credited_at,omitempty"`
	Network         string    `json:"network" gorm:"not null"`
	CryptoType      string    `json:"crypto_type" gorm:"not null"`
	
	// Relationships
	Address CryptoAddress `json:"address,omitempty" gorm:"foreignKey:AddressID"`
}

// Base CryptoWithdrawal model
type CryptoWithdrawal struct {
	gorm.Model
	FromAddressID   uint      `json:"from_address_id" gorm:"not null;index"`
	ToAddress       string    `json:"to_address" gorm:"not null"`
	Amount          string    `json:"amount" gorm:"type:decimal(28,18);not null"`
	Fee             string    `json:"fee" gorm:"type:decimal(28,18);not null"`
	TxHash          string    `json:"tx_hash,omitempty"`
	Status          string    `json:"status" gorm:"default:pending"` // pending, broadcasting, confirmed, failed
	FailureReason   string    `json:"failure_reason,omitempty"`
	BlockHeight     uint64    `json:"block_height,omitempty"`
	Confirmations   uint      `json:"confirmations" gorm:"default:0"`
	BroadcastAt     *time.Time `json:"broadcast_at,omitempty"`
	ConfirmedAt     *time.Time `json:"confirmed_at,omitempty"`
	Network         string    `json:"network" gorm:"not null"`
	CryptoType      string    `json:"crypto_type" gorm:"not null"`
	Memo            string    `json:"memo,omitempty"`
	
	// Relationships
	FromAddress CryptoAddress `json:"from_address,omitempty" gorm:"foreignKey:FromAddressID"`
}

// Base CryptoUTXO model (mainly for Bitcoin-like cryptocurrencies)
type CryptoUTXO struct {
	gorm.Model
	AddressID     uint   `json:"address_id" gorm:"not null;index"`
	TxHash        string `json:"tx_hash" gorm:"not null"`
	Vout          uint   `json:"vout" gorm:"not null"` // Output index
	Amount        string `json:"amount" gorm:"type:decimal(28,18);not null"`
	ScriptPubKey  string `json:"script_pub_key"`
	IsSpent       bool   `json:"is_spent" gorm:"default:false"`
	SpentTxHash   string `json:"spent_tx_hash,omitempty"`
	SpentAt       *time.Time `json:"spent_at,omitempty"`
	BlockHeight   uint64 `json:"block_height,omitempty"`
	Confirmations uint   `json:"confirmations" gorm:"default:0"`
	
	// Relationships
	Address CryptoAddress `json:"address,omitempty" gorm:"foreignKey:AddressID"`
}

// Bitcoin specific transaction
type BitcoinTransaction struct {
	CryptoTransaction
	Size        uint   `json:"size,omitempty"`
	VSize       uint   `json:"vsize,omitempty"`       // Virtual size (for SegWit)
	Weight      uint   `json:"weight,omitempty"`      // Transaction weight
	Version     uint   `json:"version" gorm:"default:1"`
	LockTime    uint   `json:"lock_time" gorm:"default:0"`
	InputCount  uint   `json:"input_count"`
	OutputCount uint   `json:"output_count"`
	FeeRate     string `json:"fee_rate,omitempty" gorm:"type:decimal(10,8)"` // sat/vB
	RBF         bool   `json:"rbf" gorm:"default:false"` // Replace-by-fee
}

// Bitcoin specific deposit
type BitcoinDeposit struct {
	CryptoDeposit
	Vout         uint   `json:"vout"` // Output index in transaction
	ScriptPubKey string `json:"script_pub_key,omitempty"`
	CoinbaseJustification string `json:"coinbase_justification,omitempty"` // If from coinbase tx
}

// Bitcoin specific withdrawal
type BitcoinWithdrawal struct {
	CryptoWithdrawal
	ChangeAddress string `json:"change_address,omitempty"`
	ChangeAmount  string `json:"change_amount,omitempty" gorm:"type:decimal(28,18)"`
	FeeRate       string `json:"fee_rate,omitempty" gorm:"type:decimal(10,8)"` // sat/vB
	RBF           bool   `json:"rbf" gorm:"default:false"`
	InputCount    uint   `json:"input_count"`
	OutputCount   uint   `json:"output_count"`
	Size          uint   `json:"size,omitempty"`
	VSize         uint   `json:"vsize,omitempty"`
}

// Ethereum specific transaction
type EthereumTransaction struct {
	CryptoTransaction
	Nonce           uint64 `json:"nonce"`
	GasPrice        string `json:"gas_price" gorm:"type:decimal(28,0)"`        // in wei
	GasLimit        uint64 `json:"gas_limit"`
	GasUsed         uint64 `json:"gas_used,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	Input           string `json:"input,omitempty" gorm:"type:text"`           // Contract call data
	IsContract      bool   `json:"is_contract" gorm:"default:false"`
	TokenTransfer   bool   `json:"token_transfer" gorm:"default:false"`
	TokenSymbol     string `json:"token_symbol,omitempty"`
	TokenAmount     string `json:"token_amount,omitempty" gorm:"type:decimal(28,18)"`
}

// Ethereum specific deposit
type EthereumDeposit struct {
	CryptoDeposit
	GasPrice        string `json:"gas_price,omitempty" gorm:"type:decimal(28,0)"`
	GasUsed         uint64 `json:"gas_used,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	IsTokenTransfer bool   `json:"is_token_transfer" gorm:"default:false"`
	TokenSymbol     string `json:"token_symbol,omitempty"`
	TokenAmount     string `json:"token_amount,omitempty" gorm:"type:decimal(28,18)"`
	LogIndex        uint   `json:"log_index,omitempty"` // For token transfers
}

// Ethereum specific withdrawal
type EthereumWithdrawal struct {
	CryptoWithdrawal
	Nonce           uint64 `json:"nonce"`
	GasPrice        string `json:"gas_price" gorm:"type:decimal(28,0)"` // in wei
	GasLimit        uint64 `json:"gas_limit"`
	GasUsed         uint64 `json:"gas_used,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	IsTokenTransfer bool   `json:"is_token_transfer" gorm:"default:false"`
	TokenSymbol     string `json:"token_symbol,omitempty"`
	TokenAmount     string `json:"token_amount,omitempty" gorm:"type:decimal(28,18)"`
	Data            string `json:"data,omitempty" gorm:"type:text"` // Contract call data
}