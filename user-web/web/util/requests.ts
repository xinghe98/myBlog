import axios from "axios";

const request = axios.create({
	baseURL: "http://127.0.0.1:3001",
	// baseURL: "http://192.168.31.9:3001",
});

export default request;
