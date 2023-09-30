import { adminStore } from "@/store/adminStore";
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
		{
			path: "/admin",
			name: "allarticle",
			component: () => import("@/views/adminhome.vue"),
			redirect: "/admin/readall",
			children: [
				{
					path: "readall",
					name: "readall",
					component: () => import("@/views/readall.vue"),
				},
				{
					path: "create",
					name: "create",
					component: () => import("@/views/create.vue"),
				},
				{
					path: "tags",
					name: "tags",
					component: () => import("@/views/tagOp.vue"),
				},
			],
		},
	],
});

// 路由守卫
router.beforeEach(async (to) => {
	const admin = adminStore();
	if (
		// 检查用户是否已登录
		!admin.isAuthenticated &&
		// ❗️ 避免无限重定向
		to.name !== "home"
	) {
		console.log(admin.isAuthenticated);
		// 将用户重定向到登录页面
		return { name: "home" };
	}
});
export default router;
