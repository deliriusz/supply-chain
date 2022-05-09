export interface Specification {
   name: string
   value: string
}

export interface Image {
   id: number
   productId: number
   name: string
   url: string
}

export default interface ProductModel {
   id: number | undefined
   images: Image[]
   title: string
   description: string
   price: number
   quantity: number
   specification: Specification[]
}
