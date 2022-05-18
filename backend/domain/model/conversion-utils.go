package model

import (
	"fmt"

	"rafal-kalinowski.pl/config"
)

func ToProductMetadataDTO(productMeta ProductMetadata) ProductMetadataDTO {
	return ProductMetadataDTO{
		Name:        productMeta.Name,
		Description: productMeta.Description,
		Image:       productMeta.Image,
		Attributes:  productMeta.Attributes,
		ExternalUri: productMeta.ExternalUri,
	}
}

func ToProductMetadata(productMeta ProductMetadataDTO) ProductMetadata {
	return ProductMetadata{
		Name:        productMeta.Name,
		Description: productMeta.Description,
		Image:       productMeta.Image,
		Attributes:  productMeta.Attributes,
		ExternalUri: productMeta.ExternalUri,
	}
}

func ToPurchaseDTO(purchase PurchaseOrder) PurchaseOrderDTO {
	totalPrice := uint(0)
	minimalState := Delivered
	for _, product := range purchase.Product {
		totalPrice += product.Price

		if minimalState > product.State {
			minimalState = product.State
		}
	}

	return PurchaseOrderDTO{
		Id:      purchase.Id,
		UserId:  purchase.UserId,
		Product: purchase.Product,
		Price:   totalPrice,
		Date:    purchase.Date,
		Status:  minimalState.String(),
	}
}

func ToPurchase(purchase PurchaseOrderDTO) PurchaseOrder {
	return PurchaseOrder{
		Id:      purchase.Id,
		UserId:  purchase.UserId,
		Product: purchase.Product,
		Date:    purchase.Date,
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
	return ProductDTO{
		Id:    product.Id,
		Model: ToProductModelDTO(product.Model),
		State: product.State.String(),
		Owner: product.Owner,
		Price: product.Price,
	}
}

func ToProductModelDTO(product ProductModel) ProductModelDTO {
	var imgDtos []ImageDTO
	var specDtos []SpecificationDTO

	for _, img := range product.Img {
		imgDtos = append(imgDtos, ToImageDTO(img))
	}

	for _, spec := range product.Specification {
		specDtos = append(specDtos, ToSpecificationDTO(spec))
	}

	return ProductModelDTO{
		Id:            product.Id,
		Title:         product.Title,
		Description:   product.Description,
		BasePrice:     product.BasePrice,
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
	state, _ := ProductStateFromString(productDto.State)

	return Product{
		Id:    productDto.Id,
		Model: ProductModel{Id: productDto.ModelId},
		State: state,
		Owner: productDto.Owner,
		Price: productDto.Price,
	}
}

func ToProductModel(productDto ProductModelDTO) ProductModel {
	var imgs []Image
	var specs []Specification

	for _, img := range productDto.Img {
		imgs = append(imgs, ToImage(img))
	}

	for _, spec := range productDto.Specification {
		specs = append(specs, ToSpecification(spec))
	}

	return ProductModel{
		Id:            productDto.Id,
		Title:         productDto.Title,
		Description:   productDto.Description,
		BasePrice:     productDto.BasePrice,
		Quantity:      productDto.Quantity,
		Img:           imgs,
		Specification: specs,
	}
}
