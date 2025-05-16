import { useMutation } from '@tanstack/vue-query'
import { registerMutation } from '../services'
import type { RegisterRequest } from '../types'

export function useRegisterMutation() {
  return useMutation({
    mutationFn: async (payload: RegisterRequest) => registerMutation(payload),
  })
}
