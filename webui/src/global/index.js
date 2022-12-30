const loggedUser = {
	token: -1,
	username: "",

	setLoggedUser(token, username) {
		this.token = token;
		this.username = username;
	},

	logout() {
		this.token = -1;
		this.username = "";
	}
}

export default loggedUser;
