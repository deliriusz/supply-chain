export default interface LoginRequest {
   signature?: string,
   data?: string,
   address: string,
   nonce?: string | number,
}