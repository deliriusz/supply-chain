import Product from "../interfaces/Product";

interface ResponseContent {
   data?: any,
   isOk: boolean,
   status: number,
}

const getProducts = async (offset: number = 0, limit: number = 10): Promise<Product[]> => {
   const requestOptions = {
      method: "GET",
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
   const requestOptions = {
      method: "GET",
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
   const requestOptions = {
      method: "POST",
      body: JSON.stringify(product),
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
