import axios from 'axios'


export function credentialStmpList() {
    return axios.get('/api/credential/stmp')
}


export function credentialStmpPost(data: any) {
    return axios.post('/api/credential/stmp', data)
}


export function credentialStmpDelete(id: number) {
    return axios.delete(`/api/credential/stmp/${id}`)
}


export function credentialWecomList() {
    return axios.get('/api/credential/wecom')
}


export function credentialWecomPost(data: any) {
    return axios.post('/api/credential/wecom', data)
}


export function credentialWecomDelete(id: number) {
    return axios.delete(`/api/credential/wecom/${id}`)
}

export function credentialDingList() {
    return axios.get('/api/credential/dingtalk')
}


export function credentialDingPost(data: any) {
    return axios.post('/api/credential/dingtalk', data)
}
export function credentialDingDelete(id: number) {
    return axios.delete(`/api/credential/dingtalk/${id}`)
}
export function credentialLarkList() {
    return axios.get('/api/credential/lark')
}


export function credentialLarkPost(data: any) {
    return axios.post('/api/credential/lark', data)
}
export function credentialLarkDelete(id: number) {
    return axios.delete(`/api/credential/lark/${id}`)
}

 