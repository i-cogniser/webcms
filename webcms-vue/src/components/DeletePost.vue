<template>
  <div>
    <h1>Удалить запись</h1>
    <p>Вы уверены, что хотите удалить запись <strong>{{ post.title }}</strong>?</p>
    <b-button @click="deletePost" variant="danger">Удалить</b-button>
    <b-button to="/posts" variant="secondary">Отмена</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const post = ref(null)
const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/posts/${route.params.id}`)
    post.value = response.data
  } catch (error) {
    console.error('Не удалось получить данные записи:', error)
  }
})

const deletePost = async () => {
  try {
    await axios.delete(`/api/posts/${route.params.id}`)
    await router.push('/posts')
  } catch (error) {
    console.error('Не удалось удалить запись:', error)
  }
}
</script>
