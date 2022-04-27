package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"rafal-kalinowski.pl/adapter/api"
	"rafal-kalinowski.pl/config"
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
			respRecorder := ServeTestRequest(router, "POST", PRODUCT_BASE_URI, []byte("{}"), nil)
			savedProducts, count := productRepository.GetProducts(10, 0)

			g.Assert(respRecorder.Code).Equal(http.StatusBadRequest)
			g.Assert(count == 0).IsTrue()
			g.Assert(len(savedProducts)).Equal(0)
		})

		g.It("Should create product from valid request", func() {
			respRecorder := ServeTestRequest(router, "POST", PRODUCT_BASE_URI, []byte(productJson), nil)
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

		g.It("Should get single existing product", func() {
			respRecorder := ServeTestRequest(router, "GET", PRODUCT_BASE_URI+"/1", nil, nil)
			product := getSingleProductResponseFromByteArray(respRecorder.Body)

			g.Assert(respRecorder.Code).Equal(http.StatusOK)
			g.Assert(product.Id).IsNotZero()
		})

		g.It("Should fail on wrong product id", func() {
			respRecorder := ServeTestRequest(router, "GET", PRODUCT_BASE_URI+"/sdf", nil, nil)
			product := getSingleProductResponseFromByteArray(respRecorder.Body)

			g.Assert(respRecorder.Code).Equal(http.StatusBadRequest)
			g.Assert(product.Id).IsZero()
		})

		g.It("Should fail on not existing product id", func() {
			respRecorder := ServeTestRequest(router, "GET", PRODUCT_BASE_URI+"/999999", nil, nil)
			product := getSingleProductResponseFromByteArray(respRecorder.Body)

			g.Assert(respRecorder.Code).Equal(http.StatusNotFound)
			g.Assert(product.Id).IsZero()
		})
	})
}

func TestImage(t *testing.T) {
	g := Goblin(t)

	CREATE_IMAGE_BASE_URI := "/product/%d/image"
	CONTENT_TYPE_HEADER_NAME := "Content-Type"
	CREATE_IMAGE_HEADERS := map[string]string{}

	g.Describe("Test Image", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.JustBeforeEach(fillProductsRepo)
		g.After(Cleanup)

		g.It("Should fail when trying to create image for nonexisting product", func() {
			g.Timeout(time.Hour)

			bytePayload, contentType, err := getMimeFileBytesFromName("../testdata/", "image.png")
			if err != nil {
				g.Fail(err)
			}

			CREATE_IMAGE_HEADERS[CONTENT_TYPE_HEADER_NAME] = contentType

			respRecorder := ServeTestRequest(router, "POST",
				fmt.Sprintf(CREATE_IMAGE_BASE_URI, 99999), *bytePayload, CREATE_IMAGE_HEADERS)

			g.Assert(respRecorder.Code).Equal(http.StatusNotFound)
		})

		g.It("Should create single image for a product", func() {
			g.Timeout(time.Hour)

			bytePayload, contentType, err := getMimeFileBytesFromName("../testdata/", "image.png")

			g.Assert(err).IsNil()

			CREATE_IMAGE_HEADERS[CONTENT_TYPE_HEADER_NAME] = contentType

			respRecorder := ServeTestRequest(router, "POST",
				fmt.Sprintf(CREATE_IMAGE_BASE_URI, 1), *bytePayload, CREATE_IMAGE_HEADERS)

			image := getImageResponseFromByteArray(respRecorder.Body)

			g.Assert(respRecorder.Code).Equal(http.StatusOK)
			g.Assert(image.Id > 0).IsTrue()
			g.Assert(len(image.Name) > 0).IsTrue()
			g.Assert(image.Url).Equal(config.IMAGE_REPO_BASE_URI + image.Name)

			product, err := productRepository.GetProduct(1)

			g.Assert(err).IsNil()
			g.Assert(len(product.Img)).Equal(1)
			g.Assert(product.Img[0].Id == 1).IsTrue()
			g.Assert(product.Img[0].Name).Equal(image.Name)
		})

		g.It("Should create multiple images for a products", func() {
			g.Timeout(time.Hour)

			for i := 0; i < GET_PRODUCTS_GENERATED_COUNT; i++ {
				bytePayload, contentType, err := getMimeFileBytesFromName("../testdata/", "image.png")

				g.Assert(err).IsNil()

				CREATE_IMAGE_HEADERS[CONTENT_TYPE_HEADER_NAME] = contentType

				respRecorder := ServeTestRequest(router, "POST",
					fmt.Sprintf(CREATE_IMAGE_BASE_URI, 1), *bytePayload, CREATE_IMAGE_HEADERS)

				image := getImageResponseFromByteArray(respRecorder.Body)

				g.Assert(respRecorder.Code).Equal(http.StatusOK)
				g.Assert(image.Id > 0).IsTrue()
				g.Assert(len(image.Name) > 0).IsTrue()
				g.Assert(image.Url).Equal(config.IMAGE_REPO_BASE_URI + image.Name)

				product, err := productRepository.GetProduct(1)

				g.Assert(err).IsNil()
				g.Assert(len(product.Img)).Equal(i + 1)
				g.Assert(product.Img[i].Id == uint(i+1)).IsTrue()
				g.Assert(product.Img[i].Name).Equal(image.Name)
			}
		})

		g.It("Should create single image for multiple products", func() {
			g.Timeout(time.Hour)

			for i := 0; i < GET_PRODUCTS_GENERATED_COUNT; i++ {
				bytePayload, contentType, err := getMimeFileBytesFromName("../testdata/", "image.png")

				g.Assert(err).IsNil()

				CREATE_IMAGE_HEADERS[CONTENT_TYPE_HEADER_NAME] = contentType

				respRecorder := ServeTestRequest(router, "POST",
					fmt.Sprintf(CREATE_IMAGE_BASE_URI, i+1), *bytePayload, CREATE_IMAGE_HEADERS)

				image := getImageResponseFromByteArray(respRecorder.Body)

				g.Assert(respRecorder.Code).Equal(http.StatusOK)
				g.Assert(image.Id > 0).IsTrue()
				g.Assert(len(image.Name) > 0).IsTrue()
				g.Assert(image.Url).Equal(config.IMAGE_REPO_BASE_URI + image.Name)

				product, err := productRepository.GetProduct(uint(i + 1))

				g.Assert(err).IsNil()
				g.Assert(len(product.Img)).Equal(1)
				g.Assert(product.Img[0].Id == uint(i+1)).IsTrue()
				g.Assert(product.Img[0].Name).Equal(image.Name)
			}
		})
	})
}

func fillProductsRepo() {
	count := GET_PRODUCTS_GENERATED_COUNT

	for i := 0; i < count; i++ {
		product := &model.ProductModel{
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

func getMimeFileBytesFromName(dir, name string) (*[]byte, string, error) {
	var bBuffer bytes.Buffer
	multipartFileWriter := multipart.NewWriter(&bBuffer)

	imageFile, err := os.Open(config.IMAGE_LOCAL_STORAGE + dir + name)
	if err != nil {
		return nil, "", err
	}

	defer imageFile.Close()
	fw, err := multipartFileWriter.CreateFormFile("upload", name)
	if err != nil {
		return nil, "", err
	}

	if _, err = io.Copy(fw, imageFile); err != nil {
		return nil, "", err
	}

	contentType := multipartFileWriter.FormDataContentType()
	multipartFileWriter.Close()

	bytes, err := ioutil.ReadAll(&bBuffer)
	return &bytes, contentType, err
}

func getSingleProductResponseFromByteArray(data *bytes.Buffer) *model.ProductModelDTO {
	dataBytes, _ := ioutil.ReadAll(data)
	respData := model.ProductModelDTO{}
	json.Unmarshal(dataBytes, &respData)

	return &respData
}

func getImageResponseFromByteArray(data *bytes.Buffer) *model.ImageDTO {
	dataBytes, _ := ioutil.ReadAll(data)
	respData := model.ImageDTO{}
	json.Unmarshal(dataBytes, &respData)

	return &respData
}

func getProductsResponseFromByteArray(data *bytes.Buffer) *api.GetProductsResponse {
	dataBytes, _ := ioutil.ReadAll(data)
	respData := api.GetProductsResponse{}
	json.Unmarshal(dataBytes, &respData)

	return &respData
}

func assertPaginatedProductsAreValid(g *G, uri string, offset, limit int) {
	respRecorder := ServeTestRequest(router, "GET", uri, nil, nil)
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
