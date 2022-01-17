import React, { useState } from 'react';
import Product from '../interfaces/Product';
import ProductCard from './ProductCard';
import * as ProductService from '../services/ProductService'

const ProductListing = () => {
   const [products, setProducts] = useState<Product[]>(ProductService.getProducts());

   return (
      <>
         <div className='ui stackable equal width grid'>
            {
               products.map((element, idx, arr) => {
                  return (
                     <div className='column'>
                        <ProductCard
                           product={element} />
                     </div>
                  )
               })
            }
         </div>
      </>
   )
}

export default ProductListing;
