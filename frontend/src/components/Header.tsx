import React from "react";
import { Link } from "react-router-dom";
import { Container, Icon, Image } from "semantic-ui-react";
import './Header.css'
import PageInformationModal from "./PageInformationModal";

const Header = () => {
   return (
      <>

         <div className="ui fixed inverted menu">
            <Container>
               <Link className="header item" to="/">
                  <Image size="mini" src={`${process.env.PUBLIC_URL}/logo-new.png`} />&nbsp;
                  Firmex
               </Link>
               <Link className="item" to="/product">Products</Link>
               <Link className="item" to="/trace">Trace</Link>
               <div className="right menu">
                  <PageInformationModal />
                  <Link className="item" to="/admin">Log in &nbsp;
                     <Icon name="sign-in" />
                  </Link>
               </div>
            </Container>
         </div>
      </>
   );
};

export default Header;