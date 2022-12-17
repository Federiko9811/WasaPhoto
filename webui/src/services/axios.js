import axios from "axios";

const instance = axios.create({
	baseURL: "http://localhost:3000",
	timeout: 1000 * 5
});
instance.defaults.headers.common["Authorization"] = "Bearer " + localStorage.getItem("identifier");
export default instance;
