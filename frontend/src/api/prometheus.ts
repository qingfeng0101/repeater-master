import axios from 'axios'

export function prometheusDatasourceList() {
    return axios.get('/api/prometheus/datasource')
}

export function prometheusDatasourceGet(id: number) {
    return axios.get(`/api/prometheus/datasource/${id}`)
}

export function prometheusDatasourcePost(data: any) {
    return axios.post('/api/prometheus/datasource', data)
}

export function prometheusDatasourcePut(id: number, data: any) {
    return axios.put(`/api/prometheus/datasource/${id}`, data)
}

export function prometheusDatasourceDelete(id: number) {
    return axios.delete(`/api/prometheus/datasource/${id}`)
}

export function prometheusAlertsList() {
    return axios.get('/api/prometheus/alert')
}

export function prometheusLabelSetTarget(params: { source: string }) {
    return axios.get('/api/prometheus/labelset/target', { params })
}

export function prometheusSeverityTarget(params: { source: string }) {
    return axios.get('/api/prometheus/severity/target', { params })
}

export function prometheusDatasourceFakeList() {
    return axios.get('/api/prometheus/datasource/fake')
}

export function prometheusDatasourceFakeGet(id: number) {
    return axios.get(`/api/prometheus/datasource/fake/${id}`)
}

export function prometheusDatasourceFakePost(data: any) {
    return axios.post('/api/prometheus/datasource/fake', data)
}

export function prometheusDatasourceFakePut(id: number, data: any) {
    return axios.put(`/api/prometheus/datasource/fake/${id}`, data)
}

export function prometheusDatasourceFakeDelete(id: number) {
    return axios.delete(`/api/prometheus/datasource/fake/${id}`)
}
