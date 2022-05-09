import ProductModel from "./ProductModel"

export default interface Product {
   id?: number
   productModel: ProductModel
   state: string
   owner?: string
   price: number
}
