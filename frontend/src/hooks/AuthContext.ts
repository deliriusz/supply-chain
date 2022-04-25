import React from "react";
import Auth from '../interfaces/Auth'


export const initialState: Auth = {
   isAuthenticated: false,
   address: null,
   message: "",
   isError: false,
   action: "LOGOUT",
};

export const loadState = (): Auth => {
   let userString = localStorage.getItem("user")
   if (userString != null) {
      return JSON.parse(userString)
   }

   return initialState
}

interface AuthDispatchContext {
   auth: Auth,
   dispatcher: React.Dispatch<Auth> | null
}

const AuthContext = React.createContext<AuthDispatchContext>({ auth: initialState, dispatcher: null });

export const reducer = (state: Auth, action: Auth) => {
   console.log(action)
   switch (action.action) {
      case "LOGIN":
         localStorage.setItem("user", JSON.stringify(action))
         return action
      case "LOGIN_ERROR":
         localStorage.clear()
         return action
      case "LOGOUT":
         localStorage.clear()
         return action
      default:
         return state
   }
};

export default AuthContext

export type { AuthDispatchContext }