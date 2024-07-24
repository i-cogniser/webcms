<template>
  <div>
    <h1>Редактировать запись</h1>
    <b-form @submit.prevent="updatePost">
      <b-form-group label="Название">
        <b-form-input v-model="post.title" type="text" required />
      </b-form-group>
      <b-form-group label="Содержание">
        <b-form-textarea v-model="post.content" rows="10" required />
      </b-form-group>
      <b-form-group label="Автор">
        <b-form-input v-model="post.author" type="text" required />
      </b-form-group>
      <b-button type="submit" variant="primary">Обновить запись</b-button>
      <b-button to="/posts" variant="secondary">Назад к записям</b-button>
    </b-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const post = ref({
  title: '',
  content: '',
  author: '',
})

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

const updatePost = async () => {
  try {
    await axios.put(`/api/posts/${route.params.id}`, post.value)
    router.push('/posts')
  } catch (error) {
    console.error('Не удалось обновить запись:', error)
  }
}
</script>