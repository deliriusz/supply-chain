import React from "react";
import { Button, Card, Icon, Image } from "semantic-ui-react";
import Product from "../../interfaces/Product";
import PurchaseButton from "../PurchaseButton";

interface ProductCardProps {
   product: Product
}

const ProductCard = (props: ProductCardProps) => {
   return (
      <Card>
         <Card.Content>
            <Image
               src={props.product.imgUrl && props.product.imgUrl[0]}
            />
            <Card.Header>{props.product.title}</Card.Header>
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