<script setup>
import { RouterView } from 'vue-router'
</script>
<script>
	export default {
		data(){
			return {
				username: localStorage.getItem("username"),
			}
		},
		methods: {
			async handleLogout() {
				localStorage.removeItem("identifier");
				localStorage.removeItem("username");
				this.$router.push("/");
			}
		},
	}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse" v-if="$route.name !== 'Login'">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Dashboard</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/profile/' + this.username" class="nav-link">
								Profile Page
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to='/add-photo' class="nav-link">
								Add Photo
							</RouterLink>
						</li>
						<li class="nav-item">
							<div class="nav-link" @click="handleLogout">
								Logout
							</div>
						</li>

					</ul>
				</div>
			</nav>
			<main class="col-md-9 ms-sm-auto col-lg-12" v-if="$route.name === 'Login'">
				<RouterView />
			</main>
			<main class="col-md-9 ms-sm-auto col-lg-10" v-if="$route.name !== 'Login'">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
