package backend

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "b.balatamoghna@gmail.com" || pass != "Krypto" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "b.balatamoghna@gmail.com"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + name)
}

func AlertCreate(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	target, _ := strconv.ParseFloat(c.Params("target"), 64)
	alert := CreateAlert(name, c.Params("curr"), target)

	DetailsJSON, _ := json.MarshalIndent(alert, "", "\t")
	return c.SendString("Alert created for user " + name + "\n" + strings.ReplaceAll(string(DetailsJSON), "\\u0026", "&"))
}

func AlertDelete(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	id, _ := strconv.Atoi(c.Params("id"))
	DeleteAlert(id)
	return c.SendString("Alert deleted for user " + name)
}

func FetchAlerts(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)

	alert := GetAllUserAlerts(name)
	DetailsJSON, _ := json.MarshalIndent(alert, "", "\t")
	return c.SendString(strings.ReplaceAll(string(DetailsJSON), "\\u0026", "&"))

}

func FetchTriggeredAlerts(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)

	alert := GetTriggeredUserAlerts(name)
	DetailsJSON, _ := json.MarshalIndent(alert, "", "\t")
	return c.SendString(strings.ReplaceAll(string(DetailsJSON), "\\u0026", "&"))

}
