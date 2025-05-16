import axios from 'axios'

export const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 5_000,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
    Authorization: `Bearer ${localStorage.getItem('bearer-token')}`,
  },
})

const publicRoutes: string[] = ['/login', '/register']

api.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    if (error.status === 401) {
      localStorage.removeItem('bearer-token')

      if (!publicRoutes.includes(window.location.pathname)) {
        window.location.href = '/login'
      }
    }
  },
)
