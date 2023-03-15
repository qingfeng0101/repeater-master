<template>
  <div>
    <n-input v-model:value="usernameColumn.filterOptionValue as string" placeholder="过滤"
      style="width: 240px;margin:0 8px 8px 0;">
      <template #suffix>
        <n-icon :component="SearchOutline" />
      </template>
    </n-input>
    <n-button type="primary" @click="handleClick" class="mb-8px mr-8px">添加</n-button>
    <n-button type="warning" @click="handleSync" class="mb-8px" :loading="loading">同步</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
  </div>
</template>
  
<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NDataTable, NButton, NInput, NIcon } from 'naive-ui'
import type { PaginationProps, DataTableBaseColumn } from 'naive-ui'
import { contactList, contactSync,contactDelete } from "@/api/contact"
import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import SearchOutline from "@/assets/icons/SearchOutline.vue"
import ConfirmButton from '@/components/ConfirmButton.vue'

interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "username": string
  "email": string
  "mobile": string
  "wecom": string
  "dingtalk": string
  "lark": string
}
const usernameColumn = reactive<DataTableBaseColumn<DataSource>>(
  {
    title: '用户名',
    width: 140,
    filterOptionValue: null,
    key: 'username',
    filter(value, row) {
      return !!~row.username.indexOf(value.toString()) || !!~(row.name || '').indexOf(value.toString())
    }

  }
)

const columns: DataTableBaseColumn<DataSource>[] = [
  { title: '姓名', key: 'name' },
  usernameColumn,
  // { title: '用户名', key: 'username' },
  { title: '邮箱', key: 'email' },
  { title: '手机', key: 'mobile' },
  { title: '企业微信', key: 'wecom' },
  { title: '钉钉名称', key: 'dingtalk' },
  { title: '飞书名称', key: 'lark' },
  {
    title: '操作', key: 'actions', width: 180,
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            class:'mr-8px',
            onClick: () => handleEdit(row)
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


contactList()
  .then((response) => {
    datas.value = response.data
  })

const router = useRouter()
function handleClick() {
  router.push({ name: 'ContactAdd' })
}
function handleEdit(row: DataSource) {
  router.push({ name: "ContactEdit", params: { id: row.id } })
}


function handleDelete(row: DataSource) {
  contactDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}
const loading = ref<boolean>(false)
function handleSync() {
  loading.value = true
  contactSync()
    .then(() => {
      Message?.success('success')
    })
    .finally(() => {
      loading.value = false
    })

}
</script>
  
  