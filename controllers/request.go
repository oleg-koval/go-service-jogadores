package controllers

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/oleg-koval/go-service-jogadores/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateRequest - POST /api/players
func CreateRequest(ctx *fiber.Ctx) error {
	params := new(struct {
		PlayerID    string
		Interest    string
		Coordinates []float64
	})

	ctx.BodyParser(&params)

	if len(params.PlayerID) == 0 || len(params.Interest) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "PlayerID or Interest not specified.",
		})
		return errors.New("PlayerID or Interest not specified")
	}

	_, errorPlayerNotFound := models.GetPlayerByID(params.PlayerID)
	if errorPlayerNotFound != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": errorPlayerNotFound.Error(),
		})
		return errorPlayerNotFound
	}

	interest := models.CreateRequest(params.PlayerID, params.Interest, params.Coordinates[0], params.Coordinates[1])
	err := mgm.Coll(interest).Create(interest)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":       true,
		"interest": interest,
	})
	return nil
}

// GetAllPoints - GET /api/requests
func GetAllPoints(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Request{})
	requests := []models.Request{}

	err := collection.SimpleFind(&requests, bson.D{})
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(fiber.Map{
		"ok":      true,
		"players": requests,
	})
	return nil
}

type ClosestResponse struct {
	Distance   float64
	PlayerName string
}

// GetAllPoints - GET /api/requests?interest=football&lat=1.1111&lon=2.2222
func GetAllCloserRequests(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.Request{})
	requests := []models.Request{}

	err := collection.SimpleFind(&requests, bson.D{})

	interest := ctx.Query("interest")
	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)

	if err != nil {
		err := errors.New("Non valid latitude")
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
		return err
	}

	lon, err := strconv.ParseFloat(ctx.Query("lon"), 64)

	if err != nil {
		err := errors.New("Non valid longitude")
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
		return err
	}

	var filtered []*models.Request

	for i := 0; i < len(requests); i++ {
		if requests[i].Interest == interest {
			filtered = append(filtered, &requests[i])
		}
	}

	if len(filtered) == 0 {
		err := errors.New("No one with same interest")
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
		return err
	}

	closestPoint := filtered[0]
	var closestDistance float64 = 999999999999

	for i := 0; i < len(filtered); i++ {
		// That's quite a line
		a := math.Pow(math.Sin((filtered[i].Coordinates[0]-lat)/2), 2) + math.Cos(filtered[i].Coordinates[0])*math.Cos(lat)*math.Pow(math.Sin((filtered[i].Coordinates[1]-lon)/2), 2)
		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		// Earth radius on the equator
		distance := 6378000 * c
		if distance < closestDistance {
			closestDistance = distance
			closestPoint = filtered[i]
		}
	}

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	closestPlayer, err := models.GetPlayerByID(closestPoint.PlayerID)

	if err != nil {
		err := errors.New("Something wrong is not right")
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err,
		})
		return err
	}

	response := ClosestResponse{
		Distance:   (closestDistance / 100),
		PlayerName: closestPlayer.Name,
	}

	fmt.Println(response)

	ctx.JSON(fiber.Map{
		"ok":      true,
		"request": response,
	})
	return nil
}
