import React from "react";
import { Button, Icon } from "semantic-ui-react";
import './style.css'

const Footer = () => {
   return (
      <>
         <footer className="footer">
            <div className="ui center aligned inverted container">
               <div className="ui stackable inverted grid">
                  <div className="three wide column">
                     <h4 className="ui inverted header">About Us</h4>
                     <div className="ui link list inverted">
                        <a className="item" href="#">Mission</a>
                        <a className="item" href="#">News</a>
                        <a className="item" href="#">Career</a>
                        <a className="item" href="#">Contact</a>
                     </div>
                  </div>
                  <div className="three wide column">
                     <h4 className="ui header inverted">Shopping</h4>
                     <div className="ui link list inverted">
                        <a className="item" href="#">Check processing status</a>
                        <a className="item" href="#">Loans</a>
                        <a className="item" href="#">Gift card</a>
                        <a className="item" href="#">Foreign shipping</a>
                     </div>
                  </div>
                  <div className="seven wide right floated column">
                     <h4 className="ui header inverted">Social Media</h4>
                     <p>You can find us on:</p>
                     <Button
                        icon
                        compact
                        as="a"
                        href="#"
                        size="big"
                        className="facebook"
                     >
                        <Icon name="facebook" />
                     </Button>
                     <Button
                        icon
                        compact
                        as="a"
                        href="#"
                        size="big"
                        className="youtube"
                     >
                        <Icon name="youtube" />
                     </Button>
                     <Button
                        icon
                        compact
                        as="a"
                        href="#"
                        size="big"
                        className="instagram"
                     >
                        <Icon name="instagram" />
                     </Button>
                     <Button
                        icon
                        compact
                        as="a"
                        href="#"
                        size="big"
                        className="twitter"
                     >
                        <Icon name="twitter" />
                     </Button>
                  </div>
               </div>
               <div className="ui section divider"></div>
            </div>
            <p>&copy; All rights reserved, Rafał Kalinowski</p>
         </footer>
      </>
   )
}

export default Footer;