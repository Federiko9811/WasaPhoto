<script>
export default {
	data() {

	},
	methods: {

	},
	mounted() {
		if (!localStorage.getItem("identifier")) {
			this.$router.push("/");
		}
	}
}
</script>

<template>
	<div>
		<h1>Home</h1>

		<RouterLink to='/profile/Federiko9811'>
			Federiko9811
		</RouterLink>

	</div>
</template>

<style>
</style>
