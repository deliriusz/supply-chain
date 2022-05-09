import React, { useEffect, useState } from "react";
import { Table, Image, Menu, Icon, Rating } from "semantic-ui-react";
import ProductModel from "../../interfaces/ProductModel";
import * as ProductService from '../../services/ProductService'
import _ from 'lodash';

const getPages = (page: number, maxPage: number, pagesToGenerate = 5): number[] => {
   const maximumRangeAllowed = Math.min(maxPage, pagesToGenerate)
   const pagesRange = Math.floor(maximumRangeAllowed / 2)
   const lastPage = Math.min(page + pagesRange, maxPage)

   if (pagesToGenerate >= maxPage) {
      return _.range(0, maxPage + 1)
   }


   var range = _.range(page - pagesRange, lastPage + 1)

   if (range[0] < 0) {
      range = range.map((val, idx, arr) => {
         return (val + Math.abs(range[0]))
      })
   }
   return range.filter((val, idx, arr) => val <= maxPage)
}

const ProductsTable = () => {
   const PAGE_SIZE = 10
   const [products, setProducts] = useState<ProductModel[]>([]);
   const [currentPage, setCurrentPage] = useState<number>(0)
   const [productsCount, setProductsCount] = useState<number>(0)

   useEffect(() => {
      ProductService.getProducts(currentPage * PAGE_SIZE, PAGE_SIZE).then(products => {
         setProducts(products.products)
         setProductsCount(products.total)
      })
   }, [currentPage])

   return (
      <>
         <Table selectable celled>
            <Table.Header>
               <Table.Row>
                  <Table.HeaderCell>Image</Table.HeaderCell>
                  <Table.HeaderCell>Title</Table.HeaderCell>
                  <Table.HeaderCell>Rating</Table.HeaderCell>
                  <Table.HeaderCell>Price</Table.HeaderCell>
                  <Table.HeaderCell>In Stock</Table.HeaderCell>
                  <Table.HeaderCell>Details</Table.HeaderCell>
               </Table.Row>
            </Table.Header>

            <Table.Body>
               {
                  products?.map((val, idx, arr) => {
                     return (
                        <Table.Row>
                           <Table.Cell><Image size="tiny" src={val.images && val.images[0].url} /></Table.Cell>
                           <Table.Cell selectable><a href={`/product/${val.id}`}>{val.title}</a></Table.Cell>
                           <Table.Cell><Rating disabled icon='star' defaultRating={3 + idx % 3} maxRating={5} /></Table.Cell>
                           <Table.Cell>{val.price}</Table.Cell>
                           <Table.Cell>{val.quantity}</Table.Cell>
                           <Table.Cell selectable><a href={`/product/${val.id}`}>Check Details</a></Table.Cell>
                        </Table.Row>
                     )
                  })
               }
            </Table.Body>

            <Table.Footer>
               <Table.Row>
                  <Table.HeaderCell colSpan='6'>
                     <Menu floated='right' pagination>
                        <Menu.Item disabled={currentPage <= 0} as='a' icon onClick={() => setCurrentPage(currentPage - 1)}>
                           <Icon name='chevron left' />
                        </Menu.Item>
                        {
                           getPages(currentPage, Math.ceil(productsCount / PAGE_SIZE) - 1).map((val, idx, arr) => {
                              return (
                                 <Menu.Item as='a' disabled={currentPage === val} onClick={() => setCurrentPage(val)}>{val + 1}</Menu.Item>
                              )
                           })
                        }

                        <Menu.Item as='a' icon disabled={currentPage >= Math.ceil(productsCount / PAGE_SIZE) - 1} onClick={() => setCurrentPage(currentPage + 1)}>
                           <Icon name='chevron right' />
                        </Menu.Item>
                     </Menu>
                  </Table.HeaderCell>
               </Table.Row>
            </Table.Footer>
         </Table>
      </>
   )
}

export default ProductsTable;