import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const credential: RouteRecordRaw = {
    path: '/credential',
    name: 'Credential',
    component: DefaultLayout,
    meta: {
        locale: 'Credential',
        icon: 'Credentials'
    },
    children: [
        {
            path: 'stmp',
            name: 'CredentialStmp',
            component: () => import('@/views/credential/stmp/index.vue'),
            meta: {
                locale: 'Stmp',
                icon: 'EmailOutline'
            },
        },
        {
            path: 'stmp/add',
            name: 'CredentialStmpAdd',
            component: () => import('@/views/credential/stmp/Add.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'wecom',
            name: 'CredentialWecom',
            component: () => import('@/views/credential/wecom/index.vue'),
            meta: {
                locale: 'Wecom',
                icon: 'WechatLine'
            },
        },
        {
            path: 'wecom/add',
            name: 'CredentialWecomAdd',
            component: () => import('@/views/credential/wecom/Add.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'dingtalk',
            name: 'CredentialDingtalk',
            component: () => import('@/views/credential/dingtalk/index.vue'),
            meta: {
                locale: 'Dingtalk',
                icon: 'DingtalkLine'
            },
        },
        {
            path: 'dingtalk/add',
            name: 'CredentialDingtalkAdd',
            component: () => import('@/views/credential/dingtalk/Add.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'lark',
            name: 'CredentialLark',
            component: () => import('@/views/credential/lark/index.vue'),
            meta: {
                locale: 'Lark',
                icon: 'LarkLine'
            },
        },
        {
            path: 'lark/add',
            name: 'CredentialLarkAdd',
            component: () => import('@/views/credential/lark/Add.vue'),
            meta: {
                hideInMenu: true
            },
        },

    ]
}