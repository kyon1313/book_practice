package templates

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyon1313/books/helper"
	"github.com/kyon1313/books/model"
)

func RenderFlipBookPage(c *fiber.Ctx) error {
	authorBook, err := helper.CallGetAuthorBookEndpoint()
	if err != nil {
		return c.Render("flipbook", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	return c.Render("flipbook", &fiber.Map{
		"Data": authorBook,
	})
}

func RenderIndexTablePage(c *fiber.Ctx) error {
	authorBook, err := helper.CallGetAuthorBookEndpoint()
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	return c.Render("index", &fiber.Map{
		"Data": authorBook,
	})
}

func TemplateAddAuthorBook(c *fiber.Ctx) error {
	form := &model.RequestModelAuthorBook{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	marshallBookRequest, err := helper.UnmarshalFormRequest(form)
	if err != nil {
		return c.JSON(&fiber.Map{
			"Message": err.Error(),
		})

	}
	resp, bookResponseMap, err := helper.CallAddAuthorBookEndpoint(c, marshallBookRequest)
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	authorBook, err := helper.CallGetAuthorBookEndpoint()
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if resp.StatusCode == 400 {
		return c.Render("index", fiber.Map{
			"Message": bookResponseMap["message"],
			"Success": false,
			"Data":    authorBook,
		})
	}

	return c.Render("index", fiber.Map{
		"Message": bookResponseMap["message"],
		"Success": true,
		"Data":    authorBook,
	})
}

func TemplateUpdateAuthorBook(c *fiber.Ctx) error {
	isbnParam := c.Params("isbn")
	form := &model.RequestModelAuthorBook{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	marshallBookRequest, err := helper.UnmarshalFormRequest(form)
	if err != nil {
		return c.JSON(&fiber.Map{
			"Message": err.Error(),
		})

	}

	fmt.Println(form)
	resp, bookResponseMap, err := helper.CallUpdateAuthorBookEndpoint(c, marshallBookRequest, isbnParam)
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	authorBook, err := helper.CallGetAuthorBookEndpoint()
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	if resp.StatusCode == 400 {
		return c.Render("index", fiber.Map{
			"Message": bookResponseMap["message"],
			"Success": false,
			"Data":    authorBook,
		})
	}

	return c.Render("index", fiber.Map{
		"Message": bookResponseMap["message"],
		"Success": true,
		"Data":    authorBook,
	})
}

func TemplateDeleteAuthorBook(c *fiber.Ctx) error {
	isbn := c.Params("isbn")
	resp, err := helper.CallDeleteAuthorBook(isbn)
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}
	if resp.StatusCode == 204 {
		return c.Render("index", &fiber.Map{
			"Message": "No Content",
			"Success": false,
		})
	}

	authorBook, err := helper.CallGetAuthorBookEndpoint()
	if err != nil {
		return c.Render("index", &fiber.Map{
			"Message": err.Error(),
			"Success": false,
		})
	}

	return c.Render("index", &fiber.Map{
		"Message": "item deleted successfully",
		"Success": true,
		"Data":    authorBook,
	})
}
