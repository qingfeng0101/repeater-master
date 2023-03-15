import 'vue-router'
// meta: {
//       title - 标题
//       permissions - 权限列表
//       icon - 图标
//       hidden - 是否隐藏菜单
//       flat - 是否扁平菜单
//   },
declare module 'vue-router' {
  interface RouteMeta {

    requiresAuth?: boolean //依赖认证
    flatChildrenInMenu?: boolean //扁平化子菜单
    icon?: string  //图标
    hideInMenu?: boolean //不在菜单中显示
    locale?: string  //标题的国际化
    roles?: string[]  //可访问的角色
  }
}
