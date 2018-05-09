import Vue from 'vue'
import Router from 'vue-router'

import Layout from '@/components/Layout'
import Resources from '@/components/content/Resources'
import Processes from '@/components/content/Processes'
import FS from '@/components/content/FS.vue'
import Setting from '@/components/content/Setting.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Layout',
      component: Layout,
      redirect: Resources,
      children: [
        {
          path: 'resources',
          name: 'Resources',
          component: Resources,
        },
        {
          path: 'processes',
          name: 'Processes',
          component: Processes,
        },
        {
          path: 'fs',
          name: 'FS',
          component: FS,
        },
        {
          path: 'setting',
          name: 'Setting',
          component: Setting,
        },
      ],
    }
  ]
})
