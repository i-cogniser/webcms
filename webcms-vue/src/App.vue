<script setup>
import { RouterView, RouterLink } from 'vue-router';
import axios from 'axios';

// Проверка работы прокси и доступности бэкенда
axios.get('/api/health')
    .then(response => console.log('Health check response:', response.data))
    .catch(error => console.error('Health check error:', error));
</script>

<template>
  <div>
    <header>
      <nav>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" class="logo" alt="Vite logo"/>
        </a>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/register">Register</RouterLink>
        <RouterLink to="/users">Users</RouterLink>
      </nav>
    </header>
    <RouterView />
  </div>
</template>

<style scoped>
header {
  background-color: #f8f9fa;
  padding: 1em;
}

nav {
  display: flex;
  gap: 1em;
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
