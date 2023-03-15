<template>
    <n-card>
        <div v-for="v, k, idx in model" :key="k"
            :class="{ flex: true, 'mb-8px': idx !== Object.keys(model).length - 1 }">
            <n-text class="w-64px">{{ k }}:</n-text>
            <n-checkbox-group v-model:value="valueList[k]">
                <n-space item-style="display: flex">
                    <n-checkbox :value="item" :label="item" v-for="item in v" :key="item" />
                </n-space>
            </n-checkbox-group>
        </div>
    </n-card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { NCheckboxGroup, NCheckbox, NSpace, NText, NCard } from 'naive-ui'

type modelType = Record<string, string[]>
interface Props {
    model: modelType
    modelValue: modelType,
}



const props = defineProps<Props>()

interface Emits {
    (e: 'update:modelValue', modelValue: modelType): void
}

const emit = defineEmits<Emits>()


const valueList = computed({
    get() {
        return props.modelValue
    },
    set(newValue: modelType) {
        emit('update:modelValue', newValue)
    }
})

</script>