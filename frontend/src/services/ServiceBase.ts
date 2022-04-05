import ResponseContent from "../interfaces/ResponseContent";

const DEFAULT_REQUEST_OPTIONS: RequestInit = {
   method: "GET",
   mode: "cors",
   credentials: "include",
}

const callService = async <T>(uri: string, requestOptions: RequestInit = DEFAULT_REQUEST_OPTIONS): Promise<ResponseContent<T>> => {
   const options: RequestInit = {
      ...DEFAULT_REQUEST_OPTIONS, ...requestOptions
   }

   var responseContent: ResponseContent<T> = { isOk: false, status: 500 };

   return fetch(`${process.env.REACT_APP_BACKEND_URL}/${uri}`, options)
      .then(response => {

         responseContent = { isOk: response.ok, status: response.status }

         if (!response.ok) {
            throw new Error(response.statusText)
         }

         return response.text()
      }).then(response => {
         try {
            responseContent.data = JSON.parse(response)
         } catch (e) {
            responseContent.data = response as any
         }

         return responseContent
      })
}

export { callService }