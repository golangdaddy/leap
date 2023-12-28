import { useUserContext } from '@/context/user'
import * as React from 'react'

export default function Header() {

    const [userdata, setUserdata] = useUserContext()

    return (
    <header className='w-full'>
        <div className='bg-white flex flex-row justify-between p-4'>
            <div className='flex flex-col justify-center px-10'>
                <div className='text-2xl font-bold'>NPG Platform</div>
            </div>
            {
            userdata && <div className='flex flex-col'>
                <div className='text-l'>{userdata.username}</div>
                <div>{userdata.email}</div>
                <div id='account' className='cursor-pointer text-sm mt-2 font-bold'>Edit Account</div>
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