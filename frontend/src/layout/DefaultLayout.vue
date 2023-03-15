<template>
  <n-layout position="absolute">
    <n-layout-header class="h-56px" bordered>
      <Header />
    </n-layout-header>
    <n-layout has-sider position="absolute" style="top: 56px;">
      <n-layout-sider bordered :width="240" collapse-mode="width" :collapsed-width="64"
        v-model:collapsed="store.collapsed" :native-scrollbar="false">
        <n-menu :collapsed="store.collapsed" @update:value="handleUpdateValue" :options="menuOptions"
          v-model:value="active" default-expand-all />
      </n-layout-sider>
      <n-layout-content content-style="padding: 16px; height: 100%" embedded :native-scrollbar="false"  >
        <!-- <router-view :key="$route.fullPath" /> -->
        <router-view v-slot="{ Component, route }">
          <transition name="fade" mode="out-in" appear>
            <component :is="Component" :key="route.fullPath" />
          </transition>
        </router-view>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>
<script setup lang="ts">
import { NLayout, NLayoutHeader, NLayoutSider, NLayoutContent, NMenu, NIcon } from 'naive-ui'


import type { MenuOption } from 'naive-ui'

import { h, ref } from 'vue'


import { useRouter, useRoute } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useMainStore } from '@/store'

import Header from './common/Header.vue'
import icons from '@/assets/icons'
import { routes } from "@/router"

const menuOptions = ref<MenuOption[]>()

// const expandedKeys: string[] = []

const router = useRouter()
const route = useRoute()

const active = ref<string | null>(null)
active.value = route.name as string
const store = useMainStore()



function handleUpdateValue(key: string) {
  router.push({ name: key })
}

function hasChildren(item: MenuOption | RouteRecordRaw) {
  return Boolean(item.children && item.children.length)
}

function checkPermission(item: RouteRecordRaw, userLevel: string) {
  if (!item.meta?.roles || item.meta?.roles?.includes(userLevel)) {
    return true
  }
  return false
}

function setIcon(iconName: string) {
  const icon = icons[iconName]
  if (icon) {
    return () => h(NIcon, null, { default: () => h(icon) })
  }
}



function getMenu(routes: RouteRecordRaw[] | undefined, userLevel: string): MenuOption[] {
  const menu: MenuOption[] = []
  if (!routes) {
    return menu
  }
  for (let i of routes) {
    if (!i.meta || i.meta.hideInMenu || !checkPermission(i, userLevel)) {
      continue
    }
    let item: MenuOption = { label: i.meta.locale, key: i.name as string, icon: setIcon(i.meta.icon as string), }
    if (hasChildren(i)) {
      item.children = getMenu(i.children, userLevel)
    }
    if (i.meta.flatChildrenInMenu && hasChildren(item)) {
      menu.push(...item.children as MenuOption[])
    } else {
      menu.push(item)
    }
  }

  return menu
}

menuOptions.value = getMenu(routes, store.userClaims.level)

</script>

