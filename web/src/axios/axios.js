import axios from "axios";
import { ElMessage } from 'element-plus';

axios.defaults.baseURL = "http://localhost:8000";

const request = axios.create({
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
});

request.interceptors.request.use(config => {
    config.headers['token'] = localStorage.getItem('token')
    return config
})

request.interceptors.response.use(response => {
    if (response.data.code === 200){
        console.log(response)
        let msg = response.data.message
        if (msg === "登录成功" || msg === "创建成功" || msg === "删除成功" || msg === "更新成功"){
            ElMessage.success(response.data.message)
        }
        return response;
    } else {
        ElMessage.error(response.data.message)
        console.log(response)
        return Promise.reject(response)
    }
},error => {
    console.log(error)
    return Promise.reject(error)
})

export default request;