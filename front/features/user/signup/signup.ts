import axios  from 'axios';

axios.defaults.withCredentials = true;

export type UserInfo = {
    name: string,
    email: string,
    password: string
}

export const signup = (userInfo: UserInfo) => {
    axios.post('http://localhost:8080/api/user/signup', userInfo, { withCredentials: true })
    .then( res => {
        console.log(res)
    })
}