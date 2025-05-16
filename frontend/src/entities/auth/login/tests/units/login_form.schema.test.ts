import { describe, expect, test } from 'vitest'
import { LoginFormBuilder } from '../builders/LoginFormBuilder'
import { LoginFormSchema, type LoginForm } from '../../schema'

describe('LoginFormSchema', () => {
  test('should validate correct data', () => {
    // Arrange

    const form: LoginForm = LoginFormBuilder.create().get()

    // Act

    const result = LoginFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(true)
  })

  test('should invalidate incorrect email format', () => {
    // Arrange

    const form: LoginForm = LoginFormBuilder.create().withEmail('invalid-email').get()

    // Act

    const result = LoginFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Invalid email address')
  })

  test('should invalidate empty email', () => {
    // Arrange

    const form: LoginForm = LoginFormBuilder.create().withEmail('').get()

    // Act

    const result = LoginFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Required')
  })

  test('should invalidate empty password', () => {
    // Arrange

    const form: LoginForm = LoginFormBuilder.create().withPassword('').get()

    // Act

    const result = LoginFormSchema.safeParse(form)

    // Assert

    expect(result.success).toBe(false)
    expect(result.error?.issues[0].message).toBe('Required')
  })
})
