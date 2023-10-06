import { defineStore } from "pinia";
import request from "@/util/request";
import { ElMessage } from "element-plus";

export const articleStore = defineStore("article", {
	persist: false,
	state: () => ({
		ID: 0,
		title: "",
		content: "# Hello World",
		status: 0,
		tags: [],
	}),
	actions: {
		async updateArticle(id: number) {
			try {
				const res = await request.put(`/article/${id}`, {
					title: this.title,
					content: this.content,
					status: this.status,
					tags: this.tags,
				});
				ElMessage.success(res.data.msg);
			} catch (e: any) {
				if (e.response.data.data !== null) {
					const keys = Object.keys(e.response.data.data);
					ElMessage.error(e.response.data.data[keys[0]][0]);
				} else {
					ElMessage.error(e.response.data.msg);
				}
			}
		},

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
					const keys = Object.keys(e.response.data.data);
					ElMessage.error(e.response.data.data[keys[0]][0]);
				} else {
					ElMessage.error(e.response.data.msg);
				}
			}
		},
	},
});
