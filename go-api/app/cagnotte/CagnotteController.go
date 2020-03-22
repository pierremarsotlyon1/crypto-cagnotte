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
	"log"
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
	cagnotte.Status = 1
	cagnotte.Creator = oid

	// Création des addresses sur l'exchange
	if addCagnotte.UseUSDCWallet {
		usdcAddress := coinbase.GetUSDCAddress()
		if usdcAddress == nil {
			return c.JSON(400, map[string]string{
				"error": "generate_wallet_address",
			})
		}

		cagnotte.Wallets = append(cagnotte.Wallets, createWallet("USDC", usdcAddress))
	}

	if addCagnotte.UseDAIWallet {
		daiAddress := coinbase.GetDAIAddress()
		if daiAddress == nil {
			return c.JSON(400, map[string]string{
				"error": "generate_wallet_address",
			})
		}

		cagnotte.Wallets = append(cagnotte.Wallets, createWallet("DAI", daiAddress))
	}

	// Enregistrement
	insertResult, err := GetCagnottesCollection().InsertOne(context.Background(), cagnotte)
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

// createWallet - permet de créer un wallet à partir d'une réponse coinbase wallet
func createWallet(currency string, coinbaseAddress *coinbase.CoinbaseAddress) Wallet {
	wallet := new(Wallet)
	wallet.Address = coinbaseAddress.Data.Address
	wallet.Currency = currency
	wallet.Amount = 0

	return *wallet
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
	if err := GetCagnottesCollection().FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": oid}}).Decode(cagnotte); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_cagnotte",
		})
	}

	return c.JSON(200, cagnotte)
}

// Close - Permet de fermer une cagnotte
func Close(c echo.Context) error {
	idCagnotte := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(idCagnotte)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_oid",
		})
	}

	updateResult, err := GetCagnottesCollection().UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": oid}}, bson.M{"$set": bson.M{"status": 2}})

	if err != nil || updateResult.ModifiedCount != 1 {
		return c.JSON(400, map[string]string{
			"error": "close_cagnotte",
		})
	}

	return c.NoContent(200)
}

func Withdraw(c echo.Context) error {
	askWithdraw := new(AskWithdraw)
	if err := c.Bind(askWithdraw); err != nil {
		return c.JSON(400, map[string]string{
			"error": "askwithdraw_check",
		})
	}

	oidCagnotte, err := primitive.ObjectIDFromHex(askWithdraw.Id)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_oid",
		})
	}
	// On récupère la cagnotte
	cagnotte := new(Cagnotte)
	if err := GetCagnottesCollection().FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": oidCagnotte}}).Decode(cagnotte); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_cagnotte",
		})
	}

	// On parcourt les wallets de la cagnotte et on demande un versement
	save := false
	for _, wc := range cagnotte.Wallets {
		for _, withdrawWallet := range askWithdraw.WithdrawWallets {
			if wc.Currency != withdrawWallet.Currency {
				continue
			}

			withdraw := coinbase.Withdraw(wc.Currency, wc.Address, wc.AvailableAmount)
			if withdraw == nil {
				// Problème !!!
				log.Println("Erreur lors de la génération du withdraw pour la cagnotte " + oidCagnotte.Hex())
				continue
			}
			cagnotte.Withdraws = append(cagnotte.Withdraws, *withdraw)
			wc.AvailableAmount = 0
			save = true
			break
		}
	}

	if save {
		updateResults, err := GetCagnottesCollection().UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": oidCagnotte}}, bson.M{"$set": bson.M{"withdraws": cagnotte.Withdraws, "wallets": cagnotte.Wallets}})
		if err != nil || updateResults.ModifiedCount != 1 {
			// Gros problème !!!
			log.Println("Erreur lors de la sauvegarde des withdraws pour la cagnotte " + oidCagnotte.Hex())
		}
	}

	return c.JSON(200, cagnotte)
}

func GetCagnottesCollection() *mongo.Collection {
	return app.MongoDatabase.Collection("cagnottes")
}
