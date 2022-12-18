

<script>
	export default {
		name: "PhotoCard",
		data() {
			return {
				img: "",
				temp_photo: {
					id: this.$props.photo.id,
					numberOfLikes: this.$props.photo.numberOfLikes,
					numberOfComments: this.$props.photo.numberOfComments,
					isLiked: this.$props.photo.isLiked,
				},
				list_of_comments: [],
				show_comments: false,
				comment: "",
				username_req_user: localStorage.getItem("username"),
			}
		},
 		props: {
			photo: {
				type: Object,
				required: true,
			},
		},
		methods: {
			getPhoto() {
				const identifier = localStorage.getItem("identifier")
				this.$axios.get(`/user/${identifier}/photos/${this.$props.photo.id}/`).then((response) => {
					this.img = response.data;
				}).catch(
					(error) => {
						console.log(error);
					}
				);
			},
			addLike() {
				const identifier = localStorage.getItem("identifier")
				this.$axios.put(`/user/${this.$props.photo.owner}/photos/${this.$props.photo.id}/likes/${identifier}`).then((response) => {
					this.temp_photo.numberOfLikes += 1;
					this.temp_photo.isLiked = true;
				}).catch(
					(error) => {
						console.log(error);
					}
				)
			},
			removeLike() {
				const identifier = localStorage.getItem("identifier")
				this.$axios.delete(`/user/${this.$props.photo.owner}/photos/${this.$props.photo.id}/likes/${identifier}`).then((response) => {
					this.temp_photo.numberOfLikes -= 1;
					this.temp_photo.isLiked = false;
				}).catch(
					(error) => {
						console.log(error);
					}
				)
			},
			addComment() {
				const identifier = localStorage.getItem("identifier")
				this.$axios.post(`/user/${identifier}/photos/${this.$props.photo.id}/comments/`, {
					comment: this.comment,
				}).then((r) => {
					this.temp_photo.numberOfComments += 1;
					const now = new Date().toISOString();
					this.list_of_comments.unshift({
						id: r.data.comment_id,
						content: this.comment,
						created_at: now,
						owner: this.username_req_user
					})
					this.comment = "";
				}).catch(
					(error) => {
						console.log(error);
					}
				)
			},
			getComments() {
				const identifier = localStorage.getItem("identifier")
				this.$axios.get(`/user/${identifier}/photos/${this.$props.photo.id}/comments/`).then((response) => {
					this.list_of_comments = response.data;
				}).catch(
					(error) => {
						console.log(error);
					}
				);
			},
			deleteComment(commentId) {
				const identifier = localStorage.getItem("identifier")
				this.$axios.delete(`/user/${identifier}/photos/${this.$props.photo.id}/comments/${commentId}`).then(() => {
					this.temp_photo.numberOfComments -= 1;
					this.list_of_comments = this.list_of_comments.filter((comment) => comment.id !== commentId);
				}).catch(
					(error) => {
						console.log(error);
					}
				);
			},
		},
		mounted() {
			this.getPhoto();
			this.getComments();
		},
	}
</script>

<template>
	<div class="p-3 card gap-3">
		<img
			:src="img"
			:alt='"ID: " + photo.id'
			class="img-thumbnail"
			style="width: 300px; height: 300px;"
		>
		<div class="fst-italic small">
			{{photo.createdAt}}
		</div>
		<div class="d-flex justify-content-between">
			<div class="d-flex align-items-center gap-2">
				<svg
					v-if="!temp_photo.isLiked"
					xmlns="http://www.w3.org/2000/svg"
					@click.prevent="addLike"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					width="24"
					height="24"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M6.633 10.5c.806 0 1.533-.446 2.031-1.08a9.041 9.041 0 012.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 00.322-1.672V3a.75.75 0 01.75-.75A2.25 2.25 0 0116.5 4.5c0 1.152-.26 2.243-.723 3.218-.266.558.107 1.282.725 1.282h3.126c1.026 0 1.945.694 2.054 1.715.045.422.068.85.068 1.285a11.95 11.95 0 01-2.649 7.521c-.388.482-.987.729-1.605.729H13.48c-.483 0-.964-.078-1.423-.23l-3.114-1.04a4.501 4.501 0 00-1.423-.23H5.904M14.25 9h2.25M5.904 18.75c.083.205.173.405.27.602.197.4-.078.898-.523.898h-.908c-.889 0-1.713-.518-1.972-1.368a12 12 0 01-.521-3.507c0-1.553.295-3.036.831-4.398C3.387 10.203 4.167 9.75 5 9.75h1.053c.472 0 .745.556.5.96a8.958 8.958 0 00-1.302 4.665c0 1.194.232 2.333.654 3.375z" />
				</svg>
				<svg
					v-if="temp_photo.isLiked"
					xmlns="http://www.w3.org/2000/svg"
					@click.prevent="removeLike"
					fill="true"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					width="24"
					height="24"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M6.633 10.5c.806 0 1.533-.446 2.031-1.08a9.041 9.041 0 012.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 00.322-1.672V3a.75.75 0 01.75-.75A2.25 2.25 0 0116.5 4.5c0 1.152-.26 2.243-.723 3.218-.266.558.107 1.282.725 1.282h3.126c1.026 0 1.945.694 2.054 1.715.045.422.068.85.068 1.285a11.95 11.95 0 01-2.649 7.521c-.388.482-.987.729-1.605.729H13.48c-.483 0-.964-.078-1.423-.23l-3.114-1.04a4.501 4.501 0 00-1.423-.23H5.904M14.25 9h2.25M5.904 18.75c.083.205.173.405.27.602.197.4-.078.898-.523.898h-.908c-.889 0-1.713-.518-1.972-1.368a12 12 0 01-.521-3.507c0-1.553.295-3.036.831-4.398C3.387 10.203 4.167 9.75 5 9.75h1.053c.472 0 .745.556.5.96a8.958 8.958 0 00-1.302 4.665c0 1.194.232 2.333.654 3.375z" />
				</svg>
				<div>
					{{temp_photo.numberOfLikes}}
				</div>
			</div>
			<div class="d-flex align-items-center gap-2">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					width="24"
					height="24"
					@click.prevent="() => this.show_comments = !this.show_comments"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 01.865-.501 48.172 48.172 0 003.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0012 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018z" />
				</svg>
				<div>
					{{temp_photo.numberOfComments}}
				</div>
			</div>
		</div>
		<!--Lista commenti-->
		<div v-if="show_comments" class="d-flex flex-column gap-2 align-items-center justify-content-center">
			<div class="d-flex flex-column gap-2 w-100">
				<div v-for="comment in list_of_comments" class="d-flex flex-column gap-2">
					<div class="d-flex flex-column gap-2 bg-light p-2 rounded border">
						<div class="d-flex justify-content-between">
							<div class="fw-bold">
								{{comment.owner}}
							</div>
							<div class="small">
								{{comment.created_at}}
							</div>

						</div>
						<div>
							{{comment.content}}
						</div>
						<div
							v-if="comment.owner === this.username_req_user"
							class="small fst-italic text-danger"
							role="button"
							@click.prevent="deleteComment(comment.id)"
						>
							Elimina
						</div>
					</div>
				</div>
			</div>
			<div @click.prevent="() => this.show_comments=!this.show_comments" role="button" class="text-secondary">
				Nascondi Commenti
			</div>
		</div>
		<!--Form commento-->
		<form class="d-flex align-items-center justify-content-between gap-2" @submit.prevent="addComment">
			<input
				type="text"
				placeholder="Commenta..."
				class="d-flex w-100 px-2 py-1 rounded bg-light"
				style="outline: none; border: solid 1px #ced4da;"
				v-model="this.comment"
				required
			>
			<button
				class="d-flex align-items-center justify-content-center"
				style="background: none; border: none"
				type="submit"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="text-primary"
					width="24"
					height="24"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M6 12L3.269 3.126A59.768 59.768 0 0121.485 12 59.77 59.77 0 013.27 20.876L5.999 12zm0 0h7.5"
					/>
				</svg>
			</button>

		</form>
	</div>
</template>
