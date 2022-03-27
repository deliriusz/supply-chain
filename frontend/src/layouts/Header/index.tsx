import React from "react";
import { Link } from "react-router-dom";
import { Container, Icon, Image } from "semantic-ui-react";
import { getChallenge, login } from "../../services/LoginService";
import web3 from "../../web3";
import GenericErrorModal from "../../components/GenericErrorModal";
import './style.css'
import PageInformationModal from "../../components/PageInformationModal";
import AuthContext from "../../hooks/AuthContext";

const Header = () => {
   const authContext = React.useContext(AuthContext)


   const loginWithMetamask = async () => {
      if (typeof web3 === "undefined") {
         authContext.dispatcher!(
            {
               isAuthenticated: false,
               address: null,
               message: "MetaMask browser extension is required in order to use this page",
               isError: true,
               action: "LOGIN_ERROR",
            })

         return
      }
      let accounts = await web3.eth.getAccounts()

      let challenge = await getChallenge(accounts[0])
      if (challenge.isOk) {
         let hashed = web3.eth.accounts.hashMessage(challenge.data.nonce)
         let signature = await web3.eth.sign(hashed, accounts[0])

         let response = await login({
            address: accounts[0],
            signature: signature,
         })

         authContext.dispatcher!(
            {
               isAuthenticated: response.isOk,
               address: accounts[0],
               message: response.data,
               isError: !response.isOk,
               action: response.isOk ? "LOGIN" : "LOGIN_ERROR",
            })

      } else {
         authContext.dispatcher!(
            {
               isAuthenticated: false,
               address: null,
               message: "Unable to authorize for login. Please try again.",
               isError: true,
               action: "LOGIN_ERROR",
            })
      }
   }

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
                  <a className="item" onClick={loginWithMetamask}>Log in &nbsp;
                     <Icon name="sign-in" />
                  </a>
               </div>
            </Container>
         </div>
      </>
   )
}

export default Header;