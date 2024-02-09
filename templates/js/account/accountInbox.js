import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import Sidebar from '@/features/account/sidebar'

import VisitTab from '@/features/interfaces'
import { func } from 'prop-types'

import { InboxConvosGET } from '@/app/fetch'

export default function AccountInbox(props) {  

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	const [convos, setConvos] = useState([]) 

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

	function updateConvos() {
		InboxConvosGET(userdata)
        .then((res) => res.json())
		.then((data) => {
			console.log("UPDATED OBJECT",data)
			setConvos(data)
		})
		.catch((e) => {
			console.log(e)
		})		
	}

	return (
		<div className='flex flex-row text-sm cursor-pointer w-full'>
			<Sidebar/>
			<div className='flex flex-col'>
			{
				convos.map(function (item, i) {
					return (
						<div id="" className='text-xl font-bold w-auto' onClick={updateTabEvent}>
							hello
						</div>
					)
				})
			}
			</div>
		</div>
	)

}