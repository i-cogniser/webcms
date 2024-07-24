<template>
  <div>
    <h1>Удалить страницу</h1>
    <p>Вы уверены, что хотите удалить страницу <strong>{{ page.title }}</strong>?</p>
    <b-button @click="deletePage" variant="danger">Удалить</b-button>
    <b-button to="/pages" variant="secondary">Отмена</b-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'

const page = ref(null)
const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    const response = await axios.get(`/api/pages/${route.params.id}`)
    page.value = response.data
  } catch (error) {
    console.error('Не удалось получить данные страницы:', error)
  }
})

const deletePage = async () => {
  try {
    await axios.delete(`/api/pages/${route.params.id}`)
    await router.push('/pages')
  } catch (error) {
    console.error('Не удалось удалить страницу:', error)
  }
}
</script>
