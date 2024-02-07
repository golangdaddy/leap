import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import Sidebar from '@/features/account/sidebar'

import VisitTab from '@/features/interfaces'

export default function Account(props) {  

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	// update tabs handles the updated context and sends the user to a new interface
	function updateTabEvent(e) {
		const id = e.target.id
		setLocaldata(VisitTab(localdata, id))
	}

	const menuStyle = {
		width: "20vw",
		backgroundColor: "black",
		color: "white",
	}

	return (
		<div className='flex flex-row text-sm cursor-pointer w-full'>
			<Sidebar/>
			<div id="" className='text-xl font-bold w-auto' onClick={updateTabEvent}>
				hello world
			</div>
		</div>
	)

}