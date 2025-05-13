import { ResponseError, up } from 'up-fetch'

export const upfetch = up(fetch, () => ({
  baseUrl: 'http://localhost:4000',
  timeout: 5_000,
  headers: { Authorization: localStorage.getItem('bearer-token') },
  onRequest: (options) => {},
  onError: (err: unknown, options) => {
    const status = (err as ResponseError).status

    switch (status) {
      case 400:
        break
    }
  },
}))
