<template>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id"
        :loading="loading" />
</template>
<script lang="ts" setup>
import { reactive, ref, h } from "vue"
import { prometheusAlertsList } from '@/api/prometheus'
import { NDataTable, NTag, NTime } from 'naive-ui'
import type { PaginationProps, DataTableColumn } from 'naive-ui'
const pagination = reactive<PaginationProps>({
    page: 1,
    pageSize: 15,
    showSizePicker: true,
    pageSizes: [10, 15, 25, 30, 50, 100],
    onChange: (page) => {
        pagination.page = page
    },
    onPageSizeChange: (pageSize) => {
        pagination.pageSize = pageSize
        pagination.page = 1
    }
})

interface DataSource {
    "startsAt": string
    "endsAt": string
    "labels": Record<string, string>
    "annotations": {
        "description": string
        "summary": string
    }
    "status": {
        "state": string
    }
    "fingerprint": string
}
const columns: DataTableColumn<DataSource>[] = [
    { title: '环境', key: 'labels.environment', width: '100', ellipsis: { tooltip: true } },
    { title: '名称', key: 'labels.alertname', width: '180', ellipsis: { tooltip: true } },
    { title: '服务名', key: 'labels.service_name', width: '240', ellipsis: { tooltip: true } },
    {
        title: '级别', key: 'labels.severity', width: '80', render(row) {
            const clorMap: Record<string, string> = { "high": "#f0a020", "critical": "#d03050", "warning": "#af52de", "info": "#909399" }
            let severity = row.labels?.severity as string
            return h(NTag, { color: { color: clorMap[severity] }, }, { default: () => row.labels?.severity })
        }
    },

    { title: '状态', key: 'status.state', width: '100', ellipsis: { tooltip: true } },
    { title: '详情', key: 'annotations.description', ellipsis: { tooltip: true } },
    {
        title: '开始时间', key: 'startsAt', width: '170', render(row) {
            return h(NTime, { time: Date.parse(row.startsAt) })
        }
    },
    {
        title: '结束时间', key: 'endsAt', width: '170', render(row) {
            if (!row.endsAt) {
                return "——"
            }
            return h(NTime, { time: Date.parse(row.endsAt) })
        }
    },
]
const datas = ref<DataSource[]>([])
const loading = ref(true)
prometheusAlertsList()
    .then(response => {
        datas.value = response.data
    })
    .finally(() => loading.value = false)
</script>