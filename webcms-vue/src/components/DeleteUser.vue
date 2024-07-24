<template>
  <div>
    <h1>Удалить пользователя</h1>
    <p>Вы уверены, что хотите удалить пользователя <strong>{{ user.name }}</strong>?</p>
    <b-button @click="deleteUser" variant="danger">Удалить</b-button>
    <b-button to="/users" variant="secondary">Отмена</b-button>
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
    console.error('Не удалось получить данные пользователя:', error)
  }
})

const deleteUser = async () => {
  try {
    await axios.delete(`/api/users/${route.params.id}`)
    await router.push('/users')
  } catch (error) {
    console.error('Не удалось удалить пользователя:', error)
  }
}
</script>
