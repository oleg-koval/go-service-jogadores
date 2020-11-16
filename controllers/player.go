package controllers

import (
	"errors"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/oleg-koval/go-service-jogadores/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllPlayers - GET /api/players
func GetAllPlayers(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Player{})
	players := []models.Player{}

	err := collection.SimpleFind(&players, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":      true,
		"players": players,
	})
	return nil
}

// GetPlayerByID - GET /api/players/:id
func GetPlayerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	player := &models.Player{}
	collection := mgm.Coll(player)

	err := collection.FindByID(id, player)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Player not found.",
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":     true,
		"player": player,
	})
	return nil
}

// CreatePlayer - POST /api/players
func CreatePlayer(ctx *fiber.Ctx) error {
	params := new(struct {
		Name  string
		Email string
	})

	ctx.BodyParser(&params)

	if len(params.Name) == 0 || len(params.Email) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Name or email not specified.",
		})
		return errors.New("Name or email not specified")
	}

	player := models.CreatePlayer(params.Name, params.Email)
	err := mgm.Coll(player).Create(player)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":     true,
		"player": player,
	})

	return nil
}

// DeletePlayer - DELETE /api/players/:id
func DeletePlayer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	player := &models.Player{}
	collection := mgm.Coll(player)

	err := collection.FindByID(id, player)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Player not found.",
		})
		return err
	}

	err = collection.Delete(player)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":     true,
		"player": player,
	})

	return nil
}
