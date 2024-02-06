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

	const menuStyle = {
		width: "20vw",
		backgroundColor: "black",
		color: "white",
	}

	return (
		<div className='flex flex-row text-sm cursor-pointer w-full'>
			<div id="" className='flex flex-col text-xl font-bold' style={menuStyle}>
				<div className='m-2' onClick={updateTabEvent}>
					<div>inbox</div>
				</div>
				<div className='m-2' onClick={updateTabEvent}>
					<div>favourites</div>
				</div>
				<div className='m-2' onClick={updateTabEvent}>
					<div>history</div>
				</div>
				<div className='m-2' onClick={updateTabEvent}>
					<div>settings</div>
				</div>
				<div className='m-2' onClick={updateTabEvent}>
					<div>logout</div>
				</div>
			</div>
			<div id="" className='text-xl font-bold w-auto' onClick={updateTabEvent}>
				hello
			</div>
		</div>
	)

}