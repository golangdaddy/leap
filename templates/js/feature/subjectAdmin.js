import { RowDelete } from "@/components/rowDelete"
import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import { UserObjectGET } from "@/app/fetch"
import Loading from "@/app/loading"

export function {{titlecase .Object.Name}}Admin(props) {

    const [userdata, setUserdata] = useUserContext()

    const [admin, setAdmin] = useState()

    useEffect(() => {
        UserObjectGET(userdata, props.admin)
        .then((res) => res.json())
        .then((data) => {
            console.log("RESOLVED USERNAME: "+data.username)
            setAdmin(data)
        })
    }, [])

    function deleteAdmin() {
        props.delete(props.id)
    }

    return (
        <div className='flex flex-row w-full my-2 justify-between'>

            <div className='flex flex-row justify-between w-full'>
                {
                    admin && <div className='flex flex-col justify-center text-2xl uppercase m-4'>
                        { admin.username }
                    </div>
                }
            </div>

            { admin && <RowDelete id={props.id} delete={deleteAdmin}/> }
        </div>
    )
}