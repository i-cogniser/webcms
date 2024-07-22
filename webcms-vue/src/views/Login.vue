<template>
  <div>
    <h1>Login</h1>
    <b-form @submit.prevent="login">
      <b-form-group label="Email">
        <b-form-input v-model="email" type="email" required />
      </b-form-group>
      <b-form-group label="Password">
        <b-form-input v-model="password" type="password" required />
      </b-form-group>
      <b-button type="submit" variant="primary">Login</b-button>
    </b-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const email = ref('')
const password = ref('')

const login = async () => {
  try {
    await axios.post('/api/login', { email: email.value, password: password.value })
    // Handle successful login
  } catch (error) {
    console.error('Login failed:', error)
  }
}
</script>