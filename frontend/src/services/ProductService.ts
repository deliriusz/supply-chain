import Product from "../interfaces/Product";
import ResponseContent from "../interfaces/ResponseContent";

interface GetProductsResponse {
   total: number,
   products: Product[]
}

const getProducts = async (offset: number = 0, limit: number = 10): Promise<GetProductsResponse> => {
   const requestOptions: RequestInit = {
      method: "GET",
      mode: "cors",
      credentials: "include",
   }

   var responseContent: ResponseContent = { isOk: false, status: 500 };

   await fetch(`${process.env.REACT_APP_BACKEND_URL}/product?offset=${offset}&limit=${limit}`, requestOptions)
      .then(async response => {
         const responseData: any = await response.json()

         responseContent = { data: responseData, isOk: response.ok, status: response.status }
         responseContent.data = responseData
         responseContent.isOk = response.ok
         responseContent.status = response.status

         if (!response.ok) {
            return Promise.reject(responseContent)
         }

      }).catch(err => {
         console.log(err)
      })

   return responseContent.data
}

const getProduct = async (id: number): Promise<Product | undefined> => {
   const requestOptions: RequestInit = {
      method: "GET",
      mode: "cors",
      credentials: "include",
   }

   var responseContent: ResponseContent = { isOk: false, status: 500 };

   await fetch(`${process.env.REACT_APP_BACKEND_URL}/product/${id}`, requestOptions)
      .then(async response => {
         const responseData: any = await response.json()

         responseContent = { data: responseData, isOk: response.ok, status: response.status }
         responseContent.data = responseData
         responseContent.isOk = response.ok
         responseContent.status = response.status

         if (!response.ok) {
            return Promise.reject(responseContent)
         }

      }).catch(err => {
         console.log(err)
      })

   return responseContent.data
}

const createProduct = async (product: Product): Promise<ResponseContent> => {
   const requestOptions: RequestInit = {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(product),
      credentials: "include",
   }

   var responseContent: ResponseContent = { isOk: false, status: 500 };

   await fetch(`${process.env.REACT_APP_BACKEND_URL}/product`, requestOptions)
      .then(async response => {
         const responseData: any = await response.json()

         responseContent = { data: responseData, isOk: response.ok, status: response.status }
         responseContent.data = responseData
         responseContent.isOk = response.ok
         responseContent.status = response.status

         if (!response.ok) {
            return Promise.reject(responseContent)
         }

      }).catch(err => {
         // response = err
      })

   return responseContent
}

export type { ResponseContent }
export { getProduct, getProducts, createProduct }
