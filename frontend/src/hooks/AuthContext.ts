import React from "react";
import Auth from '../interfaces/Auth'


export const initialState: Auth = {
   isAuthenticated: false,
   address: null,
   message: "",
   isError: false,
   action: "LOGOUT",
};

interface AuthDispatchContext {
   auth: Auth,
   dispatcher: React.Dispatch<Auth> | null
}

const AuthContext = React.createContext<AuthDispatchContext>({ auth: initialState, dispatcher: null });

export const reducer = (state: Auth, action: Auth) => {
   switch (action.action) {
      case "LOGIN":
         localStorage.setItem("user", JSON.stringify(action.address))
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