import React, { useState } from "react";
import { Button, Card, Divider, Form, FormProps, Grid, Icon, Segment } from "semantic-ui-react";
import PurchaseCard from "../../components/PurchaseCard";
import PurchaseOrder from "../../interfaces/PurchaseOrder";

const handleSearchInput = (event: React.ChangeEvent<HTMLInputElement>) => {
   //TODO: implement
}


const MyPurchases = () => {
   const [purchaseOrders, setPurchaseOrders] = useState<PurchaseOrder[]>(
      [
         {
            id: "1",
            userId: "",
            date: "2020-05-13 15:23:58",
            price: 100,
            product: [
               {
                  price: 90,
                  state: "IN_PROGRESS",
                  productModel: {
                     description: "",
                     title: "product title",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               },
               {
                  price: 10,
                  state: "PAYED",
                  productModel: {
                     description: "",
                     title: "product title 2",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               }
            ],
            status: "IN_PROGRESS"

         },
         {
            id: "2",
            userId: "",
            date: "2021-12-24 09:30:00",
            price: 250,
            product: [
               {
                  price: 250,
                  state: "DELIVERED",
                  productModel: {
                     description: "",
                     title: "product title 3",
                     id: 1,
                     images: [],
                     specification: [],
                     quantity: 1,
                     price: 100
                  }
               }
            ],
            status: "DELIVERED"

         }
      ]
   )

   return (
      <>
         <h1>My Purchases</h1>

         <Form>
            <Form.Field>
               <Form.Input action>
                  <input type="text" placeholder="Search ..." onChange={(event) => handleSearchInput(event)} />
                  <button className="ui teal right labeled icon button">
                     <i className="search icon"></i>
                     Search
                  </button>
               </Form.Input>
            </Form.Field>
         </Form>

         <hr />

         {
            purchaseOrders?.map((val, idx, arr) => {
               return <PurchaseCard purchaseOrder={val} />
            })
         }
      </>
   )
}

export default MyPurchases;