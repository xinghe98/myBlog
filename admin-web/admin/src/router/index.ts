import { adminStore } from "@/store/adminStore";
import { articleStore } from "@/store/articleStore";
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
					// 已发布
					path: "readall",
					name: "readall",
					component: () => import("@/views/readall.vue"),
				},
				{
					// 草稿箱
					path: "readsaved",
					name: "readsaved",
					component: () => import("@/views/readsaved.vue"),
				},
				{
					// 新建
					path: "create",
					name: "create",
					component: () => import("@/views/create.vue"),
				},
				{
					// 编辑文章（复用新建文章组件）
					path: "edit/:id",
					name: "edit",
					component: () => import("@/views/create.vue"),
				},
				{
					// 标签相关操作
					path: "tags",
					name: "tags",
					component: () => import("@/views/tagOp.vue"),
				},
			],
		},
	],
});

// 路由守卫
router.beforeEach(async (to, form) => {
	const admin = adminStore();
	if (
		// 检查用户是否已登录
		!admin.isAuthenticated &&
		// ❗️ 避免无限重定向
		to.name !== "home"
	) {
		// 将用户重定向到登录页面
		return { name: "home" };
	}
	const article = articleStore();
	// 当点击的是新建文章时，清空文章数据
	if (to.name === "create") {
		article.$reset();
		// 刷新当前页面
	}
	if (form.name === "create" || form.name === "edit") {
		// 当离开新建文章页面时，弹出确认框
		let res = confirm("离开此页面将会丢失已编辑的内容");
		if (res) {
			// 确认离开
			article.$reset();
			return true;
		}
		// 留在当前页面，销毁弹出框
		return false;
	}
});
// 路由守卫：离开特定页面时，弹出确认框
export default router;
