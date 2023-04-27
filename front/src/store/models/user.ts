import { defineStore } from "pinia"
import { ref } from 'vue'
import { store } from '@/store'
import { loginApi, logoutApi} from '@/api/auth'
import type { LoginData } from "@/api/auth/types"

export const useUserStore = defineStore('user', () => {
    const nickname = ref('')
    const token = ref('')
    const avatar = ref('');
    function login(loginData : LoginData) {
        return new Promise<void>((resolve, reject) => {
            loginApi(loginData).then(response => {
                const { tokenType, accessToken } = response.data
                token.value = tokenType + ' ' + accessToken
                resolve()
            }).catch(error => {
                reject(error)
            })
        })
    }

    function logout() {
        return new Promise<void>((resolve, reject) => {
            logoutApi().then(response => {
                resetToken()
                resolve()
            }).catch(error => {
                reject(error)
            });
        })
    }
    function resetToken() {
        token.value = ''
        nickname.value = ''
    }
    return {
        nickname,
        token,
        avatar,
        login,
        logout
    }
})

export function useUserStoreHook() {
    return useUserStore(store);
}
