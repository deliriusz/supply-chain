package api_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/domain/model"
)

var GET_PRODUCTS_GENERATED_COUNT int

func TestCreateProduct(t *testing.T) {
	g := Goblin(t)
	PRODUCT_BASE_URI := "/product"
	productJson := `
	{
		"title": "title",
		"description": "lorem ipslum",
		"price": 3600,
		"quantity": 150,
		"specification": [
			{
				"name": "size",
				"value": "6x2 inch"
			},
			{
				"name": "color",
				"value": "red"
			}
		]
	}`

	g.Describe("Test Product", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		g.It("Should fail when trying to create product with invalid data", func() {
			respRecorder := ServeTestRequest(router, "POST", PRODUCT_BASE_URI, []byte("{}"))
			savedProducts, count := productRepository.GetProducts(10, 0)

			g.Assert(respRecorder.Code).Equal(http.StatusBadRequest)
			g.Assert(count == 0).IsTrue()
			g.Assert(len(savedProducts)).Equal(0)
		})

		g.It("Should create product from valid request", func() {
			respRecorder := ServeTestRequest(router, "POST", PRODUCT_BASE_URI, []byte(productJson))
			savedProducts, count := productRepository.GetProducts(10, 0)

			g.Assert(respRecorder.Code).Equal(http.StatusOK)
			g.Assert(count == 1).IsTrue()
			g.Assert(len(savedProducts)).Equal(1)
		})
	})
}

func TestGetProduct(t *testing.T) {
	g := Goblin(t)
	PRODUCT_BASE_URI := "/product"

	g.Describe("Test GetProduct", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.JustBeforeEach(fillProductsRepo)
		g.After(Cleanup)

		g.It("Should fallback to default offset for invalid offset value", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=abc&limit=12", 0, 12)
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=-45&limit=12", 0, 12)
		})

		g.It("Should fallback to default limit for invalid limit value", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=5&limit=asdf", 5, 10)
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=5&limit=-20", 5, 10)
		})

		g.It("Should return default size list when querying for multiple products without pagination", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI, 0, 10)
		})

		g.It("Should return data for non-default limit and offset properly", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=5&limit=15", 5, 15)
		})

		g.It("Should return only 5 products for offset 95", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=95&limit=12", 95, 12)
		})

		g.It("Should return no data for offset bigger than total products count", func() {
			assertPaginatedProductsAreValid(g, PRODUCT_BASE_URI+"?offset=120&limit=12", 120, 12)
		})
	})
}

func TestImage(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test Image", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

	})
}

func fillProductsRepo() {
	count := GET_PRODUCTS_GENERATED_COUNT

	for i := 0; i < count; i++ {
		product := &model.Product{
			Title:       "title",
			Description: "lorem ipslum",
			Price:       uint(i + 1),
			Quantity:    uint(i + 1),
			Specification: []model.Specification{
				{
					Name:  "size",
					Value: "6x2 inch",
				},
				{
					Name:  "color",
					Value: "red",
				},
			},
		}

		productRepository.CreateProduct(product)
	}
}

func getProductsResponseFromByteArray(data *bytes.Buffer) *api.GetProductsResponse {
	dataBytes, _ := ioutil.ReadAll(data)
	respData := api.GetProductsResponse{}
	json.Unmarshal(dataBytes, &respData)

	return &respData
}

func assertPaginatedProductsAreValid(g *G, uri string, offset, limit int) {
	respRecorder := ServeTestRequest(router, "GET", uri, nil)
	resp := getProductsResponseFromByteArray(respRecorder.Body)
	savedProducts, count := productRepository.GetProducts(uint(limit), uint(offset))

	var expectedReturnedProductsCount int
	if GET_PRODUCTS_GENERATED_COUNT-offset < limit {
		if offset > GET_PRODUCTS_GENERATED_COUNT {
			expectedReturnedProductsCount = 0
		} else {
			expectedReturnedProductsCount = GET_PRODUCTS_GENERATED_COUNT - offset
		}
	} else {
		expectedReturnedProductsCount = limit
	}

	g.Assert(respRecorder.Code).Equal(http.StatusOK)
	g.Assert(resp.Total).Equal(GET_PRODUCTS_GENERATED_COUNT)
	g.Assert(resp.Total == int(count)).IsTrue()
	g.Assert(len(resp.Products)).Equal(expectedReturnedProductsCount)

	for i := 0; i < int(expectedReturnedProductsCount); i++ {
		g.Assert(resp.Products[i].Id).Equal(savedProducts[i].Id)
		g.Assert(resp.Products[i].Title).Equal(savedProducts[i].Title)
		g.Assert(resp.Products[i].Quantity).Equal(savedProducts[i].Quantity)
		g.Assert(resp.Products[i].Price).Equal(savedProducts[i].Price)
		g.Assert(len(resp.Products[i].Specification)).Equal(len(savedProducts[i].Specification))
		g.Assert(resp.Products[i].Specification[0].Name).Equal(savedProducts[i].Specification[0].Name)
	}
}

func init() {
	GET_PRODUCTS_GENERATED_COUNT = 100
}
