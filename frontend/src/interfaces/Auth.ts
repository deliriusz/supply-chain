import { UserRole } from "./UserRole";

export default interface Auth {
   isAuthenticated: boolean,
   isError: boolean,
   address: string | null,
   action: "LOGIN" | "LOGIN_ERROR" | "LOGOUT",
   role?: UserRole
   message: string,
}