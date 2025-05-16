import { describe, test, vi, expect, beforeAll, afterAll, afterEach } from 'vitest'
import { fireEvent, getByText, render } from '@testing-library/vue'
import LoginFormComponent from '../../form/LoginForm.vue'
import { LoginFormBuilder } from '../builders/LoginFormBuilder'
import { VueQueryPlugin } from '@tanstack/vue-query'
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import { flushPromises } from '@vue/test-utils'
import waitForExpect from 'wait-for-expect'

export const handlers = [
  http.post(`${import.meta.env.VITE_API_URL}/auth/login`, () => {
    return HttpResponse.json({ token: 'mocked_token' }, { status: 200 })
  }),
]

const server = setupServer(...handlers)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('LoginFormComponent', () => {
  test('should register token in local storage', async () => {
    // Arrange

    const form = LoginFormBuilder.create().get()

    const { getByTestId } = render(LoginFormComponent, {
      global: {
        plugins: [VueQueryPlugin, autoAnimatePlugin],
      },
    })

    // Act

    await fireEvent.update(getByTestId('login-form-email'), form.email)
    await fireEvent.update(getByTestId('login-form-password'), form.password)
    await fireEvent.click(getByTestId('login-form-button'))

    // Assert

    await flushPromises()
    await waitForExpect(() => {
      expect(localStorage.getItem('bearer-token')).toBe('mocked_token')
    })
  })
})
