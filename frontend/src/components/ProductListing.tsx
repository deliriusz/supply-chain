import React, { useEffect, useState } from 'react';
import Product from '../interfaces/Product';
import ProductCard from './ProductCard';
import * as ProductService from '../services/ProductService'
import { Grid, GridColumn } from 'semantic-ui-react';

const ProductListing = () => {
   useEffect(() => {
      ProductService.getProducts().then(products => setProducts(products.products))
   }, [])
   const [products, setProducts] = useState<Product[]>([]);

   return (
      <>
         <Grid stackable columns={'equal'}>
            {
               products.map((element, idx, arr) => {
                  return (
                     <GridColumn key={`product-${idx}`} width={5}>
                        <ProductCard
                           product={element} />
                     </GridColumn>
                  )
               })
            }
         </Grid>
      </>
   )
}

export default ProductListing;
