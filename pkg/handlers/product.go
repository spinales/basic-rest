package handlers

import (
	"basic-rest/pkg/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetProducts(c echo.Context) error {
	prods, err := s.ProductService.Products()
	if err != nil {
		log.Fatalln("buscando equipos: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, prods)
}

func (s *Service) GetProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("turning ID: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	p, err := s.ProductService.Product(uint(id))
	if err != nil {
		log.Fatalln("searching product: %w", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, p)
}

func (s *Service) PostProduct(c echo.Context) error {
	p := model.Product{}
	if err := c.Bind(&p); err != nil {
		log.Fatalln("procesing product: %w", err)
		return c.NoContent(http.StatusBadRequest)
	}

	cp, err := s.ProductService.CreateProduct(&p)
	if err != nil {
		log.Fatalln("creating product: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, cp)
}

func (s *Service) PutProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("turning ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	var p model.Product
	if err := c.Bind(p); err != nil {
		log.Fatalln("procesing product: %w", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	uprod, err := s.ProductService.UpdateProduct(uint(id), &p)
	if err != nil {
		log.Fatalln("updating product: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, uprod)
}

func (s *Service) DeleteProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalln("turning ID: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	err = s.ProductService.DeleteEquipo(uint(id))
	if err != nil {
		log.Fatalln("removing product: %w", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
