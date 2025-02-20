package controllers

import (
	"api/models"
	"api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePeca(c *fiber.Ctx) error {
	var peca models.Peca
	if err := c.BodyParser(&peca); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := peca.Format(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := r.Create(peca); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(peca)

}

func ListAllPecas(c *fiber.Ctx) error {
	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	pecas, err := r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.JSON(pecas)

}

func ListPecas(c *fiber.Ctx) error {

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	pecas, err := r.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(pecas)
}

func FindPeca(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	peca, err := r.Find(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.JSON(peca)
}

func EditPeca(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	var peca models.Peca

	err = c.BodyParser(&peca)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := peca.Format(); err != nil && peca.Descricao != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	peca.ID = uint(id)
	if err := r.Update(peca); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.SendStatus(200)
}

func DeletePeca(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

	}

	r, err := repository.NewPecaRepo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	var peca models.Peca
	peca.ID = uint(id)
	if err := r.Delete(peca); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})

	}
	return c.SendStatus(200)

}

func SaerchPeca(ctx *fiber.Ctx) error {
	filter := ctx.Query("query")
	if len(filter) < 4 {
		return ctx.Status(fiber.StatusBadRequest).JSON((fiber.Map{"Error": "Insira ao menos 4 letras na pesquisa"}))

	}
	repo, err := repository.NewPecaRepo()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	results, err := repo.Search(filter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return ctx.JSON(results)
}
