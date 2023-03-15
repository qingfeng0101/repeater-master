<template>
    <n-form ref="formRef" :model="model" :rules="rules" class="max-w-640px">
        <n-form-item path="password" label="当前密码">
            <n-input v-model:value="model.password" type="password" />
        </n-form-item>
        <n-form-item path="new_password" label="新密码">
            <n-input v-model:value="model.new_password" type="password" @input="handlePasswordInput"
                @keydown.enter.prevent />
        </n-form-item>
        <n-form-item ref="rPasswordFormItemRef" first path="confirm_password" label="重复新密码">
            <n-input v-model:value="model.confirm_password" :disabled="!model.new_password" type="password"
                @keydown.enter.prevent />
        </n-form-item>
        <div style="display: flex; justify-content: center;">
            <n-button round type="primary" @click="handleValidateButtonClick">
                修改
            </n-button>
        </div>
        <!-- <n-row :gutter="[0, 24]">
            <n-col :span="24">
                <div style="display: flex; justify-content: flex-end">
                    <n-button round type="primary" @click="handleValidateButtonClick">
                        修改
                    </n-button>
                </div>
            </n-col>
            
        </n-row> -->
    </n-form>


</template>
  
<script lang="ts" setup >
import { ref, reactive } from 'vue'
import type { FormInst, FormItemInst, FormItemRule, FormRules } from 'naive-ui'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui';
import { Message } from '@/utils'
import { useMainStore } from "@/store";
import { userPasswordPatch } from '@/api/user'
interface ModelType {
    password: string
    new_password: string
    confirm_password: string
}


const formRef = ref<FormInst | null>(null)
const rPasswordFormItemRef = ref<FormItemInst | null>(null)

const model = reactive<ModelType>({
    password: '',
    new_password: '',
    confirm_password: ''
})
function validatePasswordStartWith(
    rule: FormItemRule,
    value: string
): boolean {
    return (
        !!model.new_password &&
        model.new_password.startsWith(value) &&
        model.new_password.length >= value.length
    )
}
function validatePasswordSame(rule: FormItemRule, value: string): boolean {
    return value === model.new_password
}
const rules: FormRules = {
    password: [
        {
            required: true,
            message: '请输入密码'
        }
    ],
    new_password: [
        {
            required: true,
            message: '请输入新密码'
        }
    ],
    confirm_password: [
        {
            required: true,
            message: '请再次输入新密码',
            trigger: ['input', 'blur']
        },
        {
            validator: validatePasswordStartWith,
            message: '两次新密码输入不一致',
            trigger: 'input'
        },
        {
            validator: validatePasswordSame,
            message: '两次新密码输入不一致',
            trigger: ['blur', 'password-input']
        }
    ]
}


function handlePasswordInput() {
    if (model.confirm_password) {
        rPasswordFormItemRef.value?.validate({ trigger: 'password-input' })
    }
}
const store = useMainStore()
function handleValidateButtonClick(e: MouseEvent) {
    e.preventDefault()
    formRef.value?.validate((errors) => {
        if (!errors) {
            userPasswordPatch(store.account, model)
                .then(() => Message?.success("密码修改成功"))
        } else {
            console.log(errors)
            Message?.error('验证失败')
        }
    })
}


</script>