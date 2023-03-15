import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const contact: RouteRecordRaw = {
    path: '/contact',
    name: 'Contact',
    component: DefaultLayout,
    meta: {
        flatChildrenInMenu: true
    },
    children: [
        {
            path: 'list',
            name: 'ContactList',
            component: () => import('@/views/contact/index.vue'),
            meta: {
                locale: 'Contact',
                icon: 'ContactOutline'
            },
        },
        {
            path: 'add',
            name: 'ContactAdd',
            component: () => import('@/views/contact/Action.vue'),
            meta: {
                locale: 'ContactAdd',
                hideInMenu: true,
            },
        },
        {
            path: ':id/edit',
            name: 'ContactEdit',
            props: true,
            component: () => import('@/views/contact/Action.vue'),
            meta: {
                locale: 'ContactEdit',
                hideInMenu: true,
            },
        },
    ]
}