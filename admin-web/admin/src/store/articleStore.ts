import { defineStore } from "pinia";
import request from "@/util/request";
import { ElMessage } from "element-plus";

export const articleStore = defineStore("article", {
	persist: false,
	state: () => ({
		title: "",
		content: "",
		status: 0,
		tags: [],
	}),
	actions: {
		async createArticle() {
			try {
				const res = await request.post("/article/create", {
					title: this.title,
					content: this.content,
					status: this.status,
					tags: this.tags,
				});
				ElMessage.success(res.data.msg);
			} catch (e: any) {
				if (e.response.data.data !== null) {
					// FIX: 未知原因，data为null时，无法正确提示
					console.log(e.response.data.data);
					ElMessage.error(e.response.data.data[0]);
				} else {
					ElMessage.error(e.response.data.msg);
				}
			}
		},
	},
});
