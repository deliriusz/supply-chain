import React from "react";
import { Link } from "react-router-dom";
import { Container, Icon, Image } from "semantic-ui-react";
import './style.css'
import PageInformationModal from "../../components/PageInformationModal";
import AuthContext from "../../hooks/AuthContext";
import { useLogin } from "../../hooks/useLogin";
import { useLogout } from "../../hooks/useLogout";
import { UserRole } from "../../interfaces/UserRole";

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
               <Link className="item" to="/product">Product catalogue</Link>
               <Link className="item" to="/purchase">My Purchases</Link>
               <div className="right menu">
                  {(authContext.auth.role === UserRole.Admin
                     || authContext.auth.role === UserRole.DashboardViewer) &&
                     <Link className="item" to="/admin">
                        <Icon name="th list" />&nbsp;Dashboard
                     </Link>
                  }
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