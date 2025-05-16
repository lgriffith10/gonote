import { describe, test, expect } from 'vitest'
import { RegisterFormBuilder } from '../builders/RegisterFormBuilder'
import { RegisterFormSchema, type RegisterForm } from '../../schema'

describe('LoginFormSchema', () => {
  test('should validate correct data', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(true)
  })

  test('should invalidate incorrect email format', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().withEmail('invalid-email').get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Invalid email address')
  })

  test('should invalidate empty email', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().withEmail('').get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Required')
  })

  test('should invalidate empty password', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().withPassword('').get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Password must be at least 8 characters long')
  })

  test('should invalidate empty firstname', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().withFirstname('').get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Required')
  })

  test('should invalidate empty lastname', () => {
    // Arrange

    const form: RegisterForm = RegisterFormBuilder.create().withLastname('').get()

    // Act

    const result = RegisterFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Required')
  })
})
