import React, { useState } from 'react';
import Product from '../interfaces/Product';
import ProductCard from './ProductCard';
import * as ProductService from '../services/ProductService'
import { Grid } from 'semantic-ui-react';

const ProductListing = () => {
   const [products, setProducts] = useState<Product[]>(ProductService.getProducts());

   return (
      <>
         <Grid stackable columns={'equal'}>
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
         </Grid>
      </>
   )
}

export default ProductListing;
