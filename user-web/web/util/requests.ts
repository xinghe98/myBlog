import { useFetch } from "nuxt/app";
interface Res<T> {
	code: number;
	data: responseData<T>;
	msg: string;
}
interface responseData<T> {
	articles: T[];
	current_page_size: number;
	limit: number;
	page: number;
	total: number;
}

const request = async <K>(path: string) => {
	return await useFetch<Res<K>>(() => `http://127.0.0.1:3001${path}`);
};

export default request;
