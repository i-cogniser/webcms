<template>
  <div>
    <h1>Редактировать страницу</h1>
    <b-form @submit.prevent="updatePage">
      <b-form-group label="Название">
        <b-form-input v-model="page.title" type="text" required />
      </b-form-group>
      <b-form-group label="Содержание">
        <b-form-textarea v-model="page.content" rows="10" required />
      </b-form-group>
      <b-form-group label="Автор">
        <b-form-input v-model="page.author" type="text" required />
      </b-form-group>
      <b-button type="submit" variant="primary">Обновить страницу</b-button>
      <b-button to="/pages" variant="secondary">Назад к страницам</b-button>
    </b-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const page = ref({
  title: '',
  content: '',
  author: '',
})

const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(/api/pages/${route.params.id})
    page.value = response.data
  } catch (error) {
    console.error('Не удалось получить данные страницы:', error)
  }
})

const updatePage = async () => {
  try {
    await axios.put(/api/pages/${route.params.id}, page.value)
    router.push('/pages')
  } catch (error) {
    console.error('Не удалось обновить страницу:', error)
  }
}
</script>