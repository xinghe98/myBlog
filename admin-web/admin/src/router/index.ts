import { createRouter, createWebHashHistory } from "vue-router";
const router = createRouter({
  // 打包时启用
  // history: createWebHashHistory(process.env.BASE_URL),
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/views/login.vue"),
    },
  ],
});

export default router;
