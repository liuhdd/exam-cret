import axios from 'axios';
import { ElMessage, ElMessageBox } from "element-plus";
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios';
import { useUserStoreHook } from '@/store/models/user'

const request = axios.create({
    baseURL: "http://localhost:8080",
    timeout: 5000,
    headers: { 'Content-Type': 'application/json' }
});

request.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        const userStore = useUserStoreHook()

        if (userStore.token) {
            config.headers.Authorization = userStore.token
        }
        return config;
    }, (error: any) => {
        return Promise.reject(error);
    }
);

request.interceptors.response.use(
    (response: AxiosResponse) => {
        if (response.status === 200) {
            return response.data;
        } else {
            ElMessage({
                message: response.data.message,
                type: 'error'
            });
            return Promise.reject(response.data.message || 'Error');
        }
    }, (error : any) => {

        const { message } = error.response.data
        if (error.response.status == 401) {

            ElMessageBox.confirm(message, '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'

            }).then(() => {
                localStorage.clear();
                window.location.href = '/'
            })
        } else {
            ElMessage({
                message: message,
                type: 'error'
            })
        }
        return Promise.reject(message || 'Error');
    }
)

export default request;