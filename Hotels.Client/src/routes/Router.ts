import { createRouter, createWebHistory, } from 'vue-router';
import index from '../pages/index.vue';
import Home from '../pages/Home/home.vue';
import Dashboard from '../pages/Home/Dashboard/dashboard.vue';
import Hotels from '../pages/Home/Hotels/Hotel.vue';
import Rooms from '../pages/Home/Rooms/Room.vue'
import About from '../pages/About/about.vue';
import Login from '../pages/Login/Login.vue';
import NotFound from '../pages/NotFound/NotFound.vue';



export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/home/',
      name: 'Home',
      component: Home,
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: Dashboard
        },
        {
          path: 'hotels',
          name: 'Hotels',
          component: Hotels
        },
        {
          path: 'rooms',
          name: 'Rooms',
          component: Rooms
        }
      ]
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/about',
      name: 'About',
      component: About
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: NotFound
    }
  ]
});


