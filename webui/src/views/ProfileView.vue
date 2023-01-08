<script>
	export default {
		emits: ['photoDeleted'],
		name: "ProfileView",
		data() {
			return {
				profile: {
					username: "",
					photos: [],
					numberOfFollowers: 0,
					numberOfFollowing: 0,
					numberOfPhotos: 0,
					isBanned: false,
					isFollowed: false,
					isOwner: false,
				},
				tempIsBanned: false,
				tempIsFollowed: false,
				tempUsername: "",
				showEditProfile: false,
				id: this.$loggedUser.token,
				forbidden_error: false,
			}
		},
		methods: {
			getProfile() {
				const pathUsername = this.$route.params.username

				this.$axios.get(`/user/${this.id}/profile-page/${pathUsername}`).then((response) => {
					this.profile = response.data;
					this.tempIsBanned = response.data.isBanned;
					this.tempIsFollowed = response.data.isFollowed;
					this.tempUsername = response.data.username;
				}).catch(
					(error) => {
						if (error.response.status === 403) {
							this.forbidden_error = true;
						}
					}
				);
			},
			followUser() {
				this.$axios.put(`/user/${this.id}/follow/${this.profile.username}`,).then(() => {
					this.tempIsFollowed = true;
				}).catch((error) => {
					console.log(error);
				});
			},
			unfollowUser() {
				this.$axios.delete(`/user/${this.id}/follow/${this.profile.username}`,).then(() => {
					this.tempIsFollowed = false;
				}).catch((error) => {
					console.log(error);
				});
			},
			banUser() {
				this.$axios.put(`/user/${this.id}/ban/${this.profile.username}`,).then(() => {
					this.tempIsBanned = true;
				}).catch((error) => {
					console.log(error);
				});
			},
			unbanUser() {
				this.$axios.delete(`/user/${this.id}/ban/${this.profile.username}`,).then(() => {
					this.tempIsBanned = false;
				}).catch((error) => {
					console.log(error);
				});
			},
			editProfile() {
				this.$axios.put(`/user/${this.id}/update-username`, {name: this.tempUsername}).then(() => {
					this.showEditProfile = false;
					this.$loggedUser.username = this.tempUsername;
					this.$router.push(`/profile/${this.tempUsername}`);
				}).catch((error) => {
					console.log(error);
				});
			},
		},
		mounted() {
			if (this.id === -1) {
				this.$router.push("/");
			}
			this.getProfile();
		}
	}
</script>

<template>
	<!--	Profile page using bootstrap-->
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3" v-if="this.forbidden_error">
		<div class="alert alert-danger" role="alert">
			Sei bannato da questo utente
		</div>
	</div>

	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3" v-if="!this.forbidden_error">
		<div v-if="!showEditProfile">
			<h1>{{this.tempUsername}}</h1>
		</div>
		<div v-if="showEditProfile && profile.isOwner" class="d-flex align-items-center align-items-center bg-light p-3 rounded">
			<h2>
				<input type="text" v-model="this.tempUsername" style="outline: none; border: none; background: none">
			</h2>
			<div role="button" @click.prevent="editProfile">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					width="3em"
					height="3em"
					class="text-primary"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10"
					/>
				</svg>
			</div>
		</div>
		<div class="d-flex gap-5">
			<h5>Photos: {{profile.numberOfPhotos}}</h5>
			<h5>Followers: {{profile.numberOfFollowers}}</h5>
			<h5>Following: {{profile.numberOfFollowing}}</h5>
		</div>
		<div>
			<div class="d-flex gap-5">
				<button
					class="btn btn-primary"
					v-if="!this.tempIsFollowed && !profile.isOwner"
					@click.prevent="followUser"
				>
					Follow
				</button>
				<button
					class="btn btn-primary"
					v-if="this.tempIsFollowed && !profile.isOwner"
					@click.prevent="unfollowUser"
				>
					Unfollow
				</button>
				<button
					class="btn btn-danger"
					v-if="!this.tempIsBanned && !profile.isOwner"
					@click.prevent="banUser"
				>
					Ban
				</button>
				<button
					class="btn btn-danger"
					v-if="this.tempIsBanned && !profile.isOwner"
					@click.prevent="unbanUser"
				>
					Unban
				</button>
				<button class="btn btn-primary" v-if="profile.isOwner" @click.prevent="() => this.showEditProfile=!this.showEditProfile">
					Edit Profile
				</button>
			</div>
		</div>
		<div class="d-flex flex-wrap justify-content-center gap-3">
			<PhotoCard v-for="photo in profile.photos" :key="photo.id" :photo="photo" v-on:photoDeleted="this.getProfile"/>
		</div>

	</div>
</template>
