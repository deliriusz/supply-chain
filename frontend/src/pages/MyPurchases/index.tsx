import { filter } from "lodash";
import React, { useEffect, useState } from "react";
import { Button, Card, Divider, Form, FormProps, Grid, Icon, Segment } from "semantic-ui-react";
import PurchaseCard from "../../components/PurchaseCard";
import AuthContext from "../../hooks/AuthContext";
import PurchaseOrder from "../../interfaces/PurchaseOrder";
import { getPurchaseOrdersForUser } from "../../services/PurchaseOrderService";



const MyPurchases = () => {
   const PURCHASE_LIST_PAGE_SIZE = 10
   const authContext = React.useContext(AuthContext)
   const [purchaseOrders, setPurchaseOrders] = useState<PurchaseOrder[]>([])
   const [filteredPurchaseOrders, setFilteredPurchaseOrders] = useState<PurchaseOrder[]>(purchaseOrders)
   const [hasMoreData, setHasMoreData] = useState(false)
   const [isDataFetching, setIsDataFetching] = useState(false)
   const [fetchedDataOffset, setFetchedDataOffset] = useState(0)

   useEffect(() => {
      if (authContext.auth.isAuthenticated == true) {
         setIsDataFetching(true)
         getPurchaseOrdersForUser(authContext.auth.address!, fetchedDataOffset).then(resp => {
            const concatenatedPurchaseOrders = purchaseOrders.concat(resp.products)
            setPurchaseOrders(concatenatedPurchaseOrders)
            setFilteredPurchaseOrders(concatenatedPurchaseOrders)
            setHasMoreData(fetchedDataOffset + PURCHASE_LIST_PAGE_SIZE < resp.total)
         })
         setIsDataFetching(false)
      }
   }, [fetchedDataOffset])

   const handleFetchData = () => {
      setFetchedDataOffset(fetchedDataOffset + PURCHASE_LIST_PAGE_SIZE)
   }

   //TODO: add fetch from backend
   const handleSearchInput = (event: React.ChangeEvent<HTMLInputElement>) => {
      const filterText = event.target.value.toLowerCase()
      setFilteredPurchaseOrders(purchaseOrders.filter((po, idx, arr) => {
         return po.product.some((prod, idxProd, arrProd) => { return prod.productModel.title.toLowerCase().includes(filterText) })
      })
      )
   }

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
            filteredPurchaseOrders?.map((val, idx, arr) => {
               return <PurchaseCard purchaseOrder={val} />
            })
         }

         {hasMoreData &&
            <Button
               color="teal"
               loading={isDataFetching}
               onClick={() => handleFetchData()}
            >
               Load more
            </Button>
         }
      </>
   )
}

export default MyPurchases;