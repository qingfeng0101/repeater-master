<template>
  <n-loading-bar-provider>
    <n-dialog-provider ref="dialogProviderRef">
      <n-notification-provider ref="notificationProviderRef">
        <n-message-provider ref="messageProviderRef">
          <NaiveProviderRegister />
          <slot></slot>
        </n-message-provider>
      </n-notification-provider>
    </n-dialog-provider>
  </n-loading-bar-provider>


</template>

<script setup lang="ts">

import type { NotificationProviderInst, DialogProviderInst, MessageProviderInst } from 'naive-ui'
import { NNotificationProvider, NLoadingBarProvider, NDialogProvider, NMessageProvider } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import NaiveProviderRegister from './NaiveProviderRegister.vue'


const notificationProviderRef = ref<NotificationProviderInst>()
const dialogProviderRef = ref<DialogProviderInst>()
const messageProviderRef = ref<MessageProviderInst>()
const router = useRouter()

onMounted(() => {
  // 每次路由切换时销毁当前的实例
  router.afterEach(() => {
    notificationProviderRef.value!.destroyAll()
    dialogProviderRef.value!.destroyAll()
    messageProviderRef.value!.destroyAll()
  })
})
</script>