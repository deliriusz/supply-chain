import React from "react";
import { getChallenge, login } from "../services/LoginService";
import web3 from "../web3";
import { AuthDispatchContext } from "../hooks/AuthContext";
import { UserRole } from "../interfaces/UserRole";

const useLogin = (authContext: AuthDispatchContext) => async () => {
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
   if (challenge.isOk && challenge.data && challenge.data.nonce) {
      let hashed = web3.eth.accounts.hashMessage(challenge.data.nonce + '')
      let signature = await web3.eth.sign(hashed, accounts[0])

      let response = await login({
         address: accounts[0],
         signature: signature,
      })

      authContext.dispatcher!(
         {
            isAuthenticated: response.isOk,
            address: accounts[0],
            message: response.isOk ? '' : JSON.stringify(response.data),
            isError: !response.isOk,
            role: response.data?.role || UserRole.Unauthorized,
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

export { useLogin }