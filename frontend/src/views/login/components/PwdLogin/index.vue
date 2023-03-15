<template>
  <n-form ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <n-form-item path="userName">
      <n-input v-model:value="model.username" placeholder="请输入用户名" />
    </n-form-item>
    <n-form-item path="password">
      <n-input v-model:value="model.password" type="password" show-password-on="click" placeholder="请输入密码"
        @keyup.enter="handleSubmit" />
    </n-form-item>
    <n-space :vertical="true" :size="24">
      <div class="flex-y-center justify-between">
        <n-checkbox v-model:checked="rememberMe">记住我</n-checkbox>
        <n-button :text="true" @click="Alert">忘记密码？</n-button>
      </div>
      <n-button type="primary" size="large" :block="true" :round="true" :loading="loading" @click="handleSubmit">
        确定
      </n-button>
    </n-space>
  </n-form>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { NForm, NFormItem, NInput, NSpace, NCheckbox, NButton } from 'naive-ui'
import type { FormInst, FormRules } from 'naive-ui'

import { useMainStore } from '@/store'
import { Message } from '@/utils'
import { useRoute, useRouter } from 'vue-router'

const { Login } = useMainStore()
const router = useRouter()
const route = useRoute()

const formRef = ref<HTMLElement & FormInst>()

const model = reactive({
  username: 'admin',
  password: ''
})

const loading = ref(false)

const rules: FormRules = {
  username: { required: true, message: '请输入密码' },
  password: { required: true, message: '请输入密码' }
}

const rememberMe = ref(false)

async function handleSubmit() {
  await formRef.value?.validate()
  const { username, password } = model
  loading.value = true
  Login(username, password)
    .then((ret) => {
      if (ret) {
        Message?.success("登录成功")
        router.push(route.query.redirect as string || "/")
      }
    })
    .finally(() => {
      loading.value = false
    })
}

function Alert() {
  alert('请联系管理员')
}

</script>

<style scoped>

</style>
