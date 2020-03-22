package coinbase

import (
	"context"
	"crypto-cagnotte/go-api/app/cagnotte"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"sync"
)

var mux sync.Mutex

func ReceiveNotification(c echo.Context) error {
	mux.Lock()
	defer mux.Unlock()

	notification := new(Notification)
	if err := c.Bind(notification); err != nil {
		return c.NoContent(400)
	}

	// TODO check coinbase notification origin
	if notification.Type == "ping" {
		return c.NoContent(200)
	}

	if notification.Type != "wallet:deposit:completed" {
		// On ignore pour l'instance
		// A voir après
		return c.NoContent(200)
	}

	// On récupère la cagnotte associée
	ca := new(cagnotte.Cagnotte)
	if err := cagnotte.GetCagnottesCollection().FindOne(context.Background(), bson.M{"wallets.address": bson.M{"$eq": notification.Data.Address}}).Decode(ca); err != nil {
		log.Println("Erreur lors de la récupération de la cagnotte pour la notification " + notification.Id)
		return c.NoContent(400)
	}

	amountFloat, err := strconv.ParseFloat(notification.AdditionalData.Amount.Amount, 64)
	if err != nil {
		log.Println("Erreur lors de la récupération du montant du paiement pour la notification " + notification.Id)
		return c.NoContent(400)
	}

	// On récupère le wallet associé à la notification
	for _, wc := range ca.Wallets {
		if wc.Address == notification.Data.Address && wc.Currency == notification.AdditionalData.Amount.Currency {
			wc.AvailableAmount += amountFloat
			wc.Amount += amountFloat
			ca.TotalAmount += amountFloat

			cagnotte.GetCagnottesCollection().UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": ca.ID}}, bson.M{"$set": bson.M{"wallets": ca.Wallets, "totalAmount": ca.TotalAmount}})
			break
		}
	}

	return c.NoContent(200)
}
