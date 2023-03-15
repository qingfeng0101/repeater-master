<template>
    <n-popconfirm @positive-click="handlePositiveClick" :positive-button-props="buttonProps"
        :negative-button-props="buttonProps">
        <template #trigger>
            <n-button :type="type" :size="size" :class="class">
                <slot></slot>
            </n-button>
        </template>
        {{ message }}
    </n-popconfirm>
</template>

<script lang="ts" setup>
import { NPopconfirm, NButton } from 'naive-ui'
import type { ButtonProps } from 'naive-ui'


interface Props {
    type: ButtonProps['type']
    message?: string,
    size: ButtonProps['size']
    class?: unknown
}
const props = withDefaults(defineProps<Props>(), {
    type: 'error',
    message: '请确认',
    size: 'small',
})
const buttonProps: ButtonProps = {
    size: 'tiny'
}
interface Emits {
    (e: 'click'): void
}

const emit = defineEmits<Emits>()
function handlePositiveClick() {
    emit('click')
}
function handleNegativeClick() {

}


</script>