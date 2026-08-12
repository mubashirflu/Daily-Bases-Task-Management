import { createRouter, createWebHistory } from "vue-router";

import Login from "../components/login.vue";
import Register from "../components/Register.vue";
import Dashboard from "../views/Dashboard.vue";

const router = createRouter({
  history: createWebHistory(),

  routes: [
    {
      path: "/",
      redirect: "/login",
    },

    {
      path: "/login",
      component: Login,
    },

    {
      path: "/register",
      component: Register,
    },

    {
      path: "/dashboard",
      component: Dashboard,
    },
  ],
});

export default router;