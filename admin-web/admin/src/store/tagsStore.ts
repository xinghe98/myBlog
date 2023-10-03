import { defineStore } from "pinia";

export const tagsStore = defineStore("tags", {
	persist: false,
	state: () => ({
		isShow: false,
		windowName: "",
		rowID: 0,
	}),
	actions: {
		addTag(windowName: string) {
			this.isShow = true;
			this.windowName = windowName;
		},
		editTag(windowName: string, rowID: number) {
			this.isShow = true;
			this.windowName = windowName;
			this.rowID = rowID;
		},
	},
});
