import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { ElMessage, ElMessageBox } from "element-plus";
import { localStorage } from "@/utils/storage";

const service = axios.create({
    baseURL: "http://localhost:8080",
    timeout: 5000,
    headers: { 'Content-Type': 'application/json' }
});

service.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        if (localStorage.getItem('token')) {
            config.headers['Authorization'] = localStorage.getItem('token');
        }
        return config;
    }, (error: any) => {
        return Promise.reject(error);
    }
);

service.interceptors.response.use(
    (response: AxiosResponse) => {
        if (response.data.code === 200) {
            return response.data;
        } else {
            ElMessage({
                message: response.data.message,
                type: 'error'
            });
            return Promise.reject(response.data.message || 'Error');
        }
        return response;

    }, (error : any) => {

        const { code, msg } = error.response.data
        if (code === 401) {

            ElMessageBox.confirm(msg, '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'

            }).then(() => {
                localStorage.removeItem('token');
                location.reload();
            })
        } else {
            ElMessage({
                message: msg,
                type: 'error'
            })
        }
        return Promise.reject(msg || 'Error');
    }
)

export default service;