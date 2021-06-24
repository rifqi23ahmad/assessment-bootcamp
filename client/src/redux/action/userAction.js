import apiBase from "../../API/passManApi"

export const register = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const { data } = await apiBase({
                method : "POST",
                url : "/user/register",
                data : payload,
            })

            console.log(data)

            return dispatch({ type : "USER_REGISTER", payload : data})

        } catch(err) {
            console.log(err.response)
        }
    }
}

export const login = (payload, history) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const response = await apiBase({
                method : "POST",
                url : "/user/login",
                data : payload,
            })

            console.log("ini respon", response.data)

            localStorage.setItem("access_token", response.data.authorization)

            console.log("data",response.data.authorization)
            console.log(payload)

            history.push("/pass")

            // return dispatch({ type : "USER_LOGIN", payload : data})


        } catch(err) {
            console.log(payload)
            dispatch({type : "USER_ERROR", payload : err})
            console.log(err.response)
        }
    }
}

export const logoutUser = () => ({
    type : "LOGOUT_USER"
})
