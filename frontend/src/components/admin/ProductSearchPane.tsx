import React, { useState } from "react";
import { Button, Form, FormProps } from "semantic-ui-react";
import Product from '../../interfaces/Product'
import './ProductSearchPane.css'

const EMPTY_SPEC = { name: '', value: '' }

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
   }

   const addSpecFormParameter = () => {
      setProduct({ ...product, specification: [...product.specification, EMPTY_SPEC] })
   }

   const removeSpecFormParameter = (idx: number) => {
      var specArray = product.specification
      console.log(product.specification)

      if (product.specification.length <= 1) {
         specArray = [EMPTY_SPEC]
      } else {
         specArray.splice(idx, 1)
      }

      setProduct({ ...product, specification: specArray })
   }

   const submitForm = (event: React.FormEvent<HTMLFormElement>, data: FormProps) => {
      event.preventDefault()
   }

   return (
      <>
         <Form onSubmit={(event: React.FormEvent<HTMLFormElement>, data: FormProps) => submitForm(event, data)}>
            <Form.Field>
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  id="title" label='Title' placeholder='Title' />
            </Form.Field>
            <Form.Group inline>
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Price' placeholder='Price' id="price" width={3} />
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Quantity' id="quantity" placeholder='Available Quantity' width={3} />
            </Form.Group>
            <hr />
            <h5>Specification</h5>
            {
               product.specification.map((val, idx, arr) => {
                  return (
                     <Form.Group inline>
                        <Form.Input onChange={(_, data) => setSpecFormParameter("name", idx, data.value)}
                           id={`spec-name-${idx}`} value={product.specification[idx].name} label='Name' placeholder='Name' width={4} />
                        <Form.Input onChange={(_, data) => setSpecFormParameter("value", idx, data.value)}
                           id={`spec-value-${idx}`} label='Value' value={product.specification[idx].value} placeholder='Value' width={6} />
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

            <Form.Button>Submit</Form.Button>

         </Form>
      </>
   )

}

export default AddProductPane