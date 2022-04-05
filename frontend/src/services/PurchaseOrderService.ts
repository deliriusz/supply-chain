import PurchaseOrder from "../interfaces/PurchaseOrder";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";

interface GetPurchaseOrdersResponse {
   total: number,
   products: PurchaseOrder[]
}

//TODO: map userId
const getPurchaseOrdersForUser = async (userId: number, offset: number = 0, limit: number = 10): Promise<GetPurchaseOrdersResponse> => {

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
