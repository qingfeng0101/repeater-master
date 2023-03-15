<template>
  <div>
    <n-button class="mb-8px" type="primary" @click="handleClick">添加</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
  </div>

</template>
  
<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NDataTable, NButton, NTag } from 'naive-ui'
import type { DataTableColumn, PaginationProps } from 'naive-ui'
import { senderSetsList,senderSetstDelete } from "@/api/sender"
import ConfirmButton from '@/components/ConfirmButton.vue'
import { useRouter } from 'vue-router'


interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "sets": string[]

}
const columns: DataTableColumn<DataSource>[] = [
  { title: '名称', key: 'name' },
  {
    title: '消息发送器', key: 'sets', render(row) {
      return !row?.sets || row.sets.map(item => h(
        NTag,
        {
          size: 'small',
          type: 'warning',
          class: 'mr-8px'
        },
        { default: () => item }
      ))
    }
  },

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
            onClick: () => router.push({ name: "SenderSetsEdit", params: { id: row.id } })
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


senderSetsList()
  .then((response) => {
    datas.value = response.data
  })



const router = useRouter()

function handleClick() {
  router.push({ name: 'SenderSetsAdd' })
}
function handleDelete(row: DataSource) {
  senderSetstDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}
</script>
  
  