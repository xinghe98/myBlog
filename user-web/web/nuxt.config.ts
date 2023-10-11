// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: ["@element-plus/nuxt"],
	devtools: { enabled: true },
	vite: {
		vue: {
			customElement: true,
		},
	},
});
