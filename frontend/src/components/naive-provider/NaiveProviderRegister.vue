<script setup lang="ts">
import { useNotification, useDialog, useMessage, useLoadingBar } from 'naive-ui'
import { registerProviders } from '@/utils'


const notification = useNotification()
const dialog = useDialog()
const message = useMessage()
const loadingBar = useLoadingBar()

registerProviders({ notification, dialog, message, loadingBar })
</script>

<!-- eslint-disable-next-line vue/valid-template-root -->
<template></template>
