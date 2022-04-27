import React, { useEffect, useState } from "react";
import { Divider, Grid, Header, Icon, Segment, Table } from "semantic-ui-react";
import Product from "../../interfaces/Product";
import { useParams } from 'react-router-dom';
import * as ProductService from '../../services/ProductService'
import ImageCarousel from "../../components/ImageCarousel";
import PurchaseButton from "../../components/PurchaseButton";

const ProductDetails = () => {
   let { productId } = useParams()
   useEffect(() => {
      ProductService.getProduct(parseInt(productId!))
         .then(product => setProduct(product))
   }, [productId])
   let [product, setProduct] = useState<Product | undefined>()

   return (
      <Grid>
         <Grid.Row>
            <Grid.Column width={10}>
               <ImageCarousel images={product?.images} />
            </Grid.Column>
            <Grid.Column width={6}>

               <h1>{product?.title}</h1>
               <Divider horizontal>
                  <Header as='h4'>
                     <Icon name='ethereum' />
                     Buy
                  </Header>
               </Divider>
               <Segment>
                  <Grid columns={2} relaxed='very'>
                     <Grid.Column>
                        <h2>{product?.price} wei </h2>
                     </Grid.Column>
                     <Grid.Column>
                        <PurchaseButton
                           primary
                           size='small'
                           active={(!!product && !!product?.id)}
                           purchaseProductId={product?.id || -1}
                           purchaseAmount={product?.price || -1}
                        >
                        </PurchaseButton>
                     </Grid.Column>
                  </Grid>

                  <Divider vertical></Divider>
               </Segment>
               <Grid>

                  <Grid.Row>
                     <Grid.Column>
                        <Divider horizontal>
                           <Header as='h4'>
                              <Icon name='chart bar' />
                              Specifications
                           </Header>
                        </Divider>

                        <Table definition>
                           <Table.Body>
                              {
                                 product?.specification.map((val, idx, arr) => {
                                    return (<Table.Row key={idx}>
                                       <Table.Cell>{val.name}</Table.Cell>
                                       <Table.Cell>{val.value}</Table.Cell>
                                    </Table.Row>)
                                 })
                              }
                           </Table.Body>
                        </Table>
                     </Grid.Column>

                  </Grid.Row>
               </Grid>
            </Grid.Column>
         </Grid.Row>
         <Grid.Row>
            <Grid.Column>
               <Divider horizontal>
                  <Header as='h4'>
                     <Icon name='tag' />
                     Description
                  </Header>
               </Divider>
               <p>{product?.description}</p>
            </Grid.Column>
         </Grid.Row>
      </Grid>
   )
}

export default ProductDetails;