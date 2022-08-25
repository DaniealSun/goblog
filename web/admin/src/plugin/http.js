import Vue from 'vue'
import axios from 'axios'

let Url = 'http://10.116.32.5:8777/api/v1/'

axios.defaults.baseURL = Url

// 请求拦截器，配置axios请求头携带token
axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})

Vue.prototype.$http = axios

export { Url }

