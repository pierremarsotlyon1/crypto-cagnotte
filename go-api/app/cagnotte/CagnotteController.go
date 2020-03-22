package cagnotte

import (
	"context"
	"crypto-cagnotte/go-api/app"
	"crypto-cagnotte/go-api/app/coinbase"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var apiKey = "m0ytTDsih1T/16d30RlRK4KGBxUUwJvLRFOxWHgPr/qVmahTcMyrKVEN"
var secretKey = "kp8+pU3vM7sYzmvJS2lEkJM9Yk7ZdFigkmqVddeu0brDcX/7YP6DE+gwMLq7qSp2VfHXoMWhC062RTl7LVQLRQ=="

// Add - permet d'ajouter une cagnotte
func Add(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	userId := user.Claims.(jwt.MapClaims)["id"].(string)

	// Récupération de l'id de l'user connecté
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "objectid_user",
		})
	}

	addCagnotte := new(AddCagnotte)
	if err := c.Bind(addCagnotte); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_cagnotte",
		})
	}

	if !addCagnotte.IsValid() {
		return c.JSON(400, map[string]string{
			"error": "cagnotte_check_infos",
		})
	}

	cagnotte := new(Cagnotte)
	cagnotte.Days = addCagnotte.Days
	cagnotte.Description = addCagnotte.Description
	cagnotte.Name = addCagnotte.Name
	cagnotte.Creator = oid

	// Création des addresses sur l'exchange
	if addCagnotte.UseUSDCWallet {
		usdcAddress := coinbase.GetUSDCAddress()
		if usdcAddress == nil {
			return c.JSON(400, map[string]string{
				"error": "generate_wallet_address",
			})
		}

		usdcWallet := new(Wallet)
		usdcWallet.Address = usdcAddress.Data.Address
		usdcWallet.Currency = "USDC"
		cagnotte.Wallets = append(cagnotte.Wallets, *usdcWallet)
	}

	if addCagnotte.UseDAIWallet {
		daiAddress := coinbase.GetDAIAddress()
		if daiAddress == nil {
			return c.JSON(400, map[string]string{
				"error": "generate_wallet_address",
			})
		}

		daiWallet := new(Wallet)
		daiWallet.Address = daiAddress.Data.Address
		daiWallet.Currency = "DAI"
		cagnotte.Wallets = append(cagnotte.Wallets, *daiWallet)
	}

	// Enregistrement
	insertResult, err := getCagnottesCollection().InsertOne(context.Background(), cagnotte)
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

// Get - permet de récupérer une cagnotte
func Get(c echo.Context) error {
	idCagnotte := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(idCagnotte)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_oid",
		})
	}

	cagnotte := new(Cagnotte)
	if err := getCagnottesCollection().FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": oid}}).Decode(cagnotte); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_cagnotte",
		})
	}

	return c.JSON(200, cagnotte)
}

func getCagnottesCollection() *mongo.Collection {
	return app.MongoDatabase.Collection("cagnottes")
}
