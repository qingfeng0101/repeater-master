<template>
  <n-card :title="`${props.id ? '编辑' : '添加'}Prometheus数据源`">
    <n-skeleton text v-if="loading" :repeat="12" />
    <n-form v-else ref="formRef" :model="model" label-placement="left" :label-width="100" :rules="rules"
      class="max-w-95%">
      <n-form-item label="名称:" path="name">
        <n-input v-model:value="model.name" />
      </n-form-item>

      <n-form-item label="LabelSetList:" path="label_set_list">
        <n-dynamic-input v-model:value="model.label_set_list" :on-create="() => { return {} }">
          <template #create-button-default>
            添加
          </template>
          <template #default="{ index }">
            <pre class="flex-1 text-18px m-a">{{ JSON.stringify(model.label_set_list[index]) }}</pre>
            <n-button text @click="edit(index)" class="text-18px">
              <n-icon>
                <EditBoxLine />
              </n-icon>
            </n-button>
          </template>
        </n-dynamic-input>
      </n-form-item>

      <n-form-item label="SeverityList:" path="severity_list">
        <n-dynamic-tags v-model:value="model.severity_set" />
      </n-form-item>
      <n-form-item label="备注: " path="comment">
        <n-input v-model:value="model.comment" type="textarea" :autosize="{ minRows: 3, maxRows: 5 }" />
      </n-form-item>
      <div class="flex justify-end">
        <n-button @click="submitButtonClick" type="primary">提交</n-button>
      </div>

      <n-modal :show="showModal">
        <n-card class="w-800px" :bordered="false">
          <template #header>
            <pre class="m-a">{{ JSON.stringify(label_set_list) }}</pre>
          </template>
          <KVEdit v-model="label_set_list" />
          <div class="flex justify-end m-t-12px">
            <n-button @click="showModal = false" type="primary">完成</n-button>
          </div>
        </n-card>

      </n-modal>

    </n-form>
  </n-card>
</template>

<script setup lang="ts">
import { NForm, NFormItem, NInput, NButton, NCard, NSkeleton, NDynamicInput, NModal, NIcon, NDynamicTags } from 'naive-ui'
import type { FormInst } from 'naive-ui'

import { reactive, ref, computed } from "vue"
import { prometheusDatasourceFakePost, prometheusDatasourceFakePut, prometheusDatasourceFakeGet } from "@/api/prometheus"
import KVEdit from './components/KVEdit.vue'
import EditBoxLine from '@/assets/icons/EditBoxLine.vue'
import { useRouter } from 'vue-router'
import { Message } from '@/utils'

const props = defineProps({
  id: { type: Number, required: false }
})

// const label_set_list = ref<Record<string, string>>({})

const label_set_index = ref<number>()
const formRef = ref<FormInst | null>(null)

interface ModelType {
  name: string
  label_set_list: Record<string, string>[]
  severity_set: string[],
  comment: string
}

const showModal = ref(false)
const model = reactive<ModelType>({
  name: '',
  label_set_list: [{}],
  severity_set: [],
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

function edit(index: number) {
  showModal.value = true
  label_set_index.value = index
}


if (props.id) {
  loading.value = true;
  prometheusDatasourceFakeGet(props.id)
    .then(response => {
      const data = response.data
      model.name = data.name

      model.label_set_list = data.label_set_list
      model.severity_set = data.severity_set
      model.comment = data.comment

      loading.value = false
    })
}
const label_set_list = computed({
  get() {
    return model.label_set_list[label_set_index.value as number]
  },
  set(newValue: Record<string, string>) {
    model.label_set_list[label_set_index.value as number] = newValue
  }
})
const router = useRouter()

function submitButtonClick() {

  formRef.value?.validate((errors) => {
    if (!errors) {
      const promise = props.id ? prometheusDatasourceFakePut(props.id, model) : prometheusDatasourceFakePost(model)
      promise
        .then((response) => {
          Message?.success('success')
          router.push({ name: 'FakePrometheusDatasourceList' })
        })
    } else {
      console.log('errors')
      Message?.error('验证失败')
    }
  })

}

</script>

