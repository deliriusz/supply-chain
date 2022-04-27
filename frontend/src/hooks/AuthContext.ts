import React from "react";
import Auth from '../interfaces/Auth'
import { UserRole } from "../interfaces/UserRole";


const initialState: Auth = {
   isAuthenticated: false,
   address: null,
   message: "",
   isError: false,
   role: UserRole.Unauthorized,
   action: "LOGOUT",
};

const loadState = (): Auth => {
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

const reducer = (state: Auth, action: Auth) => {
   switch (action.action) {
      case "LOGIN":
         localStorage.setItem("user", JSON.stringify(action))
         return action
      case "LOGIN_ERROR":
         localStorage.removeItem("user")
         return action
      case "LOGOUT":
         localStorage.removeItem("user")
         return action
      default:
         return state
   }
};

export default AuthContext
export { reducer, initialState, loadState }

export type { AuthDispatchContext }