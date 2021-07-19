package handlers

import "basic-rest/pkg/storage/sqlite"

type Service struct {
	ProductService *sqlite.ProductService
}
