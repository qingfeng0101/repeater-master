<template>
  <n-config-provider :locale="locale" :date-locale="dateLocale" :theme-overrides="themeOverrides"
    :theme="themeStore.naiveTheme" class="h-full">
    <NaiveProvider>
      <router-view />
    </NaiveProvider>

  </n-config-provider>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router'
import { computed } from 'vue'
import { zhCN, enUS, dateEnUS, dateZhCN, NConfigProvider } from 'naive-ui'
import type { GlobalThemeOverrides } from 'naive-ui'
import { useThemeStore } from './store'
import useLocale from '@/hooks/locale'
import { NaiveProvider } from '@/components/naive-provider'

const { currentLocale } = useLocale()
const locale = computed(() => {
  switch (currentLocale.value) {
    case 'zh-CN':
      return zhCN
    case 'en-US':
      return enUS
    default:
      return enUS
  }
})

const dateLocale = computed(() => {
  switch (currentLocale.value) {
    case 'zh-CN':
      return dateZhCN
    case 'en-US':
      return dateEnUS
    default:
      return dateZhCN
  }
})

const themeStore = useThemeStore()

const themeOverrides: GlobalThemeOverrides = computed(
  (): GlobalThemeOverrides => ({
    "common": {
      ...themeStore.naiveThemeColors,
    },

  })
) as GlobalThemeOverrides


</script>
