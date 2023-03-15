<template>

  <n-result class="absolute top-50% left-50% translate--50%" status="404" :subtitle="'not found'">
    <template #footer>
      <n-button @click="back">back</n-button>
    </template>
  </n-result>

</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { NResult, NButton } from 'naive-ui'
const router = useRouter()
const back = () => {
  // warning： Go to the node that has the permission
  router.push({ name: 'root' })
}
</script>

<style scoped lang="less">
</style>
