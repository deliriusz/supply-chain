import Product from "./Product"

export default interface PurchaseOrder {
   id: string,
   userId: string
   product: Product[],
   price: number,
   date: string,
   status: string
}