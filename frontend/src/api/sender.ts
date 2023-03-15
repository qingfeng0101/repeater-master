import axios from 'axios'

export function senderList() {
    return axios.get('/api/sender')
}

export function senderSetsGet(id: number) {
    return axios.get(`/api/sender/set/${id}`)
}

export function senderSetsList() {
    return axios.get('/api/sender/set')
}

export function senderSetsPost(data: any) {
    return axios.post('/api/sender/set', data)
}
export function senderSetstDelete(id: number,) {
    return axios.delete(`/api/sender/set/${id}`)
}


export function senderSetsPut(id: number, data: any) {
    return axios.put(`/api/sender/set/${id}`, data)
}
