<script>
	export default {
		name: "SearchView",
		data() {
			return {
				searchResults: [],
				username: "",
				identifier: localStorage.getItem("identifier")
			}
		},
		methods: {
			searchUsers() {
				if (this.username.length >= 3 && this.username.length <= 16) {
					this.$axios.get(`/user/${this.identifier}/search/${this.username}`).then((response) => {
						this.searchResults = response.data;
						console.log(this.searchResults);
					}).catch((error) => {
						console.log(error);
					});
				}
			},
		},
	}
</script>

<template>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<div>
			<h1>Search User</h1>
		</div>
		<h2 class="bg-light border rounded">
			<input
				type="text"
				class="d-flex px-2 py-1"
				v-model="this.username"
				style="outline: none; border: none; background: none"
				placeholder="Search User"
				@input.prevent="searchUsers"
			>
		</h2>
		<div class="d-flex flex-column gap-2 w-50">
			<router-link
				v-for="user in searchResults"
				:key="user"
				:to="'/profile/' + user"
				class="text-decoration-none text-dark bg-light border px-2 py-1 rounded"
			>
				<h3>{{user}}</h3>
			</router-link>
		</div>
	</div>
</template>
