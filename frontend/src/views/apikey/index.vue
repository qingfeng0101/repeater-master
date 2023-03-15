<template>
  <div>
    <n-input v-model:value="nameColumn.filterOptionValue as string" placeholder="过滤"
      style="width: 240px;margin:0 8px 8px 0;">
      <template #suffix>
        <n-icon :component="SearchOutline" />
      </template>
    </n-input>
    <n-button type="primary" @click="showAPIKeyModal = true" class="mb-8px mr-8px">添加</n-button>
    <n-data-table :data="datas" :pagination="pagination" :columns="columns" :row-key="row => row.id" />
    <n-modal v-model:show="showAPIKeyModal" preset="card" :mask-closable="false" title="创建API Key" style="width:680px">
      <n-form ref="formRef" :model="model" label-placement="left" :label-width="120" :rules="rules">
        <n-form-item label="名称:" path="name">
          <n-input v-model:value="model.name" />
        </n-form-item>
        <n-form-item label="有效期:" path="expire_delta">
          <n-input-group>
            <n-input-number class="w-66%" v-model:value="expire_delta" :show-button="false" />
            <n-input-group-label>天</n-input-group-label>
          </n-input-group>
        </n-form-item>
      </n-form>
      <div class="flex justify-end">
        <n-button round type="primary" @click="handleSubmitClick">
          提交
        </n-button>
      </div>
    </n-modal>
  </div>
</template>
  
<script setup lang="ts">
import { h, ref, reactive, computed } from 'vue'
import { NDataTable, NButton, NInput, NIcon, NTime, NModal, NForm, NFormItem, NInputNumber, NInputGroup, NInputGroupLabel } from 'naive-ui'
import type { PaginationProps, DataTableBaseColumn, FormInst, FormRules } from 'naive-ui'
import { apikeyList, apikeyRefresh, apikeyDelete, apikeyPost } from '@/api/apikey'

import { Message } from '@/utils'
import SearchOutline from '@/assets/icons/SearchOutline.vue'
import ConfirmButton from '@/components/ConfirmButton.vue'

interface DataSource {
  "id": number
  "created_at": number
  "updated_at": number
  "name": string
  "expire_delta": number
  "key": string
}
const showAPIKeyModal = ref(false)

const model = reactive<Pick<DataSource, 'name' | 'expire_delta'>>({
  name: '',
  expire_delta: 1 * 24 * 60 * 60
})
const nameColumn = reactive<DataTableBaseColumn<DataSource>>(
  {
    title: '名称',
    width: 140,
    filterOptionValue: null,
    key: 'name',
    filter(value, row) {
      return !!~(row.name || '').indexOf(value.toString())
    }

  }
)

const expire_delta = computed({
  get(): number {
    return model.expire_delta / 24 / 60 / 60
  },
  set(value: number) {
    model.expire_delta = value * 24 * 60 * 60
  }
})
const columns: DataTableBaseColumn<DataSource>[] = [

  nameColumn,
  {
    title: '创建时间', key: 'created_at', sorter: 'default',
    render(row) {
      return h(
        NTime,
        {
          time: row.created_at,
          unix: true
        }
      )
    }
  },
  {
    title: '失效时间', key: 'updated_at',
    render(row) {
      return h(
        NTime,
        {
          time: row.updated_at + row.expire_delta,
          unix: true
        }
      )
    }
  },

  { title: 'Key', key: 'key' },

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
            onClick: () => handleRefresh(row)
          },
          { default: () => '展期' }
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
const loading = ref<boolean>(false)
function getData() {
  loading.value = true
  apikeyList()
    .then((response) => {
      datas.value = response.data
    })
    .finally(() => loading.value = false)
}
getData()



function handleDelete(row: DataSource) {
  apikeyDelete(row.id)
    .then(() => {
      datas.value.splice(datas.value.indexOf(row), 1)
      Message?.success('success')
    })

}

function handleRefresh(row: DataSource) {
  apikeyRefresh(row.id)
    .then(() => {
      Message?.success('success')
      getData()
    })
}
const formRef = ref<FormInst | null>(null)

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  expire_delta: {
    type: 'number',
    required: true,
    message: '请输入有效期',
    trigger: 'blur'
  }

}
function handleSubmitClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {

      apikeyPost(model)
        .then((response) => {
          Message?.success('success')
          showAPIKeyModal.value = false
          getData()
        })

    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}
</script>
  
  