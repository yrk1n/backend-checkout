package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yrk1n/backend-checkout/models"
	"github.com/yrk1n/backend-checkout/services"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(s *services.CartService) *CartHandler {
	return &CartHandler{
		service: s,
	}
}

func (ch *CartHandler) DisplayItems(c *fiber.Ctx) error {
	items, err := ch.service.GetItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(items)
}

func (ch *CartHandler) AddItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	err := ch.service.AddItemToCart(1, &item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item added successfully",
	})
}

func (ch *CartHandler) AddVasItemToItem(c *fiber.Ctx) error {
	var vasItem models.VasItem
	if err := c.BodyParser(&vasItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	itemId, err := c.ParamsInt("itemId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	if vasItem.ParentItemId != itemId {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Mismatch between vasItem's ParentItemId and itemId in the URL",
		})
	}

	_, err = ch.service.GetItemByID(vasItem.ParentItemId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Item with ID %d not found", vasItem.ParentItemId),
		})
	}
	err = ch.service.GetVasItemRepo().CreateVasItemForItem(itemId, &vasItem)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create and associate VasItem with Item: %s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"message": "VasItem added successfully",
	})
}

func (ch *CartHandler) ResetCart(c *fiber.Ctx) error {
	err := ch.service.ResetCart(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cart reset successfully",
	})
}
