package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"bagas": {
		AuthToken: "CBA321",
		Username:  "bagas",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"bagas": {
		Coins:    100,
		Username: "bagas",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func init() {
	mockLoginDetails["alice"] = LoginDetails{AuthToken: "ABC123", Username: "alice"}
	mockLoginDetails["bob"] = LoginDetails{AuthToken: "DEF456", Username: "bob"}
	mockLoginDetails["charlie"] = LoginDetails{AuthToken: "GHI789", Username: "charlie"}
	mockLoginDetails["dave"] = LoginDetails{AuthToken: "JKL012", Username: "dave"}
	mockLoginDetails["eve"] = LoginDetails{AuthToken: "MNO345", Username: "eve"}
	mockLoginDetails["frank"] = LoginDetails{AuthToken: "PQR678", Username: "frank"}
	mockLoginDetails["grace"] = LoginDetails{AuthToken: "STU901", Username: "grace"}
	mockLoginDetails["heidi"] = LoginDetails{AuthToken: "VWX234", Username: "heidi"}
	mockLoginDetails["ivan"] = LoginDetails{AuthToken: "YZA567", Username: "ivan"}
	mockLoginDetails["judy"] = LoginDetails{AuthToken: "BCD890", Username: "judy"}

	mockCoinDetails["alice"] = CoinDetails{Coins: 200, Username: "alice"}
	mockCoinDetails["bob"] = CoinDetails{Coins: 150, Username: "bob"}
	mockCoinDetails["charlie"] = CoinDetails{Coins: 300, Username: "charlie"}
	mockCoinDetails["dave"] = CoinDetails{Coins: 250, Username: "dave"}
	mockCoinDetails["eve"] = CoinDetails{Coins: 100, Username: "eve"}
	mockCoinDetails["frank"] = CoinDetails{Coins: 350, Username: "frank"}
	mockCoinDetails["grace"] = CoinDetails{Coins: 400, Username: "grace"}
	mockCoinDetails["heidi"] = CoinDetails{Coins: 450, Username: "heidi"}
	mockCoinDetails["ivan"] = CoinDetails{Coins: 500, Username: "ivan"}
	mockCoinDetails["judy"] = CoinDetails{Coins: 550, Username: "judy"}
}
