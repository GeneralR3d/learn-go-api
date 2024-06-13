package tools

import(
	log "github.com/sirupsen/logrus"
)

// database collections
type LoginDetails struct{
	AuthToken string
	Username string
}

type CoinDetails struct{
	Coins int64
	Username string
}

type DatabaseInterface interface{	//	interface w 3 broad methods
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDataBase() (*DatabaseInterface, error){
	
	var database DatabaseInterface = &mockDB{}	// custom struct

	var err error = database.SetupDatabase()	// create database
	if err != nil{
		log.Error(err)
		return nil, err
	}
	return &database, nil
}