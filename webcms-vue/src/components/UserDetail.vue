<template>
  <div v-if="user">
    <h1>User Details</h1>
    <p><strong>Name:</strong> {{ user.name }}</p>
    <p><strong>Email:</strong> {{ user.email }}</p>
    <b-button to="/users" variant="secondary">Back to Users</b-button>
    <b-button @click="deleteUser" variant="danger">Delete User</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const user = ref(null)
const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/users/${route.params.id}`)
    user.value = response.data
  } catch (error) {
    console.error('Failed to fetch user details:', error)
  }
})

const deleteUser = async () => {
  try {
    await axios.delete(`/api/users/${route.params.id}`)
    router.push('/users')
  } catch (error) {
    console.error('Failed to delete user:', error)
  }
}
</script>
