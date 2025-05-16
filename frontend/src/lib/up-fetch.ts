import { up } from 'up-fetch'

export const upfetch = up(fetch, () => ({
  baseUrl: import.meta.env.VITE_API_URL,
  timeout: 5_000,
  headers: {
    'Content-Type': 'application/json',
    Authorization: localStorage.getItem('bearer-token'),
  },
}))
