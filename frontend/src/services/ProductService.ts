import ProductModel from "../interfaces/ProductModel";
import Product from "../interfaces/Product";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";

interface GetProductModelsResponse {
   total: number,
   products: ProductModel[]
}

interface GetProductsResponse {
   total: number,
   products: Product[]
}

const getProductModels = async (offset: number = 0, limit: number = 10): Promise<GetProductModelsResponse> => {
   return callService<GetProductModelsResponse>(`product-model?offset=${offset}&limit=${limit}`)
      .then(response => {
         return {
            total: response.data?.total || 0,
            products: response.data?.products || [],
         }
      })
}

const getProductModel = async (id: number): Promise<ProductModel | undefined> => {
   return callService<ProductModel | undefined>(`product-model/${id}`)
      .then(response => {
         return response.data
      })
}

const createProductModel = async (product: ProductModel): Promise<ResponseContent<any>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(product),
   }

   return callService<any>(`product-model`, requestOptions)
}

const createImage = async (productId: number, file: File): Promise<ResponseContent<any>> => {
   const data = new FormData()
   data.append("upload", file)
   const requestOptions: RequestInit = {
      method: "POST",
      body: data
   }

   return callService<any>(`product-model/${productId}/image`, requestOptions)
}

const getProduct = async (id: number): Promise<Product | undefined> => {
   return callService<Product | undefined>(`product/${id}`)
      .then(response => {
         return response.data
      })
}

const getProductsForUser = async (user: string, offset: number = 0, limit: number = 10): Promise<GetProductsResponse | undefined> => {
   return callService<GetProductsResponse | undefined>(`product/user/${user}`)
      .then(response => {
         return response.data
      })
}

export { getProductModel, getProductModels, createProductModel, getProduct, getProductsForUser, createImage }