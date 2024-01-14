import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import VisitTab from '@/features/interfaces'

export default function EditProfile(props) {  

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	// update tabs handles the updated context and sends the user to a new interface
	function updateTabEvent(e) {
		const id = e.target.id
		setLocaldata(VisitTab(localdata, id))
	}

	return (
		<div className='flex flex-col text-sm cursor-pointer'>
			<div id="" className='text-xl font-bold' onClick={updateTabEvent}>
				hello
			</div>
		</div>
	)

}