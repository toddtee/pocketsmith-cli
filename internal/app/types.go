package app

import (
	"net/http"
	"time"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
type Client struct {
	HTTPClient Doer
	Config     *Config
}

type Config struct {
	BaseURL string
	User    string
	APIKey  string
}

type Account struct {
	ID                           int                       `json:"id"`
	Title                        string                    `json:"title"`
	CurrencyCode                 string                    `json:"currency_code"`
	Type                         string                    `json:"type"`
	IsNetWorth                   bool                      `json:"is_net_worth"`
	PrimaryTransactionAccount    PrimaryTransactionAccount `json:"primary_transaction_account"`
	PrimaryScenario              PrimaryScenario           `json:"primary_scenario"`
	TransactionAccounts          TransactionAccounts       `json:"transaction_accounts"`
	Scenarios                    Scenarios                 `json:"scenarios"`
	CreatedAt                    string                    `json:"created_at"`
	UpdatedAt                    string                    `json:"updated_at"`
	CurrentBalance               float64                   `json:"current_balance"`
	CurrentBalanceDate           string                    `json:"current_balance_date"`
	CurrentBalanceInBaseCurrency float64                   `json:"current_balance_in_base_currency"`
	CurrentBalanceExchangeRate   float64                   `json:"current_balance_exchange_rate"`
	SafeBalance                  float64                   `json:"safe_balance"`
	SafeBalanceInBaseCurrency    float64                   `json:"safe_balance_in_base_currency"`
}
type Institution struct {
	CurrencyCode string    `json:"currency_code"`
	Title        string    `json:"title"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	ID           int       `json:"id"`
}

type PrimaryTransactionAccount struct {
	ID                           int         `json:"id"`
	Name                         string      `json:"name"`
	Number                       string      `json:"number"`
	CurrentBalance               float64     `json:"current_balance"`
	CurrentBalanceDate           string      `json:"current_balance_date"`
	CurrentBalanceInBaseCurrency float64     `json:"current_balance_in_base_currency"`
	CurrentBalanceExchangeRate   float64     `json:"current_balance_exchange_rate"`
	SafeBalance                  float64     `json:"safe_balance"`
	SafeBalanceInBaseCurrency    float64     `json:"safe_balance_in_base_currency"`
	StartingBalance              float64     `json:"starting_balance"`
	StartingBalanceDate          string      `json:"starting_balance_date"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Institution                  Institution `json:"institution"`
}

type PrimaryScenario struct {
	ID                           int       `json:"id"`
	Title                        string    `json:"title"`
	Description                  string    `json:"description"`
	InterestRate                 float64   `json:"interest_rate"`
	InterestRateRepeatID         int       `json:"interest_rate_repeat_id"`
	Type                         string    `json:"type"`
	MinimumValue                 int       `json:"minimum-value"`
	MaximumValue                 int       `json:"maximum-value"`
	AchieveDate                  string    `json:"achieve_date"`
	StartingBalance              int       `json:"starting_balance"`
	StartingBalanceDate          string    `json:"starting_balance_date"`
	ClosingBalance               float64   `json:"closing_balance"`
	ClosingBalanceDate           string    `json:"closing_balance_date"`
	CurrentBalance               float64   `json:"current_balance"`
	CurrentBalanceDate           string    `json:"current_balance_date"`
	CurrentBalanceInBaseCurrency float64   `json:"current_balance_in_base_currency"`
	CurrentBalanceExchangeRate   float64   `json:"current_balance_exchange_rate"`
	SafeBalance                  float64   `json:"safe_balance"`
	SafeBalanceInBaseCurrency    float64   `json:"safe_balance_in_base_currency"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}
type Scenario struct {
	ID                           int       `json:"id"`
	Title                        string    `json:"title"`
	Description                  string    `json:"description"`
	InterestRate                 float64   `json:"interest_rate"`
	InterestRateRepeatID         int       `json:"interest_rate_repeat_id"`
	Type                         string    `json:"type"`
	MinimumValue                 int       `json:"minimum-value"`
	MaximumValue                 int       `json:"maximum-value"`
	AchieveDate                  string    `json:"achieve_date"`
	StartingBalance              int       `json:"starting_balance"`
	StartingBalanceDate          string    `json:"starting_balance_date"`
	ClosingBalance               float64   `json:"closing_balance"`
	ClosingBalanceDate           string    `json:"closing_balance_date"`
	CurrentBalance               float64   `json:"current_balance"`
	CurrentBalanceDate           string    `json:"current_balance_date"`
	CurrentBalanceInBaseCurrency float64   `json:"current_balance_in_base_currency"`
	CurrentBalanceExchangeRate   float64   `json:"current_balance_exchange_rate"`
	SafeBalance                  float64   `json:"safe_balance"`
	SafeBalanceInBaseCurrency    float64   `json:"safe_balance_in_base_currency"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}
type Scenarios struct {
	Scenarios []Scenario
}

type TransactionAccount struct {
	ID                           int         `json:"id"`
	Name                         string      `json:"name"`
	Number                       string      `json:"number"`
	CurrentBalance               float64     `json:"current_balance"`
	CurrentBalanceDate           string      `json:"current_balance_date"`
	CurrentBalanceInBaseCurrency float64     `json:"current_balance_in_base_currency"`
	CurrentBalanceExchangeRate   float64     `json:"current_balance_exchange_rate"`
	SafeBalance                  float64     `json:"safe_balance"`
	SafeBalanceInBaseCurrency    float64     `json:"safe_balance_in_base_currency"`
	StartingBalance              float64     `json:"starting_balance"`
	StartingBalanceDate          string      `json:"starting_balance_date"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Institution                  Institution `json:"institution"`
	CurrencyCode                 string      `json:"currency_code"`
	Type                         string      `json:"type"`
}
type TransactionAccounts struct {
	TransactionAccounts []TransactionAccount
}

type User struct {
	Id                         int    `json:"id"`
	Login                      string `json:"login"`
	Name                       string `json:"name"`
	Email                      string `json:"email"`
	Avatar_url                 string `json:"avatar_url"`
	Beta_user                  bool   `json:"beta_user"`
	Time_zone                  string `json:"time_zone"`
	Week_start_day             int    `json:"week_start_day"`
	Is_reviewing_transactions  bool   `json:"is_reviewing_transactions"`
	Base_currency_code         string `json:"base_currency_code"`
	Always_show_base_currency  bool   `json:"always_show_base_currency"`
	Using_multiple_currencies  bool   `json:"using_multiple_currencies"`
	Available_accounts         int    `json:"available_accounts"`
	Available_budgets          int    `json:"available_budgets"`
	Forecast_last_updated_at   string `json:"forecast_last_updated_at"`
	Forecast_last_accessed_at  string `json:"forecast_last_accessed_at"`
	Forecast_start_date        string `json:"forecast_start_date"`
	Forecast_end_date          string `json:"forecast_end_date"`
	Forecast_defer_recalculate bool   `json:"forecast_defer_recalculate"`
	Forecast_needs_recalculate bool   `json:"forecast_needs_recalculate"`
	Last_logged_in_at          string `json:"last_logged_in_at"`
	Last_activity_at           string `json:"last_activity_at"`
	Created_at                 string `json:"created_at"`
	Updated_at                 string `json:"updated_at"`
}

const (
	BaseURL = "https://api.pocketsmith.com/v2"
)

const (
	ResourceType = "application/json; charset=utf-8"
)
