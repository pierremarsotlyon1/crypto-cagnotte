package cagnotte

import (
	"context"
	"crypto-cagnotte/go-api/app"
	"crypto-cagnotte/go-api/app/coinbase"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var apiKey = "m0ytTDsih1T/16d30RlRK4KGBxUUwJvLRFOxWHgPr/qVmahTcMyrKVEN"
var secretKey = "kp8+pU3vM7sYzmvJS2lEkJM9Yk7ZdFigkmqVddeu0brDcX/7YP6DE+gwMLq7qSp2VfHXoMWhC062RTl7LVQLRQ=="

func Add(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	userId := user.Claims.(jwt.MapClaims)["id"].(string)

	// Récupération de l'id de l'user connecté
	fmt.Println(userId)
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]string{
			"error": "objectid_user",
		})
	}

	cagnotte := new(Cagnotte)
	if err := c.Bind(cagnotte); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_cagnotte",
		})
	}

	if !cagnotte.IsValid() {
		return c.JSON(400, map[string]string{
			"error": "cagnotte_check_infos",
		})
	}

	// Création des addresses sur l'exchange
	usdcAddress := coinbase.GetUSDCAddress()
	daiAddress := coinbase.GetDAIAddress()

	if usdcAddress == nil || daiAddress == nil {
		return c.JSON(400, map[string]string{
			"error": "generate_wallet_address",
		})
	}

	// Création des wallets
	usdcWallet := new(Wallet)
	usdcWallet.Address = usdcAddress.Data.Address
	usdcWallet.Currency = "USDC"

	daiWallet := new(Wallet)
	daiWallet.Address = daiAddress.Data.Address
	daiWallet.Currency = "DAI"

	// Ajout des wallets
	cagnotte.Wallets = append(cagnotte.Wallets, *usdcWallet)
	cagnotte.Wallets = append(cagnotte.Wallets, *daiWallet)

	cagnotte.Creator = oid

	// Enregistrement
	cagnottesCollection := app.MongoDatabase.Collection("cagnottes")
	insertResult, err := cagnottesCollection.InsertOne(context.Background(), cagnotte)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "add_cagnotte",
		})
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return c.JSON(400, map[string]string{
			"error": "oid",
		})
	}

	cagnotte.ID = oid
	return c.JSON(200, cagnotte)
}
