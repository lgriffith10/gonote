import { upfetch } from '@/lib/up-fetch'
import type { LoginRequest } from '../types'

export async function loginMutation(request: LoginRequest) {
  return await upfetch('/auth/login', {
    method: 'POST',
    body: JSON.stringify(request),
  })
}
