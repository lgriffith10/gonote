import { useMutation } from '@tanstack/vue-query'
import type { LoginRequest } from '../types'
import { loginMutation } from '../services'

export function useLoginMutation() {
  return useMutation({
    mutationFn: async (payload: LoginRequest) => loginMutation(payload),
  })
}
