import React from "react";
import { Button, Card, Grid, Icon, Image } from "semantic-ui-react";
import ProductModel from "../../interfaces/ProductModel";
import PurchaseOrder from "../../interfaces/PurchaseOrder";
import AddProductPane from "../../pages/AdminDashboard/AddProductPane";
import PurchaseButton from "../PurchaseButton";

interface PurchaseCardProps {
   purchaseOrder: PurchaseOrder
}

const PurchaseCard = (props: PurchaseCardProps) => {
   return (
      <Card fluid>
         <Card.Content>
            <Card.Header>{props.purchaseOrder.status}</Card.Header>
            <Grid columns={2}>
               <Grid.Column textAlign="right">
                  <Icon name="ethereum" />
                  {props.purchaseOrder.price} wei
               </Grid.Column>
               <Grid.Column textAlign="left">
                  <Icon name="calendar alternate outline" />
                  {props.purchaseOrder.date}
               </Grid.Column>
            </Grid>
         </Card.Content>
         <Card.Content>
            <Grid divided="vertically">
               {
                  props.purchaseOrder.product.map((elem, idx, arr) => {
                     return <Grid.Row>
                        <Grid.Column width={3}>
                           <img src='https://react.semantic-ui.com/images/wireframe/image.png' />
                        </Grid.Column>
                        <Grid.Column width={10}>
                           <a href={`/product/${elem.id}`} >{elem.productModel?.title}</a>
                        </Grid.Column>
                        <Grid.Column width={3}>
                           {elem.price} wei
                        </Grid.Column>
                     </Grid.Row>
                  })
               }
            </Grid>
         </Card.Content>
      </Card>
   )
}

export default PurchaseCard;