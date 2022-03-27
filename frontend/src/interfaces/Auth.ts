export default interface Auth {
   isAuthenticated: boolean,
   isError: boolean,
   address: string | null,
   action: "LOGIN" | "LOGIN_ERROR" | "LOGOUT",
   message: string,
}