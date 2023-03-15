import axios from 'axios'

export function userPasswordPatch(username: string, data: any) {
    return axios.patch(`/api/user/${username}/password`, data)
}