package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
	OrderTypeStop   OrderType = "stop"
)

type AccountType string

const (
	AccountTypeSpot    AccountType = "spot"
	AccountTypeMargin  AccountType = "margin"
	AccountTypeFutures AccountType = "futures"
	AccountTypeSavings AccountType = "savings"
	AccountTypeStaking AccountType = "staking"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique;not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
	KYCStatus string `json:"kyc_status" gorm:"default:'pending'"`
}

type Account struct {
	gorm.Model
	UserID           string          `json:"user_id" gorm:"not null"`
	Type             AccountType     `json:"type" gorm:"not null"`
	SubType          string          `json:"sub_type"`
	Currency         string          `json:"currency" gorm:"not null"`
	Balance          decimal.Decimal `json:"balance" gorm:"type:decimal(20,8);default:0"`
	AvailableBalance decimal.Decimal `json:"available_balance" gorm:"type:decimal(20,8);default:0"`
	LockedBalance    decimal.Decimal `json:"locked_balance" gorm:"type:decimal(20,8);default:0"`

	MarginLevel       *decimal.Decimal `json:"margin_level" gorm:"type:decimal(10,4)"`
	MaintenanceMargin *decimal.Decimal `json:"maintenance_margin,omitempty" gorm:"type:decimal(20,8)"`
	UnrealizedPnL     *decimal.Decimal `json:"unrealized_pnl,omitempty" gorm:"type:decimal(20,8)"`
	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
	// Orders    []Order    `json:"orders,omitempty" gorm:"foreignKey:AccountID"`
	// Positions []Position `json:"positions,omitempty" gorm:"foreignKey:AccountID"`
}

// AutoMigrate performs automatic migration for all models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// Core models
		&User{},
		&Account{},
		
		// Crypto wallet models
		&CryptoAddress{},
		&CryptoTransaction{},
		&CryptoDeposit{},
		&CryptoWithdrawal{},
		&CryptoUTXO{},
		
		// Extended crypto models  
		&BitcoinAddress{},
		&EthereumAddress{},
		&BitcoinTransaction{},
		&EthereumTransaction{},
		&BitcoinDeposit{},
		&EthereumDeposit{},
		&BitcoinWithdrawal{},
		&EthereumWithdrawal{},
	)
}
