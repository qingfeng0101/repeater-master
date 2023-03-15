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
import { credentialStmpList, credentialStmpDelete } from "@/api/stmp"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import ConfirmButton from '@/components/ConfirmButton.vue'

interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "address": string
  "username": string
  "password": string
  "auth_type": string
}
const columns: DataTableColumn<DataSource>[] = [
  { title: '名称', key: 'name' },
  { title: '地址', key: 'address' },
  { title: '用户名', key: 'username' },
  { title: '密码', key: 'password' },
  { title: '认证类型', key: 'auth_type' },
  {
    title: '操作', key: 'actions', width: 180,
    render(row) {
      return [
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


credentialStmpList()
  .then((response) => {
    datas.value = response.data
  })


const router = useRouter()

function handleClick() {
  router.push({ name: 'CredentialStmpAdd' })
}



function handleDelete(row: DataSource) {
  credentialStmpDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}
</script>
  
  