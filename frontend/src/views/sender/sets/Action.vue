<template>
  <n-card :title="`${props.id ? '编辑' : '添加'}添加STMP凭证`">
    <n-skeleton text v-if="loading" :repeat="12" />
    <n-form ref="formRef" v-else :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '640px' }">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" :disabled="!!props.id" />
      </n-form-item>
      <n-form-item label="发送器:" path="sets">
        <n-select v-model:value="model.sets" multiple :options="options" />
      </n-form-item>

      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>
    </n-form>
  </n-card>
</template>
  
<script setup lang="ts">
import { NForm, NFormItem, NInput, NButton, NCard, NSelect } from 'naive-ui'
import type { FormInst, SelectOption, FormRules } from 'naive-ui'

import { reactive, ref } from "vue"
import { senderSetsPost, senderList, senderSetsPut,senderSetsGet } from "@/api/sender"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'

const props = defineProps({
  id: { type: Number, required: false }
})


const formRef = ref<FormInst | null>(null)
const model = reactive({
  name: '',
  sets: []
})


const loading = ref(false)
if (props.id) {
  loading.value = true;
  senderSetsGet(props.id)
    .then(response => {
      const data = response.data
      model.name = data.name
      model.sets = data.sets
      loading.value = false
    })

}

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  sets: {
    type: 'array',
    required: true,
    message: '请选择发送器',
    trigger: ['blur', 'change']
  }

}
const options = ref<SelectOption[]>()



senderList()
  .then((response) => {
    options.value = response.data.map((item: { index: string }) => {
      return { label: item.index, value: item.index } as SelectOption
    })
  })


const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      const promise = props.id ? senderSetsPut(props.id, model) : senderSetsPost(model)
      promise
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'SenderSets' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })
}
</script>
  
  