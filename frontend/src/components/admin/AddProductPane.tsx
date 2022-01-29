import React, { useState } from "react";
import { Button, Form, FormProps, Icon, Image, Label } from "semantic-ui-react";
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
   const [images, setImages] = useState<File[]>([])
   const [imageUrls, setImageUrls] = useState<string[]>([])

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

   const submitForm = (event: React.FormEvent<HTMLFormElement>, data: FormProps) => {
      event.preventDefault()
   }

   const removeImage = (id: number) => {
      setImages(images.filter((_, idx) => idx !== id))
      setImageUrls(imageUrls.filter((_, idx) => idx !== id))
   }

   return (
      <>
         <Form onSubmit={(event: React.FormEvent<HTMLFormElement>, data: FormProps) => submitForm(event, data)}>
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
                  fluid='' id="description" label='Description' placeholder='Description' />
            </Form.Field>
            <hr />
            <h5>Images</h5>
            <Form.Group>
               <input type="file" id="upload-file" onChange={(e) => setImage(e)}
                  accept="image/png, image/jpg" multiple></input>
               <label htmlFor="file">
                  <Button as="span" onClick={e => e.stopPropagation()}>
                     Upload
                  </Button>
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
                              // bordered
                              className="image-content"
                              size="medium"
                              src={val}
                           />
                        </div>
                     )
                  })
               }
            </Image.Group>
            <Form.Button>Submit</Form.Button>

         </Form>
      </>
   )

}

export default AddProductPane