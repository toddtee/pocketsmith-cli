package app

import (
	"fmt"
	"strings"
	"time"
)

// Client for Pocketsmith API
type Client struct {
	BaseURL string
	User    string
	APIKey  string
}

func (c Client) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "BaseURL: %v\n", c.BaseURL)
	fmt.Fprintf(&b, "User: %v\n", c.User)
	fmt.Fprintf(&b, "Api Key: %v\n", c.APIKey)

	return b.String()
}

// Account represents a Pocketsmith account summary
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

// Institution represents a financial institution
type Institution struct {
	CurrencyCode string    `json:"currency_code"`
	Title        string    `json:"title"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	ID           int       `json:"id"`
}

// PrimaryTransactionAccount represents the primary transaction account associated with the Pocketsmith account
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

// PrimaryScenario defines the primary forecasting scenario associated with the primary transaction account
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

// Scenario defines a forecasting scenario
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

// Scenarios is a list of forecasting scenarios created within the Pocketsmith Account
type Scenarios struct {
	Scenarios []Scenario
}

// TransactionAccount defines a banking transaction account
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

// TransactionAccounts defines a list of banking transaction accounts
type TransactionAccounts struct {
	TransactionAccounts []TransactionAccount
}

// User defines a Pocketsmith user
type User struct {
	ID                       int    `json:"id"`
	Login                    string `json:"login"`
	Name                     string `json:"name"`
	Email                    string `json:"email"`
	AvatarURL                string `json:"avatar_url"`
	BetaUser                 bool   `json:"beta_user"`
	TimeZone                 string `json:"time_zone"`
	WeekStartDay             int    `json:"week_start_day"`
	IsReviewingTransactions  bool   `json:"is_reviewing_transactions"`
	BaseCurrencyCode         string `json:"base_currency_code"`
	AlwaysShowBaseCurrency   bool   `json:"always_show_base_currency"`
	UsingMultipleCurrencies  bool   `json:"using_multiple_currencies"`
	AvailableAccounts        int    `json:"available_accounts"`
	AvailableBudgets         int    `json:"available_budgets"`
	ForecastLastUpdatedAt    string `json:"forecast_last_updated_at"`
	ForecastLastAccessedAt   string `json:"forecast_last_accessed_at"`
	ForecastStartDate        string `json:"forecast_start_date"`
	ForecastEndDate          string `json:"forecast_end_date"`
	ForecastDeferRecalculate bool   `json:"forecast_defer_recalculate"`
	ForecastNeedsRecalculate bool   `json:"forecast_needs_recalculate"`
	LastLoggedInAt           string `json:"last_logged_in_at"`
	LastActivityAt           string `json:"last_activity_at"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
}

func (u User) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Name: %v\n", u.Name)
	fmt.Fprintf(&b, "Email: %v", u.Email)

	return b.String()
}

// Base URL for the Pocketsmith API
const (
	BaseURL = "https://api.pocketsmith.com/v2"
)

// Resource type for Pocketsmith HTTP responses
const (
	ResourceType = "application/json; charset=utf-8"
)
