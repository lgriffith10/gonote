import { api } from '@/lib/api'
import type { RegisterRequest } from '../types'

export async function registerMutation(request: RegisterRequest) {
  return await api.post('/auth/register', request).catch((err) => {
    throw err
  })
}
