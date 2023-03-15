<template>
  <div>
    <n-button type="primary" @click="handleClick" class="mb-8px">添加</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
  </div>

</template>

<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NDataTable, NButton, NTooltip, NTag } from 'naive-ui'
import type { DataTableColumn, PaginationProps } from 'naive-ui'
import { prometheusDatasourceFakeList, prometheusDatasourceFakeDelete } from "@/api/prometheus"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import ConfirmButton from '@/components/ConfirmButton.vue'

interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "label_set_list": Array<any>
  "severity_set": string[]
  "comment": string
}
const columns: DataTableColumn<DataSource>[] = [
  { title: '名称', key: 'name' },
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
  {
    title: '严重程度列表', key: 'severity_set', render(row) {
      if (!row.severity_set) return ''
      return row.severity_set.map(item => h(NTag,
        { class: "mb-4px mr-8px", size: 'small' },
        { default: () => item })
      )
    }
  },
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
            onClick: () => router.push({ name: 'FakePrometheusDatasourceEdit', params: { id: row.id } })
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


prometheusDatasourceFakeList()
  .then((response) => {
    datas.value = response.data
  })


const router = useRouter()

function handleClick() {
  router.push({ name: 'FakePrometheusDatasourceAdd' })
}



function handleDelete(row: DataSource) {
  prometheusDatasourceFakeDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}
</script>

