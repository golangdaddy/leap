import * as React from 'react'
import { useState } from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useMessagingContext } from '@/context/messaging'

import VisitTab from '@/features/interfaces'

export default function Sidebar() {

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [messaging, setMessaging] = useMessagingContext()

	const [feed, setFeed] = useState([])

	const c = userdata.Meta.ID;
	console.log("CONNECTING TO PUSHER", c)
	messaging.channel = messaging.pusher.subscribe(c);
	messaging.channel.bind('create', data => {
		console.log("create MESSAGE !!!!!!!!!!!!")
		setFeed([...feed, data]);
	});
	messaging.channel.bind('update', data => {
		console.log("update MESSAGE !!!!!!!!!!!!")
		setFeed([...feed, data]);
	});

	return (
		<div className='flex flex-col bg-gray-200' style={ {width:"30vw"} }>
			{
				feed.map(function (item, i) {
					return (
						<div key={i} className='m-4 p-4 border-solid'>
							New {item.Meta.Class}
						</div>
					)
				})
			}
		</div>
	)
}