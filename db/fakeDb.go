package db

type fakeDb struct{}

var loginMap = map[string]LoginDetails{
	"alice": LoginDetails{
		Username:  "alice",
		AuthToken: "AliceToken",
	},
	"bob": LoginDetails{
		Username:  "bob",
		AuthToken: "BobToken",
	},
}

var coinMap = map[string]CoinDtl{
	"alice": CoinDtl{
		Username: "alice",
		Coins:    330,
	},
	"bob": CoinDtl{
		Username: "bob",
		Coins:    400,
	},
}

func (d *fakeDb) GetUserLoginDtl(username string) *LoginDetails {
	temp := loginMap[username]
	return &temp
}

func (d *fakeDb) GetUserCoins(username string) *CoinDtl {
	temp := coinMap[username]
	return &temp
}

func (d *fakeDb) SetupDb() error {
	return nil
}
