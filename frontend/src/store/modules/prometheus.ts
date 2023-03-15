import { defineStore } from 'pinia'
interface PrometheusState {
    sender_filter: Record<string, string[]>
}

export const usePrometheusStore = defineStore('prometheus', {
    state: (): PrometheusState => ({
        sender_filter: {}
    }),
    persist: {
        enabled: true,
        strategies: [
            {
                storage: localStorage,
            },
        ],
    }
})