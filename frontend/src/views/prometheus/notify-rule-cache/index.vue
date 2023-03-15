<template>
  <vue-tree class="w-full h-full bg-#fff break-all" :dataset="dataset" :config="treeConfig">
    <template v-slot:node="{ node, collapsed }">
      <span class="bg-blue p-2px rd-t-5%  w-full " :style="{ border: collapsed ? '1px solid grey' : '', }"
        :class="{ 'text-center': node.key?.indexOf('=') === -1 }"> {{ node.key }}</span>
      <span v-if="node.value" @click="valueClick(node.value)" class="bg-#e2e2e2 w-full text-center"> {{ node.value }}
      </span>
    </template>
  </vue-tree>
</template>
<script lang="ts" setup>
import { ref, h } from 'vue'
import { useDialog, NDescriptions, NDescriptionsItem, NTag } from 'naive-ui';
// @ts-ignore
import VueTree from '@ssthouse/vue3-tree-chart'
import '@ssthouse/vue3-tree-chart/dist/vue3-tree-chart.css'
import { notifyRulesCacheTree, notifyRulesCacheList } from "@/api/notify"
const dataset = ref({})

const treeConfig = { nodeWidth: 120, nodeHeight: 60, levelHeight: 130 }

const cache = ref<any>()
notifyRulesCacheList()
  .then((response) => {
    cache.value = response.data
  })

notifyRulesCacheTree()
  .then(response => {
    dataset.value = response.data
  })
const dialog = useDialog()
function valueClick(value: string) {
  console.log(cache.value[value])
  let rule = cache.value[value]
  if (!rule) {
    return
  }
  dialog.create({
    title: value,
    positiveText: '确定',
    content: () => {
      // const renderMap={'Contacts':'通知目标',SenderSet}
      return h(NDescriptions, { column: 1, labelPlacement: 'left' }, {
        default: () => {
          return [
            h(NDescriptionsItem, { label: '通知目标' }, {
              default: () => {
                return rule.Contacts.map((item: string) => h(NTag, { size: 'small', class: 'mr-4px' }, { default: () => item }))
              }
            }),
            h(NDescriptionsItem, { label: '发送器' }, {
              default: () => {
                return rule.SenderSet.map((item: string) => h(NTag, { size: 'small', class: 'mr-4px' }, { default: () => item }))
              }
            }),
            h(NDescriptionsItem, { label: '发送器过滤' }, { default: () => JSON.stringify(rule.SenderFilter) }),
          ]
        }
      })
    }

  })
}
</script>


<style lang="scss">
.dom-container {
  .node-slot {
    flex-direction: column;
  }
}
</style>
