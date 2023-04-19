import { defineStore } from "pinia"
import { ref } from 'vue'


export const useUserStore = defineStore('user', () => {
    const nickname = ref('')
    const token = ref('')

    return {
        nickname,
        token
    }
}

