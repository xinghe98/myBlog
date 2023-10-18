// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: ["@element-plus/nuxt", "@pinia/nuxt"],
	devtools: { enabled: false },
	vite: {
		vue: {
			customElement: true,
		},
	},
	runtimeConfig: {
		public: {
			baseURL: "http://localhost:3001",
		},
	},
});
