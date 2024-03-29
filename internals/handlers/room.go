package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) {

	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func Room(c *fiber.Ctx) error {

	uuid := c.Params("uuid")

	if uuid == "" {
		c.Status(400)
		return nil
	}
	uuid, suuid, _ := createOrGetRoom(uuid)
}

func RoomWebSocket(c *websocket.Com) {
	uuid := c.Params("uuid")

	if uuid == "" {
		return
	}
	_, _, room := createOrGetRoom(uuid)

}

func createOrGetRoom(uuid string) (string, string, Room) {

}
