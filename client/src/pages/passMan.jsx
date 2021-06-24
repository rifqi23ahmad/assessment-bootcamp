import { deletePass, fetchPass } from "../redux/action/passAction"
import Navbar from "../component/navbar"
import { useDispatch, useSelector } from "react-redux"
import React, { useEffect } from "react"
import { useHistory } from "react-router-dom"
import { logoutUser } from "../redux/action/userAction"


function TablePass() {
    const history = useHistory()
    const dispatch = useDispatch()

    const { passw, isLoading } = useSelector(state => state.pass)

    console.log("inipassw",passw)

    useEffect(() => {
        dispatch(fetchPass())
    },[])

 

    // console.log(isLoading)
    console.log(passw.data)

    // if (isLoading){
    //     return (<h1>Loading</h1>) 
    // } else {

    const gotoCreate = () => {
        history.push("/add-site")
    }

    // const gotoLogin = () =>{
    //     history.push("/login")
    // }

    // const gotoUpdate = () =>{
    //     history.push("/updatePass/:pass_id")
    // }

        return(
            <div>
                <Navbar/>
                <div style={{padding: "10%"}}>
                    <button class="btn btn-primary" onClick={gotoCreate}>Create New Site</button>
                    <table class="table table-hover">
                        <thead>
                            <tr>
                            <th>NO</th>
                            <th scope="col">Website</th>
                            <th scope="col">Password</th>
                            <th scope="col">Update</th>
                            <th scope="col">Delete</th>
                            </tr>
                        </thead>
                        <tbody>
                        {passw && passw.map((pass, index) =>{
                            return(
                                <tr>
                                <td>{index+1}</td>
                                <td>{pass.webite}</td>
                                <td>{pass.site_pass}</td>
                                <td>
                                <button class="btn btn-warning" onClick={(e)=>{
                                    e.preventDefault()
                                        history.push("/edit-site/" + pass.ID)
                                }}>Update</button>
                                </td>
                                <td>
                                <button class="btn btn-danger" onClick={(e)=>{
                                    e.preventDefault()
                                    dispatch(deletePass(pass.ID, history))
                                    dispatch(fetchPass())
                                }}>Delete</button>
                                </td>
                            </tr>
                            )
                        })}
                        </tbody>
                    </table>
                </div>
            {/* <Footer/> */}
            </div>
        )
    // }
}


export default TablePass