import { api } from '@/lib/api'
import type { RegisterRequest } from '../types'

export function registerMutation(request: RegisterRequest) {
  return api.post('/auth/register', request)
}
