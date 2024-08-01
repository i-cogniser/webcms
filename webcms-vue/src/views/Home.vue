<template>
  <div>
    <h1>Добро пожаловать в CMSKA</h1>
    <b-button to="/users" variant="primary">Управление пользователями</b-button>
     <b-button @click="fetchUserData" variant="success">Загрузить данные пользователя</b-button>
    <b-button to="/register" variant="warning">Регистрация</b-button> <!-- Кнопка регистрации -->
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

const fetchUserData = () => {
  store.dispatch('fetchUser');
};
</script>
