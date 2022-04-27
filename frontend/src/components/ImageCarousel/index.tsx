import React, { useEffect, useState } from "react"
import { Card, Grid, Image } from "semantic-ui-react";
import { Image as ProductImage } from "../../interfaces/Product";
import './style.css'

const IMAGE_PREVIEW_WIDTH = 3;
const IMAGE_PREVIEW_SILDER_WIDTH = 1;
const IMAGE_PREVIEW_COUNT = Math.floor(16 / IMAGE_PREVIEW_WIDTH) - 1;

const ImageCarousel: React.FC<{ images: ProductImage[] | undefined }> = ({ images }) => {

   const getImageForLocation = (imgLocation: string | undefined) => {
      let imageToDisplay: any;
      if (imgLocation) {
         imageToDisplay = <Image
            src={imgLocation} />
      } else {
         imageToDisplay = <div className="ui placeholder"><div className="image"></div></div>
      }

      return imageToDisplay
   }

   const setImageComponentToPreview = (imgLocation: string | undefined) => {
      setImgToDisplay(getImageForLocation(imgLocation))
   }

   const increasePreviewPage = () => {
      const imagesCount = images ? images.length : 0
      const newPage = page + 1
      if (IMAGE_PREVIEW_COUNT * newPage < imagesCount) {
         setPage(newPage)
      }
   }

   const decreasePreviewPage = () => {
      const newPage = page - 1
      if (newPage <= 0) {
         setPage(0)
      } else {
         setPage(newPage)
      }
   }

   let [imgToDisplay, setImgToDisplay] = useState<any>(getImageForLocation(images?.[0].url))
   let [page, setPage] = useState<number>(0)

   useEffect(() => {
      setImageComponentToPreview(images?.[0].url)
   }, [images?.[0].url])

   return (
      <>
         <Card fluid raised>
            {imgToDisplay}
         </Card>

         <Grid centered>
            <Grid.Column className="paginator" verticalAlign="middle" floated="left"
               onClick={decreasePreviewPage} width={IMAGE_PREVIEW_SILDER_WIDTH}>
               <span>⟨</span>
            </Grid.Column>
            {images?.map((val, idx, arr) => {
               const idxToDisplayFrom = IMAGE_PREVIEW_COUNT * page
               if (idx >= idxToDisplayFrom && idx < idxToDisplayFrom + IMAGE_PREVIEW_COUNT) {
                  return (
                     <Grid.Column id={idx} verticalAlign="middle"
                        onClick={() => setImageComponentToPreview(val.url)} width={IMAGE_PREVIEW_WIDTH}>
                        <Image
                           size='small'
                           src={val.url} />
                     </Grid.Column>
                  )
               }
            })}
            <Grid.Column className="paginator" verticalAlign="middle" floated="right"
               onClick={increasePreviewPage} width={IMAGE_PREVIEW_SILDER_WIDTH}>
               <span>⟩</span>
            </Grid.Column>
         </Grid>
      </>
   )
}

export default ImageCarousel