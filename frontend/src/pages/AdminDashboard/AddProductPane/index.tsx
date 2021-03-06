import React, { useState } from "react";
import { Button, Form, FormProps, GridRow, Header, Icon, Image, Label, Message } from "semantic-ui-react";
import ProductModel from '../../../interfaces/ProductModel'
import ResponseContent from "../../../interfaces/ResponseContent";
import { createImage, createProductModel } from '../../../services/ProductService'
import './style.css'

const EMPTY_SPEC = { name: '', value: '' }

//TODO: add images when creating product, add loading when from is submitted
const AddProductPane = () => {
   const [product, setProduct] = useState<ProductModel>({
      description: '',
      title: '',
      price: 0,
      quantity: 0,
      id: undefined,
      images: [],
      specification: [{ name: '', value: '' }],
   })
   const [images, setImages] = useState<File[]>([])
   const [imageUrls, setImageUrls] = useState<string[]>([])
   const [formSubmitError, setFormSubmitError] = useState<boolean>(false)
   const [formSubmitSuccess, setFormSubmitSuccess] = useState<boolean>(false)
   const [formSubmitErrorMessage, setFormSubmitErrorMessage] = useState<string>("")
   const [createdProductId, setCreatedProductId] = useState<string>("")

   const readImage = (file: File) => {
      const reader = new FileReader();

      reader.addEventListener("load", async function () {
         // convert image file to base64 string
         const result = reader.result?.toString()
         console.log(result)
         console.log(imageUrls)
         await new Promise(r => setTimeout(r, 2000));
         if (result)
            setImageUrls([...imageUrls, result])
      }, false);

      reader.readAsDataURL(file)
   }

   const setImage = (e: React.ChangeEvent<HTMLInputElement>) => {

      const imagesLength = e.target.files?.length || 0
      if (imagesLength > 0) {

         const incomingImages: File[] = Array.of()
         for (let i = 0; i < imagesLength; i++) {
            const file = e.target.files?.item(i)
            if (file) {
               readImage(file)
               incomingImages.push(file)
            }
         }
         setImages([...images, ...incomingImages])
      }
   }

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

      if (product.specification.length <= 1) {
         specArray = [EMPTY_SPEC]
      } else {
         specArray.splice(idx, 1)
      }

      setProduct({ ...product, specification: specArray })
   }

   const submitForm = async (event: React.FormEvent<HTMLFormElement>, data: FormProps) => {
      event.preventDefault()

      const normalizedProduct: ProductModel = (JSON.parse(JSON.stringify(product)))
      if (typeof normalizedProduct.price === "string") {
         normalizedProduct.price = parseInt(normalizedProduct.price)
      }
      if (typeof normalizedProduct.quantity === "string") {
         normalizedProduct.quantity = parseInt(normalizedProduct.quantity)
      }

      const response: ResponseContent<any> = await createProductModel(normalizedProduct)

      if (!response.isOk) {
         setFormSubmitErrorMessage(JSON.stringify(response.data))
         setFormSubmitError(true)
         setFormSubmitSuccess(false)
      } else {
         setCreatedProductId(response.data.id)

         await Promise.all(
            images.map((val, idx, arr) => { return createImage(response.data.id, val) })
         )

         setFormSubmitError(false)
         setFormSubmitSuccess(true)
      }
   }

   const removeImage = (id: number) => {
      setImages(images.filter((_, idx) => idx !== id))
      setImageUrls(imageUrls.filter((_, idx) => idx !== id))
   }

   return (
      <>
         <Form error={formSubmitError} success={formSubmitSuccess}
            onSubmit={(event: React.FormEvent<HTMLFormElement>, data: FormProps) => submitForm(event, data)}>
            <Form.Field>
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  id="title" label='Title' placeholder='Title' />
            </Form.Field>
            <Form.Group className="field-group">
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Price' placeholder='Price' id="price" width={3} />
               <Form.Input onChange={(event, data) => setFormParameter(event, data.value)}
                  label='Quantity' id="quantity" placeholder='Available Quantity' width={3} />
            </Form.Group>
            <hr />
            <h5>Specification</h5>
            {
               product.specification.map((val, idx, _) => {
                  return (
                     <Form.Group className="field-group" inline>
                        <Form.Input onChange={(_, data) => setSpecFormParameter("name", idx, data.value)}
                           id={`spec-name-${idx}`} value={val.name} label='Name' placeholder='Name' width={4} />
                        <Form.Input onChange={(_, data) => setSpecFormParameter("value", idx, data.value)}
                           id={`spec-value-${idx}`} label='Value' value={val.value} placeholder='Value' width={6} />
                        <Button as='a' className="specification-button" onClick={addSpecFormParameter}>+</Button>
                        <Button as='a' className="specification-button" onClick={() => removeSpecFormParameter(idx)}>-</Button>
                     </Form.Group>
                  )
               })
            }

            <Form.Field>
               <Form.TextArea onChange={(event, data) => setFormParameter(event, data.value ? data.value.toString() : '')}
                  id="description" label='Description' placeholder='Description' />
            </Form.Field>
            <hr />
            <h5>Images</h5>
            <Form.Group>
               <label className="custom-file-upload">
                  <input type="file" id="upload-file" onChange={(e) => setImage(e)}
                     accept="image/png, image/jpg" multiple></input>
                  <GridRow>
                     <Icon size="massive" name="upload" />
                  </GridRow>
                  <GridRow>
                     <Header size="huge">Click to upload</Header>
                  </GridRow>
                  <GridRow>
                     <Header size="medium">... Or drag and drop here</Header>
                  </GridRow>
               </label>
            </Form.Group>

            <Image.Group className="image-group" size="medium">
               {
                  imageUrls.map((val, idx, _) => {
                     return (
                        <div className="image-container">
                           <Label className="remove-image" onClick={(e) => removeImage(idx)} corner="right">
                              <Icon name="close" />
                           </Label>
                           <Image
                              rounded
                              className="image-content"
                              size="medium"
                              src={val}
                           />
                        </div>
                     )
                  })
               }
            </Image.Group>
            <Message
               success
               header='Product successfully created'
               content={<p>You can check it by visiting <a href={`/product/${createdProductId}`}>THIS LINK</a></p>}
            />
            <Message
               error
               header='Error while sending a form'
               content={formSubmitErrorMessage}
            />
            <Form.Button>Submit</Form.Button>

         </Form>
      </>
   )

}

export default AddProductPane