<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="username">Username:</label>
        <input v-model="username" id="username" type="text" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="password" id="password" type="password" required />
      </div>
      <button type="submit">Login</button>
      <div v-if="error" class="error">{{ error }}</div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const username = ref('')
const password = ref('')
const error = ref(null)
const router = useRouter()

const handleLogin = async () => {
  try {
    const response = await axios.post('/api/login', {
      email: this.email.value,
      password: password.value
    });
    localStorage.setItem('jwtToken', response.data.token);
    await router.push('/users');
  } catch (err) {
    error.value = 'Failed to login: ' + err.message;
    console.error('Failed to login:', err);
  }
};
</script>

<style>
.error {
  color: red;
}
</style>
