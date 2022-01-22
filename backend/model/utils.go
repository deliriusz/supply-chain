package model

import (
	"fmt"
)

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
		Id:                product.Id,
		Title:             product.Title,
		Description:       product.Description,
		Price:             product.Price,
		AvailableQuantity: product.AvailableQuantity,
		Img:               imgDtos,
		Specification:     specDtos,
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
		Id:  image.Id,
		Url: fmt.Sprintf("/image/%s", image.ImageName),
	}
}

func ToImage(imgDto ImageDTO) Image {
	return Image{
		Id:        imgDto.Id,
		ProductId: imgDto.ProductId,
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
		Id:                productDto.Id,
		Title:             productDto.Title,
		Description:       productDto.Description,
		Price:             productDto.Price,
		AvailableQuantity: productDto.AvailableQuantity,
		Img:               imgs,
		Specification:     specs,
	}
}
