interface Specification {
   name: string
   value: string
}

export default interface Product {
   id: number | undefined
   imgUrl: string[]
   title: string
   description: string
   price: number
   quantity: number
   specification: Specification[]
}
