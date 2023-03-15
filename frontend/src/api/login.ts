import axios from 'axios'


export function login(userInfo: { username: string, password: string }) {
    return axios.post('/api/user/login', userInfo, { headers: { 'Content-Type': 'multipart/form-data' } })
}