import Product from "../interfaces/Product";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";

interface GetProductsResponse {
   total: number,
   products: Product[]
}

const getProductModels = async (offset: number = 0, limit: number = 10): Promise<GetProductsResponse> => {
   return callService<GetProductsResponse>(`product-model?offset=${offset}&limit=${limit}`)
      .then(response => {
         return {
            total: response.data?.total || 0,
            products: response.data?.products || [],
         }
      })
}

const getProductModel = async (id: number): Promise<Product | undefined> => {
   return callService<Product | undefined>(`product-model/${id}`)
      .then(response => {
         return response.data
      })
}

const createProductModel = async (product: Product): Promise<ResponseContent<any>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(product),
   }

   return callService<any>(`product-model`, requestOptions)
}

export { getProductModel as getProduct, getProductModels as getProducts, createProductModel as createProduct }