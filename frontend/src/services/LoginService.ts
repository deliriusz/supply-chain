import Login from "../interfaces/LoginRequest";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";


const getChallenge = async (address: string): Promise<ResponseContent<Login>> => {
   let request: Login = { address }

   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(request),
   }

   return callService<Login>('auth/challenge', requestOptions)
}

const login = async (login: Login): Promise<ResponseContent<Login>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(login),
   }

   return callService<Login>('auth/login', requestOptions)
}

const logout = async (): Promise<ResponseContent<any>> => {
   return callService<any>('auth/logout')
}

export { getChallenge, login, logout }