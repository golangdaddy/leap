import * as React from 'react'
import { useState, useEffect } from 'react'

import { useLocalContext } from '@/context/local';
import { useUserContext } from '@/context/user';
import { HandcashMintPOST } from '@/app/fetch';

import { GoBack } from '@/features/interfaces'
import Loading from '@/app/loading';

export default function MintConfirmation() {

	const [localdata, setLocaldata] = useLocalContext()
	const [userdata, setUserdata] = useUserContext()

    const [mint] = useState(localdata.tab.context.mint)
    const [toggle, setToggle] = useState(true)

    function confirm() {
        setToggle(false)
        HandcashMintPOST(userdata.handcashToken, mint)
        .then(function () {
            setLocaldata(GoBack(localdata))
        })
        .catch(function (e) {
            console.error(e)
        })
    }

    return (        
    <div className="w-full p-10 flex flex-col items-center">
        <div className="flex flex-col">
        {
            mint.items.map(function (p, i) {
                return (
                    <div key={i} className='flex flex-row justify-between m-4 items-center' style={ {minWidth:"50vw"} }>
                        <div className='font-bold'>
                            {p.hello}
                        </div>
                    </div>
                )
            })
        }
        </div>
        <hr className='w-full'/>
        {
           toggle && <button onClick={confirm} className="m-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                Mint Tokens
            </button>
        }
        {
           toggle || <Loading/>
        }
    </div>
    )
}