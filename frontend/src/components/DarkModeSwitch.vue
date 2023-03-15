<template>
  <div class="text-20px cursor-pointer" @click="handleSwitch">
    <n-icon>
      <DarkModeOutline v-if="darkMode" />
      <LightModeOutline v-else />
    </n-icon>
  </div>
</template>

<script setup  lang="ts" >
import { computed } from 'vue'
import { NIcon } from 'naive-ui'
import LightModeOutline from '../assets/icons/LightModeOutline.vue'
import DarkModeOutline from '../assets/icons/DarkModeOutline.vue'



interface Props {
  dark?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  dark: false
})

interface Emits {
  (e: 'update:dark', darkMode: boolean): void
}

const emit = defineEmits<Emits>()


const darkMode = computed({
  get() {
    return props.dark
  },
  set(newValue: boolean) {
    emit('update:dark', newValue)
  }
})

function handleSwitch() {
  darkMode.value = !darkMode.value
}
</script>


