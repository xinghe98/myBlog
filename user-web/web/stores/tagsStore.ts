import { defineStore } from "pinia";

export const useTagsStore = defineStore("tagsStore", {
	state: () => ({
		tags: "",
	}),
});
