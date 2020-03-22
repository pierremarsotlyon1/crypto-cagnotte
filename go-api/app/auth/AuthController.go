package auth

import (
	"context"
	"crypto-cagnotte/go-api/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var SecretJwt = "posqdi974lkdjsqo@@mlds86546789"

func Register(c echo.Context) error {

	// Récupération des données
	registerModel := new(RegisterModel)
	if err := c.Bind(registerModel); err != nil {
		log.Print(err)
		return c.JSON(400, err)
	}

	// Check des infos
	if !registerModel.IsValid() {
		return c.NoContent(400)
	}

	// On regarde si un utilisateur existe déjà avec cet email
	usersCollection := app.MongoDatabase.Collection("users")
	count, err := usersCollection.CountDocuments(context.Background(), bson.M{"email": bson.M{"$eq": registerModel.Email}})
	if err != nil {
		return c.JSON(400, err)
	}

	if count > 0 {
		// User exists with same email
		return c.JSON(400, map[string]string{
			"error": "user_exists",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerModel.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "hash_err",
		})
	}

	user := new(User)
	user.Firstname = registerModel.Firstname
	user.Lastname = registerModel.Lastname
	user.Email = registerModel.Email
	user.Password = string(hash)

	insertResult, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "create_user_err",
		})
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return c.JSON(400, map[string]string{
			"error": "jwt_err",
		})
	}
	// Create token
	t, err := generateJwtToken(oid.Hex())
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "generate_jwt_token",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"token": t,
	})
}

func Login(c echo.Context) error {
	loginModel := new(LoginModel)
	if err := c.Bind(loginModel); err != nil {
		return c.JSON(400, err)
	}

	if !loginModel.IsValid() {
		return c.NoContent(400)
	}

	// Récupération de l'utilisateur
	// On regarde si un utilisateur existe déjà avec cet email
	usersCollection := app.MongoDatabase.Collection("users")
	user := new(User)
	if err := usersCollection.FindOne(context.Background(), bson.M{"email": bson.M{"$eq": loginModel.Email}}).Decode(&user); err != nil {
		return c.JSON(400, map[string]string{
			"error": "get_user_email",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password)); err != nil {
		return c.JSON(400, map[string]string{
			"error": "wrong_password",
		})
	}

	token, err := generateJwtToken(user.ID.Hex())
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "generate_jwt_token",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}

func generateJwtToken(id string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(SecretJwt))
}
