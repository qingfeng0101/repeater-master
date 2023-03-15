<template>
  <n-card :title="`${props.id ? '编辑' : '添加'}Prometheus数据源`">
    <n-skeleton text v-if="loading" :repeat="12" />
    <n-form ref="formRef" :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '640px' }">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" />
      </n-form-item>
      <n-form-item label="BaseUrl:" path="base_url">
        <n-input v-model:value="model.base_url" />
      </n-form-item>
      <n-form-item label="Alerts:" path="alerts">
        <n-input v-model:value="model.alerts" />
      </n-form-item>
      <n-form-item label="Targets:" path="targets">
        <n-input v-model:value="model.targets" />
      </n-form-item>
      <n-form-item label="Rules:" path="rules">
        <n-input v-model:value="model.rules" />
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
import { NForm, NFormItem, NInput, NButton, NCard, NSkeleton } from 'naive-ui'
import type { FormInst } from 'naive-ui'

import { reactive, ref } from "vue"
import { prometheusDatasourcePost, prometheusDatasourcePut, prometheusDatasourceGet } from "@/api/prometheus"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'
const props = defineProps({
  id: { type: Number, required: false }
})

const formRef = ref<FormInst | null>(null)
const model = reactive({
  name: '',
  base_url: '',
  alerts: '',
  targets: '',
  rules: '',
  comment: ''
})

const rules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  base_url: {
    required: true,
    message: '请输入BaseUrl',
    trigger: 'blur'
  },
}

const loading = ref(false)

if (props.id) {
  loading.value = true;
  prometheusDatasourceGet(props.id)
    .then(response => {
      const data = response.data
      model.name = data.name
      model.base_url = data.base_url
      model.alerts = data.alerts
      model.targets = data.targets
      model.rules = data.rules
      model.comment = data.comment

      loading.value = false
    })
}

const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      const promise = props.id ? prometheusDatasourcePut(props.id, model) : prometheusDatasourcePost(model)
      promise
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'PrometheusDatasourceList' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}
</script>
  
  