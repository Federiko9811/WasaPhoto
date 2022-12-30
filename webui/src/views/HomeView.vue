<script>
	import PhotoCard from "../components/PhotoCard.vue";
	export default {
		data() {
			return {
				myStream: [],
			}
		},
		methods: {
			getMyStream() {
				this.$axios.get(`/user/${this.$loggedUser.token}/photos/`).then((response) => {
					this.myStream = response.data;
				}).catch(
					(error) => {
						console.log(error);
					}
				);
			}
		},
		mounted() {
			if (this.$loggedUser.token === -1) {
				this.$router.push("/");
			}

			this.getMyStream();

		},
		components: {
			PhotoCard,
		}
	}
</script>

<template>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<div>
			<h1>Your Stream</h1>
		</div>

		<div class="d-flex flex-wrap justify-content-center gap-3">
			<PhotoCard v-for="photo in myStream" :key="photo.id" :photo="photo" />
		</div>

	</div>
</template>

<style>
</style>
