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
		img: "", // 普通封面
		headimg: "", // 头条封面
		isheadlines: false,
		isSuccess: false, //是否提交成功
	}),
	actions: {
		async updateArticle(id: number) {
			try {
				const res = await request.put(`/article/${id}`, {
					title: this.title,
					content: this.content,
					status: this.status,
					tags: this.tags,
					image: this.img,
					headimg: this.headimg,
					isheadlines: this.isheadlines,
				});
				ElMessage.success(res.data.msg);
				this.isSuccess = true;
			} catch (e: any) {
				if (e.response.data.data !== null) {
					const keys = Object.keys(e.response.data.data);
					ElMessage.error(e.response.data.data[keys[0]][0]);
				} else {
					ElMessage.error(e.response.data.msg);
				}
				this.isSuccess = false;
			}
		},

		async createArticle() {
			try {
				const res = await request.post("/article/create", {
					title: this.title,
					content: this.content,
					status: this.status,
					tags: this.tags,
					image: this.img,
					headimg: this.headimg,
					isheadlines: this.isheadlines,
				});
				ElMessage.success(res.data.msg);
				this.isSuccess = true;
			} catch (e: any) {
				if (e.response.data.data !== null) {
					const keys = Object.keys(e.response.data.data);
					ElMessage.error(e.response.data.data[keys[0]][0]);
				} else {
					ElMessage.error(e.response.data.msg);
				}
				this.isSuccess = false;
			}
		},
	},
});
