<script>
	export default {
		name: "AddPhotoView",
		data() {
			return {
				errormsg: null,
				loading: false,
				photo: null
			}
		},
		methods: {
			addPhoto() {
				this.loading = true
				const id = localStorage.getItem("identifier")
				const formData = new FormData()
				formData.append("photo", this.$refs.photo.files[0])
				this.$axios.post(`/user/${id}/photos/`, formData).then((response) => {
					console.log(response.data)
				}).catch(
					(error) => {
						this.errormsg = error.response.data.message;
						console.log(error);
					}
				).finally(() => {
					this.loading = false
				});
			}
		},
	}
</script>

<template>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<div>
			<h1>Add Photo</h1>
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
