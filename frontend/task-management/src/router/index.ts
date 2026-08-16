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
      meta:{
        requiresAuth:true,
      }
    },
  
  ],
  
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");

  // Dashboard protected hai
  if (to.meta.requiresAuth && !token) {
    next("/login");
    return;
  }

  // Already logged in hai aur login page kholne ki koshish kar raha hai
  if (to.path === "/login" && token) {
    next("/dashboard");
    return;
  }

  next();
});

export default router;