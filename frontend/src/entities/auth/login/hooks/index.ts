import { useMutation } from '@tanstack/vue-query'
import type { LoginRequest, LoginResponse } from '../types'
import { loginMutation } from '../services'
import type { AxiosResponse } from 'axios'

export function useLoginMutation() {
  return useMutation({
    mutationFn: async (payload: LoginRequest) => await loginMutation(payload),
  })
}
