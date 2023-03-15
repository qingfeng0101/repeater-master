<template>
  <div>
    <n-button type="primary" @click="handleClick" class="mb-8px">添加</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
  </div>

</template>
  
<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NDataTable, NButton } from 'naive-ui'
import type { DataTableColumn, PaginationProps } from 'naive-ui'
import { prometheusDatasourceList, prometheusDatasourceDelete } from "@/api/prometheus"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import ConfirmButton from '@/components/ConfirmButton.vue'

interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "base_url": string
  "alerts": string
  "targets": string
  "rules": string
  "comment": string
}
const columns: DataTableColumn<DataSource>[] = [
  { title: '名称', key: 'name' },
  { title: 'BaseUrl', key: 'base_url' },
  { title: 'Alerts', key: 'alerts' },
  { title: 'Targets', key: 'targets' },
  { title: 'Rules', key: 'rules' },
  { title: '备注', key: 'comment' },
  {
    title: '操作', key: 'actions', width: 180,
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            class: 'mr-8px',
            onClick: () => router.push({ name: 'PrometheusDatasourceEdit', params: { id: row.id } })
          },
          { default: () => '编辑' }
        ),
        h(
          ConfirmButton,
          {
            size: 'small',
            type: 'error',
            onClick: () => handleDelete(row)
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


prometheusDatasourceList()
  .then((response) => {
    datas.value = response.data
  })


const router = useRouter()

function handleClick() {
  router.push({ name: 'PrometheusDatasourceAdd' })
}



function handleDelete(row: DataSource) {
  prometheusDatasourceDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}
</script>
  
  