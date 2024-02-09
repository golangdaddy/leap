import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import VisitTab from '@/features/interfaces'

export default function Sidebar(props) {  

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
		<div id="" className='flex flex-col text-lg font-bold items-center' style={menuStyle}>
			<div id="accountinbox" className='m-2' onClick={updateTabEvent}>
				inbox
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
	)

}