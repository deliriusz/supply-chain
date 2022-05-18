import PurchaseOrder from "../interfaces/PurchaseOrder";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";

interface GetPurchaseOrdersResponse {
   total: number,
   products: PurchaseOrder[]
}

//TODO: map userId
const getPurchaseOrdersForUser = async (userId: string, offset: number = 0, limit: number = 10): Promise<GetPurchaseOrdersResponse> => {
   return {
      total: 12,
      products: [
         {
            id: "1",
            userId: "",
            date: "2020-05-13 15:23:58",
            price: 100,
            product: [
               {
                  price: 90,
                  state: "IN_PROGRESS",
                  productModel: {
                     description: "",
                     title: "product title",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               },
               {
                  price: 10,
                  state: "PAYED",
                  productModel: {
                     description: "",
                     title: "product title 2",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               }
            ],
            status: "IN_PROGRESS"

         },
         {
            id: "2",
            userId: "",
            date: "2021-12-24 09:30:00",
            price: 250,
            product: [
               {
                  price: 250,
                  state: "DELIVERED",
                  productModel: {
                     description: "",
                     title: "product title 3",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               }
            ],
            status: "DELIVERED"

         }
      ]
   }

   return callService<GetPurchaseOrdersResponse>(`purchase?offset=${offset}&limit=${limit}`)
      .then(response => {
         return response.data || { total: 0, products: [] }
      })
}

const getPurchaseOrder = async (id: number): Promise<PurchaseOrder | undefined> => {
   return callService<PurchaseOrder>(`purchase/${id}`)
      .then(response => {
         return response.data
      })
}

const createPurchaseOrder = async (purchaseOrder: PurchaseOrder): Promise<ResponseContent<any>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(purchaseOrder),
   }

   return callService<any>('purchase', requestOptions)
}

export { getPurchaseOrder, getPurchaseOrdersForUser, createPurchaseOrder }
