<template>
    <n-divider title-placement="center" class="pt-12px">主题颜色</n-divider>
    <n-space vertical class="p-x-12px">
        <n-grid :cols="8" :x-gap="8" :y-gap="12">
            <n-grid-item v-for="color in themeColorList" :key="color" class="flex-x-center">
                <div class="flex-center w-20px h-20px rounded-2px shadow cursor-pointer"
                    :style="{ backgroundColor: color }" @click=" theme.setThemeColor(color)">
                    <OutlineCheck v-if="color === theme.themeColor" class='text-white' />
                </div>
            </n-grid-item>
        </n-grid>

        <n-color-picker class="mt-12px" :value="theme.themeColor" :show-alpha="false"
            @update-value="theme.setThemeColor" />
    </n-space>
</template>

<script setup lang="ts" >
import { NDivider, NGrid, NGridItem, NColorPicker, NSpace } from 'naive-ui'

import OutlineCheck from '@/assets/icons/OutlineCheck.vue'
import { useThemeStore } from '@/store'
const themeColorList = [
    "#1890ff",
    "#409EFF",
    "#2d8cf0",
    "#007AFF",
    "#5ac8fa",
    "#5856D6",
    "#536dfe",
    "#9c27b0",
    "#AF52DE",
    "#0096c7",
    "#00C1D4",
    "#34C759",
    "#43a047",
    "#7cb342",
    "#c0ca33",
    "#78DEC7",
    "#e53935",
    "#d81b60",
    "#f4511e",
    "#fb8c00",
    "#ffb300",
    "#fdd835",
    "#6d4c41",
    "#546e7a"
]

const theme = useThemeStore()

</script>
