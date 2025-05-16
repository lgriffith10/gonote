import { z } from 'zod'

export const RegisterFormSchema = z
  .object({
    email: z.string().min(1, 'Required').email('Invalid email address'),
    password: z.string().min(8, 'Password must be at least 8 characters long'),
    repeatPassword: z.string(),
    firstname: z.string().min(1, 'Required'),
    lastname: z.string().min(1, 'Required'),
  })
  .refine(
    (data) => {
      return data.password === data.repeatPassword
    },
    {
      message: 'Passwords do not match',
      path: ['repeatPassword'],
    },
  )

export type RegisterForm = z.infer<typeof RegisterFormSchema>
