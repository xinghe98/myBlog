import axios, { AxiosResponse } from "axios";
import { ElMessage } from "element-plus";

const request = axios.create({
	 baseURL: "/api",
});

request.interceptors.request.use(
	(config: any) => {
		if (localStorage.getItem("adminToken")) {
			config.headers.Authorization = "Bearer " + localStorage.getItem("adminToken");
		}
		return config;
	},
	(err: any) => {
		return Promise.reject(err);
	},
);
// 响应拦截器
request.interceptors.response.use(
	(res: AxiosResponse<any, any> | Promise<AxiosResponse<any, any>>) => {
		// 在请求成功后的数据处理
		return res;
	},
	(err: { response: { status: number } }) => {
		// 在响应错误的时候的逻辑处理
		console.log(err.response);
		if (err.response.status === 401) {
			window.location.href = "/";
			ElMessage({
				message: "请先登录",
				type: "error",
			});
			localStorage.removeItem("adminToken");
		}

		return Promise.reject(err);
	},
);
export default request;
