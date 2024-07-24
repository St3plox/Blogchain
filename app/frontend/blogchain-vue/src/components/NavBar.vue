<template>
    <div>
        <b-navbar toggleable="lg" type="dark" variant="dark">
            <b-navbar-brand href="/">Blogchain</b-navbar-brand>

            <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

            <b-collapse id="nav-collapse" is-nav>
                <b-navbar-nav>
                    <b-nav-item v-if="token" href="/posts">My Posts</b-nav-item>
                </b-navbar-nav>

                <!-- Right aligned nav items -->
                <b-navbar-nav class="ml-auto">
                    <b-nav-form>
                        <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
                        <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
                    </b-nav-form>

                    <b-nav-item-dropdown text="Lang" right>
                        <b-dropdown-item href="#">EN</b-dropdown-item>
                        <b-dropdown-item href="#">RU</b-dropdown-item>
                    </b-nav-item-dropdown>

                    <b-nav-item-dropdown right v-if="token">
                        <!-- Using 'button-content' slot -->
                        <template #button-content>
                            <em>{{ username }}</em>
                        </template>
                        <b-dropdown-item href="#">Profile</b-dropdown-item>
                        <b-dropdown-item @click="logout">Sign Out</b-dropdown-item>
                    </b-nav-item-dropdown>
                    <b-navbar-nav v-else>
                        <b-nav-item href="/login">Login</b-nav-item>
                        <b-nav-item href="/register">Register</b-nav-item>
                    </b-navbar-nav>
                </b-navbar-nav>
            </b-collapse>
        </b-navbar>
    </div>
</template>

<script>
import { getToken, getUsername } from '@/utils/auth';

export default {
    name: 'NavBar',
    data() {
        return {
            token: null,
            username: null
        };
    },
    created() {
        this.token = getToken();
        this.username = getUsername()
    },
    methods: {
        logout() {
            this.token = null;
            localStorage.removeItem('token');
        }
    }
};
</script>
