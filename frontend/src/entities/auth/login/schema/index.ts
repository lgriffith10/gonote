import { z } from 'zod'

export const LoginFormSchema = z.object({
  email: z.string().min(1, 'Required').email('Invalid email address'),
  password: z.string().min(1, 'Required'),
})

export type LoginForm = z.infer<typeof LoginFormSchema>
