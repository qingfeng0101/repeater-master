<template>
  <div class="relative flex-center wh-full" :style="{ backgroundColor: bgColor }">
    <dark-mode-switch v-model:dark="themeStore.darkMode" class="absolute left-48px top-24px z-3 text-20px" />

    <n-card :bordered="false" size="large" class="z-4 !w-auto rounded-20px shadow-sm">
      <div class="w-300px sm:w-360px">
        <header class="flex-y-center justify-between">
          <div class="w-70px h-70px text-primary overflow-hidden">
            <n-icon :size="70">
              <CrisisAlert />
            </n-icon>
          </div>
          <n-gradient-text type="primary" :size="28">{{ '通知分发器' }}</n-gradient-text>
        </header>
        <main class="pt-24px">
          <h3 class="text-18px text-primary font-medium">{{ activeModule.label }}</h3>
          <div class="pt-24px">
            <transition name="fade-slide" mode="out-in" appear>
              <component :is="activeModule.component" />
            </transition>
          </div>
        </main>
      </div>
    </n-card>
    <login-bg :theme-color="bgThemeColor" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Component } from 'vue'
import DarkModeSwitch from '@/components/DarkModeSwitch.vue'
import CrisisAlert from '@/assets/icons/CrisisAlert.vue'
import { useThemeStore } from '@/store'
import { NCard, NGradientText, NIcon } from 'naive-ui'

import { getColorPalette, mixColor } from '@/utils'
import { LoginBg, PwdLogin, } from './components'

interface Props {
  /** 登录模块分类 */
  module: string
}

const props = defineProps<Props>()

const themeStore = useThemeStore()


interface LoginModule {
  key: string
  label: string
  component: Component
}

const modules: LoginModule[] = [
  { key: 'pwd-login', label: '密码登录', component: PwdLogin },

]

const activeModule = computed(() => {
  const active: LoginModule = { ...modules[0] }
  const findItem = modules.find(item => item.key === props.module)
  if (findItem) {
    Object.assign(active, findItem)
  }
  return active
})

const bgThemeColor = computed(() => (themeStore.darkMode ? getColorPalette(themeStore.themeColor, 7) : themeStore.themeColor))

const bgColor = computed(() => {
  const COLOR_WHITE = '#ffffff'
  const ratio = themeStore.darkMode ? 0.5 : 0.2
  return mixColor(COLOR_WHITE, themeStore.themeColor, ratio)
})
</script>

<style scoped>

</style>
