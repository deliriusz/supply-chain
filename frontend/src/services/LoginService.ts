import LoginRequest from "../interfaces/LoginRequest";
import ResponseContent from "../interfaces/ResponseContent";
import { callService } from "./ServiceBase";


const getChallenge = async (address: string): Promise<ResponseContent<LoginRequest>> => {
   let request: LoginRequest = { address }

   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(request),
   }

   return callService<LoginRequest>('auth/challenge', requestOptions)
}

const login = async (login: LoginRequest): Promise<ResponseContent<any>> => {
   const requestOptions: RequestInit = {
      method: "POST",
      body: JSON.stringify(login),
   }

   return callService<any>('auth/login', requestOptions)
}

const logout = async (): Promise<ResponseContent<any>> => {
   return callService<any>('auth/logout')
}

export { getChallenge, login, logout }