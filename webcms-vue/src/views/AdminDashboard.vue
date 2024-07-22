<template>
  <div>
    <h1>Панель Администратора</h1>
    <b-container>
      <b-row>
        <b-col md="4">
          <b-card title="Пользователи" class="mb-4">
            <b-card-text>
              <p>Всего пользователей: {{ userCount }}</p>
              <p>Управление пользователями системы.</p>
              <b-button @click="goToUsers" variant="primary">Перейти к пользователям</b-button>
            </b-card-text>
          </b-card>
        </b-col>
        <b-col md="4">
          <b-card title="Страницы" class="mb-4">
            <b-card-text>
              <p>Всего страниц: {{ pageCount }}</p>
              <p>Управление страницами веб-сайта.</p>
              <b-button @click="goToPages" variant="primary">Перейти к страницам</b-button>
            </b-card-text>
          </b-card>
        </b-col>
        <b-col md="4">
          <b-card title="Записи" class="mb-4">
            <b-card-text>
              <p>Всего записей: {{ postCount }}</p>
              <p>Управление записями блога.</p>
              <b-button @click="goToPosts" variant="primary">Перейти к записям</b-button>
            </b-card-text>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const userCount = ref(0)
const pageCount = ref(0)
const postCount = ref(0)

const goToUsers = () => {
  router.push('/users')
}

const goToPages = () => {
  router.push('/pages')
}

const goToPosts = () => {
  router.push('/posts')
}

onMounted(async () => {
  try {
    const userResponse = await axios.get('/api/users/count')
    userCount.value = userResponse.data.count
    const pageResponse = await axios.get('/api/pages/count')
    pageCount.value = pageResponse.data.count
    const postResponse = await axios.get('/api/posts/count')
    postCount.value = postResponse.data.count
  } catch (error) {
    console.error('Не удалось получить данные:', error)
  }
})
</script>

<style scoped>
h1 {
  color: #42b983;
  margin-bottom: 1rem;
}

b-card {
  text-align: center;
}
</style>