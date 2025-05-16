export type RegisterRequest = {
  email: string
  password: string
  firstname: string
  lastname: string
}

export type RegisterResponse = {
  token: string
}
