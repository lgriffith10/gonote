import { up } from 'up-fetch'

export const upfetch = up(fetch, () => ({
  baseUrl: import.meta.env.VITE_API_URL,
  timeout: 5_000,
  headers: {
    Authorization: localStorage.getItem('bearer-token'),
  },
}))
