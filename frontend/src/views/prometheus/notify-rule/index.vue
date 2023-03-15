<template>
  <div>
    <n-button type="primary" @click="handleClick" class="mb-8px">添加</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
  </div>

</template>
  
<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NDataTable, NButton, NTooltip, NTag, NBadge } from 'naive-ui'
import type { DataTableColumn, PaginationProps } from 'naive-ui'
import { notifyRulesList, notifyRulesCacheList, notifyRulesDelete, notifyRulesHardDelete } from "@/api/notify"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import ConfirmButton from '@/components/ConfirmButton.vue'

const router = useRouter()
interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "deleted_at": number
  "name": string
  "label_set_list": Array<any>
  "contacts": Array<any>
  "sender_set": any
  "comment": string
}
const columns: DataTableColumn<DataSource>[] = [
  { title: '名称', key: 'name', minWidth: '150' },
  {
    title: '缓存状态', key: 'cache_staus', align: 'center', minWidth: '100', render(row) {
      return h(NBadge, { dot: true, type: !cache.value ? 'info' : cache.value[row.name] ? 'success' : 'error' })
    }
  },
  {
    title: '状态', key: 'status', minWidth: '100', render(row) {
      return h(NTag, { type: row.deleted_at === 0 ? 'success' : 'warning', size: 'small', round: true }, { default: () => row.deleted_at === 0 ? '活动' : '禁用' })
    }
  },
  {
    title: '标签集合列表', key: 'label_set_list', render(row) {

      return row.label_set_list.map(item => h(NTooltip,
        { delay: 500, trigger: 'hover' },
        {
          trigger: () => h(NTag, { class: "mb-4px mr-8px", size: 'small' }, { default: () => item.service_name || item.namespace || item.instance }),
          default: () => JSON.stringify(item)
        })
      )
    }
  },
  { title: '通知发送器集', key: 'sender_set.name', minWidth: '150' },
  {
    title: '通知目标', key: 'contacts', minWidth: '250', render(row) {
      return row.contacts.map(item => h(NTag, { type: 'primary', size: 'small', class: 'mr-8px' }, { default: () => item.name }))
    }
  },
  { title: '备注', key: 'comment', minWidth: '300' },
  {
    title: '操作', key: 'actions', minWidth: '180',
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            class: 'mr-4px',
            onClick: () => router.push({ name: "PrometheusNotifyRuleEdit", params: { id: row.id } })
          },
          { default: () => '编辑' }
        ),
        h(
          ConfirmButton,
          {
            size: 'small',
            type: 'warning',
            class: 'mr-4px',
            onClick: () => handleDelete(row)
          },
          { default: () => '禁用' }
        ),
        h(
          ConfirmButton,
          {
            size: 'small',
            type: 'error',
            onClick: () => handleHardDelete(row)
          },
          { default: () => '删除' }
        )
      ]
    }
  },
]
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
const datas = ref<DataSource[]>([])


notifyRulesList()
  .then((response) => {
    datas.value = response.data
  })
const cache = ref<any>()
notifyRulesCacheList()
  .then((response) => {
    cache.value = response.data
  })

function handleClick() {
  router.push({ name: 'PrometheusNotifyRuleAdd' })
}



function handleDelete(row: DataSource) {
  notifyRulesDelete(row.id)
    .then(() => {
      notifyRulesList()
        .then((response) => {
          datas.value = response.data
        })
      Message?.success('success')
    })
}

function handleHardDelete(row: DataSource) {
  notifyRulesHardDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })
}


</script>
  
  