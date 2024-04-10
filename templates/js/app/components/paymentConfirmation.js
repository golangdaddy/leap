import * as React from 'react'
import { useState } from 'react'

import { useLocalContext } from '@/context/local';
import { useUserContext } from '@/context/user';
import { HandcashPaymentPOST } from '@/app/fetch';

export default function PaymemtConfirmation() {

	const [localdata, setLocaldata] = useLocalContext()
	const [userdata, setUserdata] = useUserContext()

    const [payment] = useState(localdata.tab.context.payment)

    function confirm() {
        HandcashPaymentPOST(payment)
    }

    return (        
    <div className="w-full px-8 py-4 flex flex-col">
        <div className="text-sm text-gray-500 sm:text-center dark:text-gray-400">
        <b>{payment.currency}</b> {payment.amount}
        </div>
        <button onClick={confirm} className="my-5 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Send Payment</button>
    </div>
    )
}