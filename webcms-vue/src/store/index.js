import {createStore} from 'vuex'

export default createStore({
    state: {
        user: null, // начальное состояние для пользователя
        settings: {}, // начальное состояние для настроек
        items: [] // начальное состояние для списка элементов
    },
    mutations: {
        setUser(state, user) {
            state.user = user;
        },
        setSettings(state, settings) {
            state.settings = settings;
        },
        addItem(state, item) {
            state.items.push(item);
        }
    },
    actions: {
        fetchUser({commit}) {
            // Имитация запроса к API
            const user = {name: 'John Doe', email: 'john@example.com'};
            commit('setUser', user);
        },
        fetchSettings({commit}) {
            // Имитация запроса к API
            const settings = {theme: 'dark'};
            commit('setSettings', settings);
        },
        async fetchItems({commit}) {
            // Имитация асинхронного запроса к API
            const items = await new Promise(resolve => {
                setTimeout(() => {
                    resolve([{name: 'Item 1'}, {name: 'Item 2'}]);
                }, 1000);
            });
            items.forEach(item => {
                commit('addItem', item);
            });
        }
    },
    modules: {
        // Здесь можно определить модули, если они у вас есть
    }
})