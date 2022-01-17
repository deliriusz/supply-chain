import Product from "../interfaces/Product";

let productList: Product[] = [
   {
      id: 1,
      imgUrl: [
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
      ],
      title: 'Bulbulator',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce porta consectetur magna sed elementum. Maecenas sed dignissim erat. Proin ullamcorper libero vel arcu vehicula ultricies. Aliquam finibus massa vel luctus porta. Pellentesque vehicula tortor in vehicula ultricies. Maecenas semper faucibus quam, in fringilla ipsum consequat sit amet. Proin eget rhoncus eros. Suspendisse at lorem dapibus, vehicula sem eget, pretium ligula. Maecenas interdum risus id malesuada scelerisque. Nam cursus pharetra nisi sed rhoncus. Aliquam fringilla rhoncus velit quis malesuada. Etiam tempus mauris ut hendrerit suscipit. Vestibulum enim eros, mattis in consequat in, mollis eu justo.',
      price: 1000,
      specification: [
         {
            name: "Size",
            value: "1\" x 2\""
         },
         {
            name: "Weight",
            value: "6 ounces"
         },
         {
            name: "Color",
            value: "red"
         },
         {
            name: "Power",
            value: "100 kWh"
         },
      ]
   },
   {
      id: 2,
      imgUrl: [
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
         `${process.env.PUBLIC_URL}/favicon.ico`,
         `${process.env.PUBLIC_URL}/logo192.png`,
      ],
      title: 'Nice & Dandy',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce porta consectetur magna sed elementum. Maecenas sed dignissim erat. Proin ullamcorper libero vel arcu vehicula ultricies. Aliquam finibus massa vel luctus porta. Pellentesque vehicula tortor in vehicula ultricies. Maecenas semper faucibus quam, in fringilla ipsum consequat sit amet. Proin eget rhoncus eros. Suspendisse at lorem dapibus, vehicula sem eget, pretium ligula. Maecenas interdum risus id malesuada scelerisque. Nam cursus pharetra nisi sed rhoncus. Aliquam fringilla rhoncus velit quis malesuada. Etiam tempus mauris ut hendrerit suscipit. Vestibulum enim eros, mattis in consequat in, mollis eu justo.',
      price: 222,
      specification: [
         {
            name: "Size",
            value: "3\" x 4\""
         },
         {
            name: "Weight",
            value: "9 ounces"
         },
         {
            name: "Color",
            value: "chroma blue"
         },
         {
            name: "Power",
            value: "256 kWh"
         },
      ]
   }
]

const getProducts = (offset: number = 0, limit: number = 10): Product[] => {
   return productList
}

const getProduct = (id: number): Product | undefined => {
   return productList.filter((v, i, a) => v.id === id)[0]
}

export { getProduct, getProducts }
