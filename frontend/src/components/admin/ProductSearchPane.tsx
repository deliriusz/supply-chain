import React, { useState } from "react";
import { Button, Form, Header, Icon, Segment } from "semantic-ui-react";
import Product from '../../interfaces/Product'
import './ProductSearchPane.css'

const AddProductPane = () => {
   const [product, setProduct] = useState<Product>({
      description: '',
      title: '',
      price: 0,
      quantity: 0,
      id: undefined,
      imgUrl: [],
      specification: [{ name: '', value: '' }],
   })

   const setFormParameter = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>, value: string) => {
      setProduct({ ...product, [e.target.id]: value })
   }

   const setSpecFormParameter = (name: string, idx: number, value: string) => {
      const newSpec = product.specification
      const newSpecItem = { ...newSpec[idx], [name]: value }
      newSpec[idx] = newSpecItem
      setProduct({ ...product, specification: newSpec })
      console.log("product after:")
      console.log(product)
   }

   const addSpecFormParameter = () => {
      setProduct({ ...product, specification: [...product.specification, { name: '', value: '' }] })
   }

   const removeSpecFormParameter = (idx: number) => {
      //TODO: fix edge cases, e.g. removing item from the middle
      var specArray: any

      if (product.specification.length >= 1) {
         specArray = []
      } else {
         specArray = product.specification.splice(idx, 1)
      }

      setProduct({ ...product, specification: specArray })
      console.log("product after:")
      console.log(product)
   }

   return (
      <>
         <Form>
            <Form.Field>
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  id="title" label='Title' placeholder='Title' />
            </Form.Field>
            <Form.Group>
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Price' placeholder='Price' id="price" width={3} />
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Available Quantity' id="quantity" placeholder='Available Quantity' width={3} />
            </Form.Group>
            <hr />
            <h5>Specification</h5>
            {
               product.specification.map((val, idx, arr) => {
                  return (
                     <Form.Group inline>
                        <Form.Input onChange={(_, data) => setSpecFormParameter("name", idx, data.value)}
                           id={`spec-name-${idx}`} label='Name' placeholder='Name' width={4} />
                        <Form.Input onChange={(_, data) => setSpecFormParameter("value", idx, data.value)}
                           id={`spec-value-${idx}`} label='Value' placeholder='Value' width={6} />
                        <Button as='a' className="specification-button" onClick={addSpecFormParameter}>+</Button>
                        <Button as='a' className="specification-button" onClick={() => removeSpecFormParameter(idx)}>-</Button>
                     </Form.Group>
                  )
               })
            }

            <Form.Field>
               <Form.TextArea onChange={(event, data) => setFormParameter(event, data.value ? data.value.toString() : '')}
                  fluid='' id="description" label='Description' placeholder='Description' />
            </Form.Field>
            <hr />
            <h5>Images</h5>
         </Form>
      </>
   )

}

export default AddProductPane