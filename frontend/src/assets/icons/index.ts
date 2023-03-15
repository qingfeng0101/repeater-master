import type { Component } from 'vue'

const modules = import.meta.glob<{ default: Component }>('./*.vue', { eager: true })


const icons: Record<string, Component> = {}

Object.keys(modules).forEach((key) => {
      let name = modules[key].default.name as string
      icons[name] = modules[key].default
})


export default icons