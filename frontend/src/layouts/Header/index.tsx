import React from "react";
import { Link } from "react-router-dom";
import { Container, Icon, Image } from "semantic-ui-react";
import { getChallenge, login } from "../../services/LoginService";
import web3 from "../../web3";
import GenericErrorModal from "../../components/GenericErrorModal";
import './style.css'
import PageInformationModal from "../../components/PageInformationModal";

const Header = () => {
   const [isError, setIsError] = React.useState(false)
   const [errorDescription, setErrorDescription] = React.useState("")

   const loginWithMetamask = async () => {
      if (typeof web3 !== "undefined") {
         let accounts = await web3.eth.getAccounts()

         let challenge = await getChallenge(accounts[0])
         console.log(challenge)
         if (challenge.isOk) {
            let hashed = web3.eth.accounts.hashMessage(challenge.data.nonce)
            let signature = await web3.eth.sign(hashed, accounts[0])

            let response = await login({
               address: accounts[0],
               signature: signature,
            })
            setIsError(response.isOk)
            setErrorDescription(response.data)

         } else {
            setErrorDescription("Unable to authorize for login. Please try again.")
            setIsError(true)
         }
      } else {
         setErrorDescription("MetaMask browser extension is required in order to use this page")
         setIsError(true)
      }
   }
   return (
      <>
         <GenericErrorModal open={isError} description={errorDescription} header="Login error" onClose={() => setIsError(false)} />

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
   );
};

export default Header;