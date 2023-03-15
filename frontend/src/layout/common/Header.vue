<template>
    <div class="header-container flex-y-center h-full">
        <div class="w-240px flex-center text-primary">
            <n-icon size="32">
                <CrisisAlert />
            </n-icon>
            <n-gradient-text :size="28" class="font-600 ml-8px">通知分发器
            </n-gradient-text>
        </div>
        <div class="flex-1 flex-y-center cursor-pointer">
            <n-icon size="18" @click="store.collapsed = !store.collapsed">
                <MenuFoldRight v-if="store.collapsed" />
                <MenuFoldLeft v-else />
            </n-icon>
        </div>
        <div class="flex-center h-full">
            <dark-mode-switch v-model:dark="themeStore.darkMode" class="w-40px flex-x-center" />
            <n-popover class="!p-0" trigger="click" placement="bottom"  >
                <template #trigger>
                    <n-icon size="18" class="w-40px cursor-pointer">
                        <Theme />
                    </n-icon>
                </template>
                <ThemeColorSelect/>
            </n-popover>

            <n-dropdown :options="droOptions" @select="handleDropdown" trigger="click">
                <div class="flex-center px-12px cursor-pointer">
                    <n-icon size="32">
                        <DefaultAvatar />
                    </n-icon>
           
                    <span class="pl-8px text-18px font-500">{{ store.userClaims.name }}</span>
                </div>
            </n-dropdown>
        </div>
    </div>
</template>
<script setup lang="ts">
import { h } from 'vue'
import type { Component } from 'vue'
import { useRouter } from 'vue-router'
import { NGradientText, NIcon, NDropdown, NPopover } from 'naive-ui'
import type { DropdownOption} from 'naive-ui'
import DarkModeSwitch from '@/components/DarkModeSwitch.vue'
import ThemeColorSelect from '@/components/ThemeColorSelect.vue'

import CrisisAlert from '@/assets/icons/CrisisAlert.vue'
import DefaultAvatar from '@/assets/icons/DefaultAvatar.vue'
import MenuFoldLeft from '@/assets/icons/MenuFoldLeft.vue'
import MenuFoldRight from '@/assets/icons/MenuFoldRight.vue'
import Password from '@/assets/icons/Password.vue'
import Logout from '@/assets/icons/Logout.vue'
import Theme from '@/assets/icons/Theme.vue'
import { useThemeStore, useMainStore } from '@/store'

const store = useMainStore()
const themeStore = useThemeStore()

function renderIcon(icon:Component) {
    return () => h(NIcon, null, { default: () => h(icon) })
}

// const UserAvatar = icons['UserAvatar']
// const Logout = icons['Logout']
// const Theme = icons['Theme']
const droOptions: DropdownOption[] = [
    {
        label: '修改密码',
        key: 'change-password',
        icon: renderIcon(Password)
    },
    {
        type: 'divider',
        key: 'divider'
    },
    {
        label: '退出登录',
        key: 'logout',
        icon: renderIcon(Logout)
    }
]
const router = useRouter()

function handleDropdown(key:string) {
    switch (key) {
        case 'change-password':
            router.push({ name: "UserChangePassword"})
            break
        case 'logout':
            router.push({ name: "Login" })
            break
    }
}



</script>

<style lang="scss" scoped>
.header-container {
    box-shadow: 0 1px 2px rgb(0 21 41 / 8%)
}
</style>