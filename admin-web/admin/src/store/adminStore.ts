import { defineStore } from "pinia";
import request from "@/util/request";
import { ElMessage } from "element-plus";

export const adminStore = defineStore("admin", {
	persist: true,
	state: () => ({
		isAuthenticated: false,
		username: "",
	}),
	actions: {
		async login(formdata: { username: string; password: string }) {
			try {
				const res = await request.post("/admin/login", formdata);
				localStorage.setItem("adminToken", res.data.data.JwtToen);
				ElMessage.success("登录成功");
				this.isAuthenticated = true;
				this.username = formdata.username;
			} catch (e: any) {
				this.isAuthenticated = false;
				ElMessage.error(e.response.data.msg);
			}
		},
	},
});
