import * as React from 'react'
import { useState, useEffect } from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'

import Sidebar from '@/features/account/sidebar'

import VisitTab from '@/features/interfaces'

import { InboxMessagesGET } from '@/app/fetch'

export default function AccountInboxMessages(props) {  

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	const [convo, setConvo] = useState(localdata.tab.context.conversation) 
	const [messages, setMessages] = useState([]) 

	function updateMessages() {
		InboxMessagesGET(userdata, convo)
        .then((res) => res.json())
		.then((data) => {
			console.log("MESSAGES", data)
			setMessages(data)
		})
		.catch((e) => {
			console.log(e)
		})		
	}

	const bodystyle = {
		borderRadius: "12px",
		border: "solid 1px black",
	}

	useEffect(() => {
		updateMessages()
	}, [])

	return (
		<div className='flex flex-row text-sm cursor-pointer w-full'>
			<Sidebar/>
			<div className='flex flex-col p-4  w-auto w-full'>
				<textarea id='body' className="w-full my-2 p-2" placeholder="your message" style={bodystyle}></textarea>
				<button style={buttonStyle} onClick={sendMessage}>Send</button>
			{
				messages.map(function (item, i) {
					return (
						<div key={i} id="" className='flex flex-row m-2'>
							<div>
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
								<path strokeLinecap="round" strokeLinejoin="round" d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5" />
								</svg>
							</div>
							<div>{item.body}</div>
						</div>
					)
				})
			}
			</div>
		</div>
	)

}