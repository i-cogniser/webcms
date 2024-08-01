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

const goToUsers = () => {
  router.push('/users')
}


onMounted(async () => {
  try {
    const userResponse = await axios.get('/api/users/count')
    userCount.value = userResponse.data.count
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