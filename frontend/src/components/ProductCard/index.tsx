import React from "react";
import { Button, Card, Icon, Image } from "semantic-ui-react";
import ProductModel from "../../interfaces/ProductModel";
import PurchaseButton from "../PurchaseButton";

interface ProductCardProps {
   product: ProductModel
}

const ProductCard = (props: ProductCardProps) => {
   return (
      <Card>
         <Card.Content>
            <Card.Header>{props.product.title}</Card.Header>
            <Image
               src={props.product.images && props.product.images[0].url}
            />
         </Card.Content>
         <Card.Content extra>
            <div className='ui two buttons'>
               <PurchaseButton
                  color="green"
                  size='small'
                  basic
                  purchaseProductId={props.product.id!}
                  purchaseAmount={props.product.price}
               >
               </PurchaseButton>
               <Button
                  as="a"
                  href={`/product/${props.product.id}`}
                  basic
                  color='blue'
               >
                  Details
               </Button>
            </div>
         </Card.Content>
      </Card>
   )
}

export default ProductCard;