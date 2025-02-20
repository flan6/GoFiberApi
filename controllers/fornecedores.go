package controllers

import (
	"api/models"
	"api/repository"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateFornecedor(c *fiber.Ctx) error {
	var fornecedor models.Fornecedor

	if err := c.BodyParser(&fornecedor); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := fornecedor.Format(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Create(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fornecedor)
}

func ListAllFornecedores(c *fiber.Ctx) error {
	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fornecedores, err := r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fornecedores)

}

func SearchFornecedores(c *fiber.Ctx) error {
	param := c.Query("query")
	if len(param) < 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("insira ao menos 4 letras na busca").Error()})
	}
	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fornecedores, err := r.Search(param)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fornecedores)
}

func FindFornecedor(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}

	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	fornecedor, err := r.Find(uint(id))
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.JSON(fornecedor)

}

func EditFornecedor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var fornecedor models.Fornecedor

	if err := c.BodyParser(&fornecedor); err != nil {

		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	fornecedor.ID = uint(id)

	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Update(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)

}

func DeleteForncedor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var fornecedor models.Fornecedor
	fornecedor.ID = uint(id)

	r, err := repository.NewFornecedorRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.Delete(fornecedor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)

}
