// 通知显示函数
import {ElMessage, ElNotification} from 'element-plus'
export function eleNotice(type, msg){
    switch (type) {
        case 'success':
            ElNotification({
                message: msg,
                type: 'success',
                duration: 2000
            })
            break
        case 'warning':
            ElMessage({
                message: msg,
                type: 'warning',
                duration: 2000,
            })
            break
        case 'error':
            ElNotification({
                title: 'ERROR',
                message: msg,
                type: 'error',
                duration: 2000
            })
            break
        default:
            ElNotification({
                title: 'ERROR',
                message: 'please input correct notice type ,it`s must a string',
                type: 'error',
                duration: 2000
            })
    }
}

export function notImplement(name) {
    eleNotice('warning',name + ' is coming soon!')
}