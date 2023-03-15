import axios from 'axios'


export function apikeyList() {
    return axios.get('/api/apikey')
}

export function apikeyPost(data: any) {
    return axios.post('/api/apikey', data)
}

export function apikeyDelete(id: number) {
    return axios.delete(`/api/apikey/${id}`)
}

export function apikeyRefresh(id: number) {
    return axios.patch(`/api/apikey/${id}/refresh`)
}