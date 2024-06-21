import { createRouter, createWebHistory } from 'vue-router'
import HomeView  from "@/views/HomeView.vue";
import ListView  from "@/views/ListView.vue";
import PlayView  from "@/views/PlayView.vue";
import SettingView   from "@/views/SettingView.vue";
import SearchView  from "@/views/SearchView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
      {
        path: '/list/:id',
        name: 'list',
        component: ListView,
      },
    {
      path:'/play/:id',
      name:'play',
      component:PlayView
    },
      {
        path:'/setting',
        name:'setting',
        component:SettingView
      },
    {
      path:'/search/:keyword',
      name:'search',
      component:SearchView
    }
  ]
})
router.beforeEach((to, from, next) => {
  localStorage.setItem('lastPath', from.path)
  localStorage.setItem('nowPath', to.path)
  next()
})
export default router
