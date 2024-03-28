import {jwtCheck} from "../api/user";
import {useUserStore} from "../store/userProfile";
import {eleNotice} from "./notice.js";

// jwt验证，用于自动登录
export function jwtLogin() {
    // 引入store
    const userProfile = useUserStore()
    // 从localStorage中获取JWT
    const token = localStorage.getItem('jwtToken')
    if (token) {
        const headers = {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
        };
        // 发送JWT到后端验证
        jwtCheck(headers)
            .then(res => {
                // 验证通过，设置用户信息
                userProfile.setUserInfo(res.data.user, res.data.user, '', '', 0)
                eleNotice('success', 'jwt自动登录, 欢迎回来！')
            })
            .catch(err => {
                // 清除无效的JWT
                localStorage.removeItem('jwtToken');
            });
    }
}
