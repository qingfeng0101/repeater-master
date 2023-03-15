<template>
  <n-card title="添加STMP凭证">
    <!-- <n-skeleton text v-if="loading" :repeat="12" /> -->
    <n-form ref="formRef" :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '640px' }">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" />
      </n-form-item>
      <n-form-item label="地址:" path="address">
        <n-input v-model:value="model.address" />
      </n-form-item>
      <n-form-item label="用户名:" path="username">
        <n-input v-model:value="model.username" />
      </n-form-item>
      <n-form-item label="密码:" path="password">
        <n-input v-model:value="model.password" />
      </n-form-item>
      <n-form-item label="认证类型:" path="auth_type">
        <n-select v-model:value="model.auth_type" placeholder="请选择" :options="typeOptions" />
      </n-form-item>
      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>
    </n-form>
  </n-card>
</template>
  
<script setup lang="ts">

import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NCard,

  NSelect
} from 'naive-ui'
import type { FormInst, SelectOption } from 'naive-ui'

import { reactive, ref } from "vue"
import { credentialStmpPost } from "@/api/stmp"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'


const formRef = ref<FormInst | null>(null)
const model = reactive({
  name: '',
  address: '',
  username: '',
  password: '',
  auth_type: 'LOGIN'
})

const rules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  address: {
    required: true,
    message: '请输入服务器地址',
    trigger: 'blur'
  },
  username: {
    required: true,
    message: '请输入用户',
    trigger: 'blur'
  },

  password: {
    required: true,
    message: '请输入密码',
    trigger: 'blur'
  },
  auth_type: {
    required: true,
    trigger: ['blur', 'change'],
    message: '请选择类型',
  },

}


const typeOptions: SelectOption[] = [{ label: 'LOGIN', value: 'LOGIN' }, { label: 'PLAIN', value: 'PLAIN' }, { label: 'CRAM-MD5', value: 'CRAM-MD5' }]


const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      credentialStmpPost(model)
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'CredentialStmp' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}
</script>
  
  