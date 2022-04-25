import React from "react";
import { logout } from "../services/LoginService";
import web3 from "../web3";
import { AuthDispatchContext } from "../hooks/AuthContext";

const useLogout = (authContext: AuthDispatchContext) => async () => {
   const accounts = await web3?.eth.getAccounts() || [""]

   logout().then(
      response => {
         authContext.dispatcher!(
            {
               isAuthenticated: false,
               address: accounts[0],
               message: response.data,
               isError: !response.isOk,
               action: "LOGOUT"
            })
      }
   )

}

export { useLogout }