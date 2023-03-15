import axios from 'axios'

export function notifyRulesCacheList() {
    return axios.get('/api/notify/rule/cache')
}

export function notifyRulesCacheTree() {
    return axios.get('/api/notify/rule/index/cache')
}

export function notifyRulesList() {
    return axios.get('/api/notify/rule')
}

export function notifyRulesGet(id: number) {
    return axios.get(`/api/notify/rule/${id}`)
}

export function notifyRulesPost(data: any) {
    return axios.post('/api/notify/rule', data)
}

export function notifyRulesPut(id: number, data: any) {
    return axios.put(`/api/notify/rule/${id}`, data)
}

export function notifyRulesDelete(id: number) {
    return axios.delete(`/api/notify/rule/${id}`)
}

export function notifyRulesHardDelete(id: number) {
    return axios.delete(`/api/notify/rule/${id}/unscoped`)
}
