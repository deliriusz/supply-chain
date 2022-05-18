package repository_test

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"rafal-kalinowski.pl/config"
	"rafal-kalinowski.pl/domain/model"
)

func TestCreateProduct(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test CreateProduct", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		g.It("Should create a product without an error", func() {
			if err := productRepo.CreateProductModel(createRandomProductModel()); err != nil {
				g.Fail(err)
			}

			testProduct := createRandomProduct(0)

			if err := productRepo.CreateProduct(testProduct); err != nil {
				g.Fail(err)
			}
		})

		g.It("Should return created product with the same data", func() {
			testProduct := createRandomProduct(0)

			if err := productRepo.CreateProduct(testProduct); err != nil {
				g.Fail(err)
			}

			createdProduct, err := productRepo.GetProduct(1)

			g.Assert(err).IsNil()

			g.Assert(createdProduct.Id == 1).IsTrue()
			g.Assert(createdProduct.Model.Id).Equal(testProduct.Model.Id)
			g.Assert(createdProduct.Model.BasePrice).Equal(testProduct.Model.BasePrice)
			g.Assert(createdProduct.Model.Description).Equal(testProduct.Model.Description)
			g.Assert(createdProduct.Model.Quantity).Equal(testProduct.Model.Quantity)
			g.Assert(createdProduct.Model.Title).Equal(testProduct.Model.Title)
			g.Assert(createdProduct.Owner).Equal(testProduct.Owner)
			g.Assert(createdProduct.Price).Equal(testProduct.Price)
			g.Assert(createdProduct.State).Equal(testProduct.State)
		})

		g.It("Should create multiple products for the same ProductModel", func() {
			productRepo.CreateProductModel(createRandomProductModel())
			createdProductCount := 100
			for i := 0; i < createdProductCount; i++ {
				testProduct := createRandomProduct(1)

				if err := productRepo.CreateProduct(testProduct); err != nil {
					g.Fail(err)
				}
			}

			createdProducts, count := productRepo.GetProducts(10, uint(createdProductCount-10))

			g.Assert(count == uint(createdProductCount)).IsTrue()
			g.Assert(len(*createdProducts)).Equal(10)

			for _, v := range *createdProducts {
				g.Assert(v.Id >= 90).IsTrue()
			}
		})

		g.It("Should create multiple products for multiple ProductModel", func() {
			createdProductCount := 100
			for i := 0; i < createdProductCount; i++ {
				testProduct := createRandomProduct(0)

				if err := productRepo.CreateProduct(testProduct); err != nil {
					g.Fail(err)
				}
			}

			createdProducts, count := productRepo.GetProducts(10, uint(createdProductCount-10))

			g.Assert(count == uint(createdProductCount)).IsTrue()
			g.Assert(len(*createdProducts)).Equal(10)

			for _, v := range *createdProducts {
				g.Assert(v.Id >= 90).IsTrue()
				g.Assert(v.Id).Equal(v.Model.Id)
			}
		})
	})
}

func TestCreateProductModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test CreateProductModel", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		g.It("Should create a product without an error", func() {
			testProduct := createRandomProductModel()

			if err := productRepo.CreateProductModel(testProduct); err != nil {
				g.Fail(err)
			}
		})

		g.It("Should return created product with the same data", func() {
			testProduct := createRandomProductModel()

			if err := productRepo.CreateProductModel(testProduct); err != nil {
				g.Fail(err)
			}

			createdProducts, count := productRepo.GetProductModels(10, 0)
			firstCreatedProduct := createdProducts[0]

			// for some reason it fails with message 0x1 is not equal to 1 when testing for equality
			g.Assert(count == 1).IsTrue()
			g.Assert(len(createdProducts)).Equal(1)

			g.Assert(firstCreatedProduct.Id == 1).IsTrue()
			g.Assert(firstCreatedProduct.Quantity).Equal(testProduct.Quantity)
			g.Assert(firstCreatedProduct.Description).Equal(testProduct.Description)
			g.Assert(firstCreatedProduct.BasePrice).Equal(testProduct.BasePrice)
			g.Assert(firstCreatedProduct.Title).Equal(testProduct.Title)

			g.Assert(len(firstCreatedProduct.Specification)).Equal(len(testProduct.Specification))

			for i := 0; i < len(firstCreatedProduct.Specification); i++ {
				g.Assert(firstCreatedProduct.Specification[i].ProductId).Equal(testProduct.Specification[i].ProductId)
				g.Assert(firstCreatedProduct.Specification[i].Name).Equal(testProduct.Specification[i].Name)
				g.Assert(firstCreatedProduct.Specification[i].Value).Equal(testProduct.Specification[i].Value)
			}
		})

		g.It("Should create multiple products", func() {
			createdProductCount := 100
			for i := 0; i < createdProductCount; i++ {
				testProduct := createRandomProductModel()

				if err := productRepo.CreateProductModel(testProduct); err != nil {
					g.Fail(err)
				}
			}

			createdProducts, count := productRepo.GetProductModels(10, uint(createdProductCount-10))

			g.Assert(count == uint(createdProductCount)).IsTrue()
			g.Assert(len(createdProducts)).Equal(10)

			for _, v := range createdProducts {
				g.Assert(v.Id >= 90).IsTrue()
			}
		})
	})
}

func TestImage(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test Image", func() {
		g.JustBeforeEach(Cleanup)
		g.JustBeforeEach(Setup)
		g.After(Cleanup)

		g.It("Should create image for existing product", func() {
			testProduct := createRandomProductModel()

			if err := productRepo.CreateProductModel(testProduct); err != nil {
				g.Fail(err)
			}

			fileReader, getFileReaderErr := getReaderFromName("../testdata/image.png")
			if getFileReaderErr != nil {
				g.Fail(getFileReaderErr)
			}

			createdImage, createImageErr := productRepo.CreateImage(1, fileReader)
			if createImageErr != nil {
				g.Fail(createImageErr)
			}

			g.Assert(len(createdImage.Name) > 0).IsTrue()

			savedFileName, savedFileDirectory, savedFileReader, getImageErr := productRepo.GetImage(createdImage.Name)

			g.Assert(len(savedFileName) > 0).IsTrue()
			g.Assert(savedFileDirectory).Equal(config.IMAGE_LOCAL_STORAGE)

			_, errTestRead := savedFileReader.ReadByte()
			g.Assert(savedFileReader.Buffered() > 0).IsTrue()
			g.Assert(errTestRead).IsNil()
			g.Assert(getImageErr).IsNil()
		})

		g.It("Should assign image to product", func() {
			testProduct := createRandomProductModel()

			if err := productRepo.CreateProductModel(testProduct); err != nil {
				g.Fail(err)
			}

			fileReader, getFileReaderErr := getReaderFromName("../testdata/image.png")
			if getFileReaderErr != nil {
				g.Fail(getFileReaderErr)
			}

			createdImage, createImageErr := productRepo.CreateImage(1, fileReader)
			if createImageErr != nil {
				g.Fail(createImageErr)
			}

			g.Assert(len(createdImage.Name) > 0).IsTrue()

			productWithImage, getProductErr := productRepo.GetProductModel(1)
			if getProductErr != nil {
				g.Fail(getProductErr)
			}

			g.Assert(len(productWithImage.Img)).Equal(1)
			g.Assert(productWithImage.Img[0].Id > 0).IsTrue()
			g.Assert(productWithImage.Img[0].ProductId == 1).IsTrue()
			g.Assert(productWithImage.Img[0].Name).Equal(createdImage.Name)
		})

		g.It("Should fail during image creation when product does not exist", func() {
			fileReader, getFileReaderErr := getReaderFromName("../testdata/image.png")
			if getFileReaderErr != nil {
				g.Fail(getFileReaderErr)
			}

			createdImage, createImageErr := productRepo.CreateImage(999, fileReader)
			g.Assert(createImageErr).IsNotNil()

			g.Assert(len(createdImage.Name)).Equal(0)
		})

		g.It("Should create multiple images to a product", func() {
			createdImagesCount := 5
			testProduct := createRandomProductModel()

			if err := productRepo.CreateProductModel(testProduct); err != nil {
				g.Fail(err)
			}

			for i := 0; i < createdImagesCount; i++ {
				fileReader, getFileReaderErr := getReaderFromName("../testdata/image.png")
				if getFileReaderErr != nil {
					g.Fail(getFileReaderErr)
				}

				createdImage, createImageErr := productRepo.CreateImage(1, fileReader)
				if createImageErr != nil {
					g.Fail(createImageErr)
				}

				createdImageFile, createdImageFileErr := os.Open(config.IMAGE_LOCAL_STORAGE + "/" + createdImage.Name)
				if createdImageFileErr != nil {
					g.Fail(createdImageFileErr)
				}
				createdImageFile.Close()
			}

			productWithImages, getProductErr := productRepo.GetProductModel(1)
			if getProductErr != nil {
				g.Fail(getProductErr)
			}

			g.Assert(len(productWithImages.Img)).Equal(createdImagesCount)
			g.Assert(productWithImages.Img[0].Id > 0).IsTrue()
			g.Assert(productWithImages.Img[0].ProductId == 1).IsTrue()
			g.Assert(productWithImages.Img[createdImagesCount-1].Id > 0).IsTrue()
			g.Assert(productWithImages.Img[createdImagesCount-1].ProductId == 1).IsTrue()
		})
	})
}

func createRandomPurchase() *model.PurchaseOrder {
	createRandomProductModel()
	createRandomProductModel()
	purchaseProduct1 := createRandomProduct(1)
	purchaseProduct2 := createRandomProduct(2)

	poTime := time.Now().Add(time.Hour * (time.Duration(-rand.Intn(5000))))
	return &model.PurchaseOrder{
		UserId: "asdf",
		Date:   poTime,
		Product: []model.Product{
			*purchaseProduct1,
			*purchaseProduct2,
		},
	}
}

func createRandomProduct(modelId uint) *model.Product {
	if modelId < 1 {
		productRepo.CreateProductModel(createRandomProductModel())
		modelId = 1
	}

	productModel, _ := productRepo.GetProductModel(modelId)

	return &model.Product{
		Model: productModel,
		State: model.InProduction,
		Owner: "",
		Price: productModel.BasePrice,
	}
}

func createRandomProductModel() *model.ProductModel {
	randomBaseNumber := rand.Int()
	randomBaseString := strconv.Itoa(randomBaseNumber)

	return &model.ProductModel{
		Title:       randomBaseString + " title",
		Description: randomBaseString + " description",
		BasePrice:   uint(randomBaseNumber),
		Quantity:    uint(randomBaseNumber),
		Specification: []model.Specification{
			{
				Name:  "size",
				Value: randomBaseString + "x" + randomBaseString,
			},
			{
				Name:  "model",
				Value: randomBaseString,
			},
		},
	}
}

func getReaderFromName(name string) (*bufio.Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil
}
