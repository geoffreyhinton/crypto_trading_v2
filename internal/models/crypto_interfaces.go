package models

import (
	"context"
	"math/big"
)

// AddressManager interface for common address operations
type AddressManager interface {
	GenerateAddress(network string) error
	ValidateAddress(address string) bool
	GetBalance(ctx context.Context) (*big.Int, error)
	SyncBalance(ctx context.Context) error
}

// TransactionManager interface for transaction operations
type TransactionManager interface {
	CreateTransaction(toAddress string, amount *big.Int) error
	BroadcastTransaction(ctx context.Context) error
	GetTransactionStatus(ctx context.Context) (string, error)
	ValidateTransaction() error
}

// WalletService interface for wallet operations
type WalletService interface {
	CreateWallet(userID uint, cryptoType string, network string) (AddressManager, error)
	ImportWallet(userID uint, privateKey string, cryptoType string, network string) (AddressManager, error)
	GetWallet(addressID uint) (AddressManager, error)
	ListWallets(userID uint, cryptoType string) ([]AddressManager, error)
}

// DepositService interface for deposit operations
type DepositService interface {
	ProcessDeposit(txHash string, addressID uint) error
	ConfirmDeposit(depositID uint) error
	CreditDeposit(depositID uint) error
	GetPendingDeposits(addressID uint) ([]CryptoDeposit, error)
}

// WithdrawalService interface for withdrawal operations
type WithdrawalService interface {
	CreateWithdrawal(fromAddressID uint, toAddress string, amount *big.Int) (CryptoWithdrawal, error)
	ProcessWithdrawal(withdrawalID uint) error
	BroadcastWithdrawal(withdrawalID uint) error
	ConfirmWithdrawal(withdrawalID uint) error
	CancelWithdrawal(withdrawalID uint, reason string) error
}

// BlockchainMonitor interface for blockchain monitoring
type BlockchainMonitor interface {
	ScanNewBlocks(ctx context.Context) error
	ProcessBlock(blockHeight uint64) error
	ScanTransactions(addresses []string) ([]CryptoTransaction, error)
	GetLatestBlockHeight() (uint64, error)
}

// CryptoValidator interface for validation operations
type CryptoValidator interface {
	ValidateAddress(address string) bool
	ValidatePrivateKey(privateKey string) bool
	ValidateAmount(amount string) bool
	ValidateNetwork(network string) bool
}

// Constants for crypto types
const (
	CryptoTypeBitcoin  = "bitcoin"
	CryptoTypeEthereum = "ethereum"
	CryptoTypeLitecoin = "litecoin"
	CryptoTypeDogecoin = "dogecoin"
)

// Constants for networks
const (
	NetworkMainnet = "mainnet"
	NetworkTestnet = "testnet"
	NetworkRegtest = "regtest"
	NetworkRopsten = "ropsten"
	NetworkGoerli  = "goerli"
	NetworkSepolia = "sepolia"
)

// Constants for transaction status
const (
	StatusPending     = "pending"
	StatusBroadcasted = "broadcasted"
	StatusConfirmed   = "confirmed"
	StatusFailed      = "failed"
	StatusCredited    = "credited"
)

// Constants for transaction direction
const (
	DirectionIncoming = "incoming"
	DirectionOutgoing = "outgoing"
	DirectionInternal = "internal"
)

// Constants for Bitcoin address types
const (
	AddressTypeP2PKH = "P2PKH" // Pay to Public Key Hash (legacy)
	AddressTypeP2SH  = "P2SH"  // Pay to Script Hash
	AddressTypeP2WPKH = "P2WPKH" // Pay to Witness Public Key Hash (native SegWit)
	AddressTypeP2WSH  = "P2WSH"  // Pay to Witness Script Hash
	AddressTypeP2TR   = "P2TR"   // Pay to Taproot (Taproot)
)

// Error types
type CryptoError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e CryptoError) Error() string {
	return e.Message
}

// Common error codes
const (
	ErrInvalidAddress     = "INVALID_ADDRESS"
	ErrInvalidAmount      = "INVALID_AMOUNT"
	ErrInsufficientFunds  = "INSUFFICIENT_FUNDS"
	ErrTransactionFailed  = "TRANSACTION_FAILED"
	ErrNetworkError       = "NETWORK_ERROR"
	ErrInvalidPrivateKey  = "INVALID_PRIVATE_KEY"
	ErrWalletNotFound     = "WALLET_NOT_FOUND"
	ErrDuplicateTransaction = "DUPLICATE_TRANSACTION"
)