import React from "react";
import { Link } from "react-router-dom";
import { Container, Icon, Image } from "semantic-ui-react";
import { getChallenge, login, logout } from "../../services/LoginService";
import web3 from "../../web3";
import './style.css'
import PageInformationModal from "../../components/PageInformationModal";
import AuthContext from "../../hooks/AuthContext";
import { useLogin } from "../../hooks/useLogin";
import { useLogout } from "../../hooks/useLogout";

const Header = () => {
   const authContext = React.useContext(AuthContext)
   const login = useLogin(authContext)
   const logout = useLogout(authContext)

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
                  {authContext.auth.isAuthenticated &&
                     <a className="item" onClick={logout}>Log out &nbsp;
                        <Icon name="sign-out" />
                     </a>
                  }
                  {!authContext.auth.isAuthenticated &&
                     <a className="item" onClick={login}>Log in &nbsp;
                        <Icon name="sign-in" />
                     </a>
                  }
               </div>
            </Container>
         </div>
      </>
   )
}

export default Header;