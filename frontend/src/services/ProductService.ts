import Product from "../interfaces/Product";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";

interface GetProductsResponse {
   total: number,
   products: Product[]
}

const getProducts = async (offset: number = 0, limit: number = 10): Promise<GetProductsResponse> => {
   return callService<GetProductsResponse>(`product?offset=${offset}&limit=${limit}`)
      .then(response => {
         return response.data || { total: 0, products: [] }
      })
}

const getProduct = async (id: number): Promise<Product | undefined> => {
   return callService<Product | undefined>(`product/${id}`)
      .then(response => {
         return response.data
      })
}

const createProduct = async (product: Product): Promise<ResponseContent<any>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(product),
   }

   return callService<any>(`product`, requestOptions)
}

export { getProduct, getProducts, createProduct }