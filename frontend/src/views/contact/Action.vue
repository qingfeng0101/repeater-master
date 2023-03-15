<template>
  <n-card :title="`${props.id ? '编辑' : '添加'}通讯录`">
    <n-skeleton text v-if="loading" :repeat="8" />
    <n-form v-else ref="formRef" :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '640px' }">
      <n-form-item label="用户名: " path="username">
        <n-input v-model:value="model.username" />
      </n-form-item>
      <n-form-item label="姓名: " path="name">
        <n-input v-model:value="model.name" />
      </n-form-item>
      <n-form-item label="邮箱: " path="email">
        <n-input v-model:value="model.email" />
      </n-form-item>
      <n-form-item label="手机: " path="mobile">
        <n-input v-model:value="model.mobile" />
      </n-form-item>
      <n-form-item label="企业微信: " path="wecom" class="wecom">
        <n-switch class="self-start mb-4px" v-model:value="single">
          <template #checked>single</template>
          <template #unchecked>multiple</template>
        </n-switch>
        <n-input v-model:value="model.wecom" v-if="single" />
        <n-dynamic-input v-model:value="modelWecom" v-else :on-create="onCreate" #="{ index, value }">
          <div class="flex">
            <!-- <n-input v-model:value="modelWecom[index].name" placeholder="Name" @keydown.enter.prevent /> -->
            <n-select v-model:value="modelWecom[index].name" :options="options" class="min-w-240px" />
            <div class="m-x-8px h-34px lh-34px"> : </div>
            <n-input v-model:value="modelWecom[index].value" placeholder="Value" />
          </div>
        </n-dynamic-input>
      </n-form-item>

      <n-form-item label="钉钉名称: " path="wecom" class="wecom">
        <n-switch class="self-start mb-4px" v-model:value="ding">
          <template #checked>single</template>
          <template #unchecked>multiple</template>
        </n-switch>
        <n-input v-model:value="model.dingtalk" v-if="single" />
        <n-dynamic-input v-model:value="modelDing" v-else :on-create="onCreate" #="{ index, value }">
          <div class="flex">
            <!-- <n-input v-model:value="modelWecom[index].name" placeholder="Name" @keydown.enter.prevent /> -->
            <n-select v-model:value="modelDing[index].name" :options="options" class="min-w-240px" />
            <div class="m-x-8px h-34px lh-34px"> : </div>
            <n-input v-model:value="modelDing[index].value" placeholder="Value" />
          </div>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item label="飞书名称: " path="wecom" class="wecom">
        <n-switch class="self-start mb-4px" v-model:value="lark">
          <template #checked>single</template>
          <template #unchecked>multiple</template>
        </n-switch>
        <n-input v-model:value="model.lark" v-if="single" />
        <n-dynamic-input v-model:value="modelLark" v-else :on-create="onCreate" #="{ index, value }">
          <div class="flex">
            <!-- <n-input v-model:value="modelWecom[index].name" placeholder="Name" @keydown.enter.prevent /> -->
            <n-select v-model:value="modelLark[index].name" :options="options" class="min-w-240px" />
            <div class="m-x-8px h-34px lh-34px"> : </div>
            <n-input v-model:value="modelLark[index].value" placeholder="Value" />
          </div>
        </n-dynamic-input>
      </n-form-item>
      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>
    </n-form>
  </n-card>
</template>

<script setup lang="ts">
import { NForm, NFormItem, NInput, NButton, NCard, NSkeleton, NSwitch, NDynamicInput, NSelect } from 'naive-ui'
import type { FormInst, SelectOption } from 'naive-ui'
import { reactive, ref } from "vue"
import { contactGet, contactPost, contactPut } from '@/api/contact'
import { senderList } from "@/api/sender"
import { useRouter } from 'vue-router'
import { Message } from '@/utils'

interface DataSource {
  "name": string
  "username": string
  "email": string
  "mobile": string
  "wecom": string
  "dingtalk": string
  "lark" : string
}

const formRef = ref<FormInst | null>(null)
const model = reactive({
  name: '',
  username: '',
  email: '',
  mobile: '',
  wecom: '',
  dingtalk: '',
  lark: '',
})

const rules = {
  username: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur'
  },
  email: {
    required: true,
    message: '请输入邮箱',
    trigger: 'blur'
  },
  name: {
    required: true,
    message: '请输入姓名',
    trigger: 'blur'
  },
  mobile: {
    required: true,
    message: '请输入手机号',
    trigger: 'blur'
  },
}

const single = ref(true)
const ding = ref(true)
const lark = ref(true)
interface MultiWecom {
  "name": string
  "value": string
}
function onCreate() {
  return {
    name: '',
    value: ''
  }
}

const options = ref<SelectOption[]>()



senderList()
  .then((response) => {
    options.value = response.data.filter((d: any) => d.type === 'wecom').map((item: { index: string }) => {
      return { label: item.index, value: item.index } as SelectOption
    })
  })
const modelWecom = ref<MultiWecom[]>([])
const modelDing = ref<MultiWecom[]>([])
  const modelLark = ref<MultiWecom[]>([])
const props = defineProps({
  id: { type: Number, required: false }
})

const loading = ref(false)

if (props.id) {
  loading.value = true;
  contactGet(props.id)
    .then(response => {
      const data = response.data as DataSource
      model.username = data.username
      model.name = data.name
      model.email = data.email
      model.mobile = data.mobile
      model.wecom = data.wecom
      model.dingtalk = data.dingtalk
      model.lark = data.lark

      try {
        let wecom = JSON.parse(model.wecom)
        let dingtalk = JSON.parse(model.dingtalk)
        let lark = JSON.parse(model.lark)

        for (let i in wecom) {
          modelWecom.value.push({ name: i, value: wecom[i] })
        }
        for (let i in dingtalk) {
          modelDing.value.push({ name: i, value: dingtalk[i] })
        }
        for (let i in lark) {
          modelLark.value.push({ name: i, value: lark[i] })
        }
      } catch (err) { console.log(err) }

      loading.value = false
    })

}
function getMultipleWecom(value: MultiWecom[]): string {
  let wecom: Record<string, string> = {}
  for (let i of value) {
    wecom[i.name] = i.value
  }

  if (Object.keys(wecom).length > 0) {
    return JSON.stringify(wecom)
  }
  return ''
}

const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      if (!single.value) {
        model.wecom = getMultipleWecom(modelWecom.value)
      }
      const promise = props.id ? contactPut(props.id, model) : contactPost(model)
      promise
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'ContactList' })
        })

    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}
</script>

<style lang="scss">
.wecom {
  .n-form-item-blank {
    flex-direction: column;
  }
}
</style>