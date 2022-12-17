<script>
	export default {
		name: "Login",
		data() {
			return {
				name: ""
			}
		},
		methods: {
			async handleLogin() {
				const response = await this.$axios.post("/session", {name: this.name});
				localStorage.setItem("identifier", response.data.identifier);
				localStorage.setItem("username", this.name);
				this.$router.push("/home");
			}
		}
	}
</script>


<template>
	<div class="container d-flex min-vh-100 align-items-center justify-content-center">
		<div class="card w-50">
			<div class="card-header">
				<h3 class="card-title mt-2">Login</h3>
			</div>
			<div class="card-body">
				<form @submit.prevent="handleLogin" class="d-flex flex-column gap-4">
					<div class="form-group">
						<label for="username">Username</label>
						<input
							type="text"
							v-model="this.name"
							class="form-control"
							placeholder="Enter username"
							required
							minlength="3"
							maxlength="16"
						>
					</div>
					<button type="submit" class="btn btn-primary">Login</button>
				</form>
			</div>
		</div>
	</div>
</template>
