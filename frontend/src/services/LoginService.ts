import ResponseContent from "../interfaces/ResponseContent";

interface LoginRequest {
   signature?: string,
   data?: string,
   address: string,
   nonce?: string | number,
}

const getChallenge = async (address: string): Promise<ResponseContent> => {
   let request: LoginRequest = { address }

   const requestOptions: RequestInit = {
      method: "POST",
      credentials: "include",
      mode: "cors",
      body: JSON.stringify(request),
   }

   var responseContent: ResponseContent = { isOk: false, status: 500 };

   await fetch(`${process.env.REACT_APP_BACKEND_URL}/auth/challenge`, requestOptions)
      .then(async response => {
         const responseData: any = await response.json()

         responseContent = { data: responseData, isOk: response.ok, status: response.status }
         responseContent.data = responseData
         responseContent.isOk = response.ok
         responseContent.status = response.status

         if (!response.ok) {
            return Promise.reject(responseContent)
         }

      }).catch(err => {
         console.log(err)
      })

   return responseContent
}

const login = async (login: LoginRequest): Promise<ResponseContent> => {

   const requestOptions: RequestInit = {
      method: "POST",
      credentials: "include",
      mode: "cors",
      body: JSON.stringify(login),
   }

   var responseContent: ResponseContent = { isOk: false, status: 500 };

   await fetch(`${process.env.REACT_APP_BACKEND_URL}/auth/login`, requestOptions)
      .then(async response => {
         const responseData: any = await response.text()

         responseContent = { data: responseData, isOk: response.ok, status: response.status }
         responseContent.data = responseData
         responseContent.isOk = response.ok
         responseContent.status = response.status

         if (!response.ok) {
            return Promise.reject(responseContent)
         }

      }).catch(err => {
         console.log(err)
      })

   return responseContent
}

export type { LoginRequest }
export { getChallenge, login }
