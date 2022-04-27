package model

import (
	"fmt"

	"rafal-kalinowski.pl/config"
)

func ToPurchaseDTO(purchase PurchaseOrder) PurchaseOrderDTO {
	return PurchaseOrderDTO{
		Id:      purchase.Id,
		UserId:  purchase.UserId,
		Product: purchase.Product,
		Price:   purchase.Price,
		Date:    purchase.Date,
		Status:  purchase.Status,
	}
}

func ToPurchase(purchase PurchaseOrderDTO) PurchaseOrder {
	return PurchaseOrder{
		Id:      purchase.Id,
		UserId:  purchase.UserId,
		Product: purchase.Product,
		Price:   purchase.Price,
		Date:    purchase.Date,
		Status:  purchase.Status,
	}
}

func ToLoginDTO(login Login) LoginDTO {
	return LoginDTO{
		Address:   login.Address,
		Role:      login.Role.String(),
		ExpiresAt: login.ExpiresAt,
		TTL:       login.TTL,
	}
}

func ToProductDTO(product Product) ProductDTO {
	var imgDtos []ImageDTO
	var specDtos []SpecificationDTO

	for _, img := range product.Img {
		imgDtos = append(imgDtos, ToImageDTO(img))
	}

	for _, spec := range product.Specification {
		specDtos = append(specDtos, ToSpecificationDTO(spec))
	}

	return ProductDTO{
		Id:            product.Id,
		Title:         product.Title,
		Description:   product.Description,
		Price:         product.Price,
		Quantity:      product.Quantity,
		Img:           imgDtos,
		Specification: specDtos,
	}
}

func ToSpecificationDTO(specification Specification) SpecificationDTO {
	return SpecificationDTO{
		Name:  specification.Name,
		Value: specification.Value,
	}
}

func ToImageDTO(image Image) ImageDTO {
	return ImageDTO{
		Id:   image.Id,
		Name: image.Name,
		Url:  fmt.Sprintf("%s%s", config.IMAGE_REPO_BASE_URI, image.Name),
	}
}

func ToImage(imgDto ImageDTO) Image {
	return Image{
		Id:        imgDto.Id,
		ProductId: imgDto.ProductId,
		Name:      imgDto.Name,
	}
}

func ToSpecification(specDto SpecificationDTO) Specification {
	return Specification{
		Name:  specDto.Name,
		Value: specDto.Value,
	}
}

func ToProduct(productDto ProductDTO) Product {
	var imgs []Image
	var specs []Specification

	for _, img := range productDto.Img {
		imgs = append(imgs, ToImage(img))
	}

	for _, spec := range productDto.Specification {
		specs = append(specs, ToSpecification(spec))
	}

	return Product{
		Id:            productDto.Id,
		Title:         productDto.Title,
		Description:   productDto.Description,
		Price:         productDto.Price,
		Quantity:      productDto.Quantity,
		Img:           imgs,
		Specification: specs,
	}
}
