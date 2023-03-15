import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const sender: RouteRecordRaw = {
    path: '/sender',
    name: 'Sender',
    component: DefaultLayout,
    meta: {
        flatChildrenInMenu: true,
        locale: 'Sender',
        icon: 'Prometheus'
    },
    children: [
        {
            path: 'sets',
            name: 'SenderSets',
            component: () => import('@/views/sender/sets/index.vue'),
            meta: {
                locale: 'SenderSets',
                icon: 'DataSwitching'
            },
        },
        {
            path: 'add',
            name: 'SenderSetsAdd',
            component: () => import('@/views/sender/sets/Action.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: ':id/edit',
            name: 'SenderSetsEdit',
            component: () => import('@/views/sender/sets/Action.vue'),
            props: true,
            meta: {
                hideInMenu: true
            },
        },
    ]
}