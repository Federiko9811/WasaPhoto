

<script>
	import PhotoCard from "../components/PhotoCard.vue";

	export default {
		name: "ProfileView",
		data() {
			return {
				errormsg: null,
				loading: false,
				profile: {
					username: "",
					photos: [],
					numberOfFollowers: 0,
					numberOfFollowing: 0,
					numberOfPhotos: 0,
				},
			}
		},
		methods: {
			getProfile() {
				this.loading = true

				//get username from the path
				const x = this.$route.params.username
				const id = localStorage.getItem("identifier")

				this.$axios.get(`/user/${id}/profile-page/${x}`).then((response) => {
					this.profile = response.data;
					this.loading = false;
				}).catch(
					(error) => {
						this.errormsg = error.response.data.message;
						console.log(error);
					}
				);
			}
		},
		mounted() {
			if (!localStorage.getItem("identifier")) {
				this.$router.push("/");
			}
			this.getProfile();
		},
		components: {
			PhotoCard,
		}
	}
</script>

<template>
<!--	Profile page using bootstrap-->
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<div>
			<h1>{{profile.username}}</h1>
		</div>
		<div class="d-flex gap-5">
			<h5>Photos: {{profile.numberOfPhotos}}</h5>
			<h5>Followers: {{profile.numberOfFollowers}}</h5>
			<h5>Following: {{profile.numberOfFollowing}}</h5>
		</div>
<!--		List all photos using photocard-->
		<div class="d-flex flex-wrap gap-3">
			<PhotoCard v-for="photo in profile.photos" :key="photo.id" :photo="photo"/>
		</div>

	</div>
</template>
