
  <template>
  <n-card title="添加lark凭证">
    <!-- <n-skeleton text v-if="loading" :repeat="12" /> -->
    <n-form ref="formRef" :model="model" label-placement="left" :label-width="160" :rules="rules"
      :style="{ maxWidth: '640px' }">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" />
      </n-form-item>
      <n-form-item label="应用Id:" path="corp_id">
        <n-input v-model:value="model.corp_id" />
      </n-form-item>
      <n-form-item label="应用Secret:" path="corp_secret">
        <n-input v-model:value="model.corp_secret" />
      </n-form-item>
      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>
    </n-form>
  </n-card>
</template>
<script setup lang="ts">
import { NForm, NFormItem, NInput, NButton, NCard, NInputNumber } from 'naive-ui'
import type { FormInst, FormRules } from 'naive-ui'

import { reactive, ref } from "vue"
import { credentialLarkPost } from "@/api/stmp"

import { useRouter } from 'vue-router'
import { Message } from '@/utils'


const formRef = ref<FormInst | null>(null)
const model = reactive({
  name: '',
  corp_id: '',
  corp_secret: '',
})

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入名称',
    trigger: 'blur'
  },
  corp_id: {
    required: true,
    message: '请输入应用Key',
    trigger: 'blur'
  },

  corp_secret: {
    required: true,
    message: '请输入应用Secret',
    trigger: 'blur'
  },


}

const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      credentialLarkPost(model)
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'CredentialLark' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}
</script>
  
  