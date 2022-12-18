import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from "../views/ProfileView.vue";
import AddPhotoView from "../views/AddPhotoView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView, name: "Login"},
		{path: '/home', component: HomeView},
		{path: '/profile/:username', component: ProfileView},
		{path: '/add-photo', component: AddPhotoView},
	]
})

export default router
