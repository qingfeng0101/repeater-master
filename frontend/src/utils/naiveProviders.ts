import type { NotificationProviderInst, LoadingBarProviderInst, MessageProviderInst, DialogProviderInst } from 'naive-ui'

let Notification: NotificationProviderInst | undefined
let Bialog: DialogProviderInst | undefined
let Message: MessageProviderInst | undefined
let LoadingBar: LoadingBarProviderInst | undefined

// 方便在 ts 中直接调用
export function registerProviders(option: {
  notification: NotificationProviderInst
  dialog: DialogProviderInst
  message: MessageProviderInst
  loadingBar: LoadingBarProviderInst
}): void {
  Notification = option.notification
  Bialog = option.dialog
  Message = option.message
  LoadingBar = option.loadingBar
}

export { Notification, Bialog, Message, LoadingBar }
