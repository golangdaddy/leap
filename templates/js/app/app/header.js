import * as React from 'react'

import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import VisitTab from '@/features/interfaces'

export default function Header() {

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext()

    function editProfile() {
        setLocaldata(VisitTab(localdata, "editprofile"))
    }

    return (
    <header className='w-full'>
        <div className='bg-black text-white flex flex-row justify-between p-2 px-6'>
            <div className='flex flex-col justify-center'>
                <div className='text-2xl font-bold'>{{.SiteName}}</div>
            </div>
            {
            userdata && <div className='flex flex-row items-center text-sm'>
                <div className='m-2'>{userdata.username}</div>
                <div className='m-2'>{userdata.email}</div>
                <div onClick={editProfile} className='m-2 cursor-pointer font-bold'>Edit Account</div>
            </div>
            }
            {
            !userdata && <div className='flex flex-col'>
                <a href="/otp">
                    <div id='account' className='cursor-pointer text-sm mt-2 font-bold'>Login</div>
                </a>
            </div>
            }
        </div>
    </header>
    )
}