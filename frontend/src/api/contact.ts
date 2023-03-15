import axios from 'axios'


export function contactList() {
    return axios.get('/api/contact')
}

export function contactGet(id: number) {
    return axios.get(`/api/contact/${id}`)
}

export function contactPost(data: any) {
    return axios.post('/api/contact', data)
}

export function contactPut(id: number, data: any) {
    return axios.put(`/api/contact/${id}`, data)
}

export function contactDelete(id: number,) {
    return axios.delete(`/api/contact/${id}`)
}

export function contactSync() {
    return axios.post('/api/contact/sync')
}