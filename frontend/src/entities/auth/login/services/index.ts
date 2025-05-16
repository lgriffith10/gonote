import { api } from '@/lib/api'
import type { LoginRequest } from '../types'

export async function loginMutation(request: LoginRequest) {
  return await api('/auth/login', {
    data: request,
  })
}
