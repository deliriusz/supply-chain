interface Specification {
   name: string
   value: string | number
}

export default interface Product {
   id: number
   imgUrl: string[]
   title: string
   description: string
   price: number
   specification: Specification[]
}
