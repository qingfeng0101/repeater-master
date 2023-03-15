<template>
    <n-card>
        <n-dynamic-input v-model:value="valueList" preset="pair" key-placeholder="name" value-placeholder="value" />
    </n-card>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import { NDynamicInput, NCard } from 'naive-ui'

type modelType = Record<string, string>
interface Props {
    modelValue: modelType
}

interface Emits {
    (e: 'update:modelValue', modelValue: modelType): void
}

const emit = defineEmits<Emits>()


const props = defineProps<Props>()
type valueType = {
    key: string;
    value: string;
}[]
const value = ref<valueType>([])
for (const key in props.modelValue) {
    value.value.push({ key: key, value: props.modelValue[key] })
}


const valueList = computed({
    get() {
        return value.value
    },
    set(newValue: valueType) {
        value.value = newValue
        let ret: modelType = {}
        for (let i of newValue) {
            ret[i.key] = i.value
        }
        emit('update:modelValue', ret)
    }
})
</script>