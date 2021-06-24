
import apiBase from "../../API/passManApi"

const access_token = !localStorage.getItem("access_token") ? "" : localStorage.getItem("access_token")

export const fetchPass = () => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const response = await apiBase({
                method : "GET",
                url : "/password/:id",
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(response.data)

           return dispatch({ type : "FETCH_PASS", payload : response.data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }
}


export const createPass = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const response =  await apiBase({
                method : "POST",
                url : "/password",
                data : payload,
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(response.data)

            return dispatch({ type : "CREATE_PASS", payload : response.data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }

}


export const updatePass = (id, payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const response = await apiBase({
                method : "PUT",
                url : `/password/:id${id}`,
                data : payload,
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(response.data)

            return dispatch({ type : "UPDATE_PASS", payload : response.data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }
}


export const deletePass = (id, history) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const response =  await apiBase({
                method : "DELETE",
                url : `/user/${id}`,
                headers : {
                    "authorization" : access_token
                }
            })
            history.push("/user")

            console.log(response.data)

           return dispatch({ type : "DELETE_PASS", payload : response.data})
        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }
}
