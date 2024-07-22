<template>
  <div>
    <h1>Register</h1>
    <b-form @submit.prevent="register">
      <b-form-group label="Name">
        <b-form-input v-model="name" type="text" required />
      </b-form-group>
      <b-form-group label="Email">
        <b-form-input v-model="email" type="email" required />
      </b-form-group>
      <b-form-group label="Password">
        <b-form-input v-model="password" type="password" required />
      </b-form-group>
      <b-button type="submit" variant="primary">Register</b-button>
    </b-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const name = ref('')
const email = ref('')
const password = ref('')

const register = async () => {
  try {
    await axios.post('/api/register', { name: name.value, email: email.value, password: password.value })
    // Handle successful registration
  } catch (error) {
    console.error('Registration failed:', error)
  }
}
</script>