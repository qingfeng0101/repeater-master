<template>
  <n-card :title="`${props.id ? '编辑' : '添加'}通知规则`" segmented>
    <template #header-extra>
      <n-tooltip trigger="hover">
        <template #trigger>
          <n-select v-model:value="source_name" :options="options" />
        </template>
        选择数据源
      </n-tooltip>

    </template>
    <n-skeleton text v-if="loading" :repeat="12" />
    <n-form ref="formRef" v-else :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '800px' }">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" :disabled="!!props.id" />
      </n-form-item>
      <n-form-item label="规则匹配标签:" path="label_set_list">
        <n-spin :show="labelSetLoading" class="wfull">
          <n-transfer v-model:value="model.label_set_list" virtual-scroll :options="labelSetOptions" source-filterable
            :render-source-label="renderLabel" :render-target-label="renderLabel" />
        </n-spin>
      </n-form-item>
      <n-form-item label="通知目标:" path="contact_ids">
        <n-transfer v-model:value="model.contact_ids" virtual-scroll source-filterable :options="contactOptions" />
      </n-form-item>
      <n-form-item label="通知发送器集合:" path="sender_set_id">
        <n-select v-model:value="model.sender_set_id" :options="senderSetOptions" />
      </n-form-item>
      <n-form-item label="严重程度发送规则:">
        <n-spin :show="labelSetLoading" class="wfull">
          <DynamicChecbox v-model="prometheus.sender_filter" :model="senderFilterData" />
        </n-spin>
      </n-form-item>
      <n-form-item label="备注: " path="comment">
        <n-input v-model:value="model.comment" type="textarea" :autosize="{ minRows: 3, maxRows: 5 }" />
      </n-form-item>
      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>
    </n-form>
  </n-card>

</template>
 
<script setup lang="ts">
import { h, reactive, ref, watch, computed } from "vue"
import { NForm, NFormItem, NInput, NButton, NCard, NSelect, NTooltip, NTransfer, NEllipsis, NSpin, NSkeleton } from 'naive-ui'
import type { FormInst, SelectOption, TransferOption, FormRules, TransferRenderSourceLabel } from 'naive-ui'
import DynamicChecbox from './components/DynamicChecbox.vue'

import { notifyRulesGet, notifyRulesPost, notifyRulesPut, notifyRulesCacheTree } from "@/api/notify"

import { prometheusLabelSetTarget, prometheusDatasourceList, prometheusDatasourceFakeList, prometheusSeverityTarget } from '@/api/prometheus'
import { senderList } from "@/api/sender"
import { contactList } from '@/api/contact'
import { senderSetsList } from "@/api/sender"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
import { usePrometheusStore } from '@/store'
const props = defineProps({
  id: { type: Number, required: false }
})


const prometheus = usePrometheusStore()
const formRef = ref<FormInst | null>(null)

const model = reactive({
  name: '',
  label_set_list: [],
  contact_ids: [],
  sender_filter: {},
  sender_set_id: undefined,
  comment: ''
})
let label_set_list: any[] = []
const loading = ref(false)
if (props.id) {
  loading.value = true;
  notifyRulesGet(props.id)
    .then(response => {
      const data = response.data
      model.name = data.name

      label_set_list = data.label_set_list

      model.sender_filter = data.sender_filter
      for (let key of Object.keys(data.sender_filter)) {
        prometheus.sender_filter[key] = data.sender_filter[key]
      }
      model.sender_set_id = data.sender_set_id

      for (let i of data.contacts) {
        model.contact_ids.push(i.id as never)
      }
      loading.value = false
    })

}

const source_name = ref('')

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  contact_ids: {
    type: 'array',
    required: true,
    message: '请选择通知目标',
    trigger: ['blur', 'change'],
  },
  sender_set_id: {
    type: 'number',
    required: true,
    trigger: ['blur', 'change'],
    message: '请选择通知分发集合',
  }
}

const options = ref<SelectOption[]>([])
const labelSetOptions = ref<any[]>()
const contactOptions = ref<TransferOption[]>()
const senderSetOptions = ref<SelectOption[]>()
let treeCache = {}
notifyRulesCacheTree()
  .then(response => {
    treeCache = response.data
  })
prometheusDatasourceList()
  .then((response) => {
    for (let item of response.data) {
      if (!source_name.value) source_name.value = item.name
      options.value.push({ label: item.name, value: item.name } as SelectOption)
    }
  })

prometheusDatasourceFakeList()
  .then((response) => {
    for (let item of response.data) {
      if (!source_name.value) source_name.value = item.name
      options.value.push({ label: item.name, value: item.name } as SelectOption)
    }
  })

contactList()
  .then((response) => {
    contactOptions.value = response.data.map((item: any) => {
      return { label: item.name, value: item.id }
    })
  })

senderSetsList()
  .then((response) => {
    senderSetOptions.value = response.data.map((item: any) => {
      return { label: item.name, value: item.id }
    })
  })

const senderSet = reactive(new Set<string>())
senderList()
  .then(response => {
    for (let i of response.data) {
      senderSet.add(i.type)
    }
  })

const severitys = ref<string[]>([])
const labelSetLoading = ref(false)
watch(source_name, (name) => {
  labelSetLoading.value = true
  rules.label_set_list = []
  prometheusLabelSetTarget({ source: name })
    .then(response => {
      labelSetOptions.value = response.data.map((item: any) => {

        return { label: item.service_name || item.namespace || item.instance, value: item, disabled: checkExist(item, treeCache, label_set_list) }
      })
      labelSetLoading.value = false
    })

  prometheusSeverityTarget({ source: name })
    .then(response => {
      severitys.value = response.data.severity_set
    })
})

const senderFilterData = computed(() => {
  let filterData: Record<string, string[]> = {}
  for (let i of severitys.value) {
    filterData[i] = Array.from(senderSet)
  }
  return filterData
})

const renderLabel: TransferRenderSourceLabel = ({ option }) => {
  return h(NTooltip, { delay: 1500, trigger: 'hover' }, {
    trigger: () => h('div', h(NEllipsis, { tooltip: false }, { default: () => option.label })),
    default: () => JSON.stringify(option.value)
  })
}

const router = useRouter()

function submitButtonClick() {
  formRef.value?.validate((errors) => {
    console.log(model)
    if (!errors) {
      let sender_filter: Record<string, string[]> = {}
      for (let key of Object.keys(senderFilterData.value)) {
        sender_filter[key] = prometheus.sender_filter[key]
      }
      model.sender_filter = sender_filter
      const promise = props.id ? notifyRulesPut(props.id, model) : notifyRulesPost(model)
      promise
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'PrometheusNotifyRule' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })
}


function checkExist(item: any, tree: any, extend: any[]) {
  let equal = (item1: any, item2: any) => {
    for (let i in item1) {
      if (item1[i] !== item2[i]) {
        return false
      }
    }
    return true
  }
  for (let i of extend) {
    if (equal(item, i)) {
      return false
    }
  }
  let keys = []
  for (let i in item) {
    keys.push(`${i}=${item[i]}`)
  }

  for (let i = 0; i < keys.length; i++) {
    if (Object.keys(tree).indexOf("children") === -1) {
      break
    }
    for (let j of tree.children) {
      if (keys[i] == j.key) {
        tree = j
        i = -1
        break
      }
    }
  }
  return !!tree.value

}
</script>

