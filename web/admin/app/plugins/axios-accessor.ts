import { Plugin } from '@nuxt/types'
import { initializeAxios } from '~/plugins/axios'

const accessor: Plugin = ({ $axios, store, redirect }) => {
  initializeAxios($axios)

  $axios.onRequest((config: any) => {
    config.baseURL = process.env.apiURL
    config.withCredentials = true

    const token: string = `Bearer ${store.getters['auth/getToken']}`
    if (token) {
      config.headers.common['Authorization'] = token
    }

    return config
  })

  $axios.onError((err: any) => {
    if (err.response.status === 401) {
      store.dispatch('auth/logout')
      redirect('/signin')
    }
  })
}

export default accessor
