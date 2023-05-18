import { defineStore } from "pinia"
import { ref } from 'vue'
import { store } from '@/store'
import { loginApi, logoutApi} from '@/api/auth'
import type { LoginData, LoginResult } from "@/api/auth/types"

export const useUserStore = defineStore('user', () => {
    const username = ref('')
    const token = useStorage('accessToken', '')
    const avatar = ref('')
    const role = useStorage('role', '')
    function login(loginData : LoginData) {
        username.value = loginData.username
        return new Promise<void>((resolve, reject) => {
            loginApi(loginData).then(({data})=> {
                token.value = data.token
                role.value = data.role
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
        username.value = ''
    }
    return {
        username,
        token,
        avatar,
        role,
        login,
        logout
    }
})

export function useUserStoreHook() {
    return useUserStore(store);
}
