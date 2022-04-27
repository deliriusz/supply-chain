import { UserRole } from "./UserRole";

export default interface Login {
   signature?: string,
   data?: string,
   address: string,
   role?: UserRole,
   expiresAt?: number,
   ttl?: number,
   nonce?: string | number,
}