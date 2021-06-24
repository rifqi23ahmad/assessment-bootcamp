import Axios from "axios"

const apiBase = Axios.create({
    baseURL : "http://root:pastilulus@localhost:8080"
})

export default apiBase