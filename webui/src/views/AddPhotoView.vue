<script>
	export default {
		name: "AddPhotoView",
		data() {
			return {
				photo: null,
				success: false,
				error: false,
			}
		},
		methods: {
			addPhoto() {
				const currentPhoto = this.$refs.photo.files[0]
				if (currentPhoto) {
					this.$axios.post(`/user/${this.$loggedUser.token}/photos/`, currentPhoto).then(() => {
						this.success = true
						this.error = false
					}).catch(
						() => {
							this.success = false
							this.error = true
							this.photo = null
						}
					)
				} else {
					this.success = false
					this.error = true
				}


			},
			closeAlert() {
				this.success = false
				this.error = false
			}
		},
		mounted() {
			if (this.$loggedUser.token === -1) {
				this.$router.push("/")
			}
		}
	}
</script>

<template>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<div>
			<h1>Add Photo</h1>
		</div>
		<div
			class="alert alert-success w-50 d-flex justify-content-between"
			role="alert"
			v-if="this.success"
		>
			<div>
				Foto Caricata correttamente
			</div>
			<div role="button" @click.prevent="closeAlert">
				X
			</div>
		</div>
		<div
			class="alert alert-danger w-50 d-flex justify-content-between"
			role="alert"
			v-if="this.error"
		>
			<div>
				Errore di caricamento
			</div>
			<div role="button" @click.prevent="closeAlert">
				X
			</div>
		</div>
		<form class="d-flex flex-column gap-3 card w-50" @submit.prevent="addPhoto">
			<label for="photo" class="card-header">Select a Photo</label>
			<div class="d-flex flex-column gap-3 p-3">
				<input type="file" ref="photo" accept="image/png">
				<button type="submit" class="btn btn-primary">Add Photo</button>
			</div>
		</form>
	</div>
</template>
