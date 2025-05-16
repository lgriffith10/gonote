import { api } from '@/lib/api'
import type { LoginRequest } from '../types'

export async function loginMutation(request: LoginRequest) {
  return await api('/auth/login', {
    method: 'POST',
    data: JSON.stringify(request),
  })
}
