import React from 'react';
import { Grid, Image } from 'semantic-ui-react';
import ProductListing from '../../components/ProductListing';

const Main = () => {
   return (
      <>
         {/* //TODO: implement real banner */}

         <Grid>
            <Grid.Row>
               <Grid.Column>
                  <Image src={`${process.env.PUBLIC_URL}/main-banner.png`} />
               </Grid.Column>
            </Grid.Row>
            <Grid.Row>
               <Grid.Column width={5}>
                  <Image size='medium' src={`${process.env.PUBLIC_URL}/hot-shot.png`} />
               </Grid.Column>
               <Grid.Column width={11}>
                  <hr />
                  <h2>Hot</h2>
                  <ProductListing />
               </Grid.Column>
            </Grid.Row>
            <Grid.Row>
               <Grid.Column>
                  <hr />
                  <h2>Promotions</h2>
                  <ProductListing />
               </Grid.Column>
            </Grid.Row>
            <Grid.Row>
               <Grid.Column>
                  <hr />
                  <Image src={`${process.env.PUBLIC_URL}/newsletter.png`} />
               </Grid.Column>
            </Grid.Row>
         </Grid>
      </>
   )
}

export default Main;