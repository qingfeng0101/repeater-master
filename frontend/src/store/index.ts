import { defineStore } from 'pinia'
import { Base64 } from 'js-base64'
import { login } from '../api/login'

export * from './modules/theme'
export * from './modules/prometheus'

function payloadDecode(authToken: string) {
    if (!authToken || authToken.constructor !== String) {
        return null
    }
    let payload = authToken.split('.')[1]
    if (!payload) {
        return null
    }
    return JSON.parse(Base64.decode(payload))
}



export const useMainStore = defineStore('main', {
    state: () => {
        return {
            token: "",
            collapsed: false,
        }
    },
    getters: {
        account: state => {
            let user = payloadDecode(state.token)
            return user?.username
        },
        userClaims: state => {
            return payloadDecode(state.token)
        },
        checkAuthToken(state) {
            let ts = new Date().getTime()
            let payload = payloadDecode(state.token)
            if (!payload || !Object.keys(payload).includes('exp')) {
                return false
            }
            return payload.exp * 1000 > ts
        },
    },
    actions: {
        async Login(username: string, password: string) {
            return await login({ username: username.trim(), password })
                .then(response => {
                    const data = response.data
                    this.token = data.access_token
                    return true
                })

        },
    },
    persist: {
        enabled: true // 启用持久插件
    }
})
