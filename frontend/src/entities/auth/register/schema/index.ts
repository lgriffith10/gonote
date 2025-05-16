import { z } from 'zod'

export const RegisterFormSchema = z.object({
  email: z.string().min(1, 'Required').email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters long'),
  firstname: z.string().min(1, 'Required'),
  lastname: z.string().min(1, 'Required'),
})

export type RegisterForm = z.infer<typeof RegisterFormSchema>
