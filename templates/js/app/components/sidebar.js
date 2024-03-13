import * as React from 'react'
import { useState, useEffect } from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useMessagingContext } from '@/context/messaging'

import VisitTab from '@/features/interfaces'

export default function Sidebar() {

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [messaging, setMessaging] = useMessagingContext()

	const [min, setMin] = useState(true)
	const [statuses, setStatuses] = useState({})
	const [hiddenFeed, setHiddenFeed] = useState([])

	function toggleMin() {
		setMin(!min)
	}

	useEffect(() => {

		const c = userdata.Meta.ID;
		console.log("CONNECTING TO PUSHER", c)
		messaging.channel = messaging.pusher.subscribe(c);
		messaging.channel.bind('create', data => {
			data.msg = "create"
			console.log("create MESSAGE !!!!!!!!!!!!", data)
			messaging.feed = [data, ...messaging.feed]
			setMessaging(messaging)
		});
		messaging.channel.bind('update', data => {
			data.msg = "update"
			console.log("update MESSAGE !!!!!!!!!!!!", data)
			messaging.feed = [data, ...messaging.feed]
			setMessaging(messaging)
		});
		messaging.channel.bind('job', data => {
			data.msg = "job"
			console.log("job MESSAGE !!!!!!!!!!!!", data)
			if (statuses[data.fields.id]) {
				console.log("HIDING", data.fields.id)
				setHiddenFeed([data])
			} else {
				messaging.feed = [data, ...messaging.feed]
				setMessaging(messaging)
			}
			statuses[data.fields.id] = data.Meta.Context.Status
			setStatuses(statuses)
		});
	}, [])


	function visitJob(e) {
		const x = parseInt(e.target.id)
		var context = {
			"object": messaging.feed[x]
		}
		setLocaldata(VisitTab(localdata, "job", context))
	}

	const jobEventStyle = {
		border: "solid 1px rgb(96, 165, 250)",
		color: "rgb(96, 165, 250)"
	}

	return (
	<div className='flex flex-col bg-black p-2'>
		{
			!min && <div onClick={toggleMin} className='flex flex-col justify-center items-center w-6 h-6 h-full mr-2 cursor-pointer' style={ {pointerEvents:"none"} }>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" style={ {pointerEvents:"none"} }>
				<path strokeLinecap="round" strokeLinejoin="round" d="M15.59 14.37a6 6 0 0 1-5.84 7.38v-4.8m5.84-2.58a14.98 14.98 0 0 0 6.16-12.12A14.98 14.98 0 0 0 9.631 8.41m5.96 5.96a14.926 14.926 0 0 1-5.841 2.58m-.119-8.54a6 6 0 0 0-7.381 5.84h4.8m2.581-5.84a14.927 14.927 0 0 0-2.58 5.84m2.699 2.7c-.103.021-.207.041-.311.06a15.09 15.09 0 0 1-2.448-2.448 14.9 14.9 0 0 1 .06-.312m-2.24 2.39a4.493 4.493 0 0 0-1.757 4.306 4.493 4.493 0 0 0 4.306-1.758M16.5 9a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z" />
				</svg>
			</div>
		}
		{ min && <div className='flex flex-col' style={ {width:"30vw"} }>
			{
				messaging.feed.map(function (item, i) {
					switch (item.msg) {
					case "job":
						return (
							<div key={i} className='mx-4 my-2 p-2 flex flex-col' style={jobEventStyle}>
								<div id={i} className='flex flex-row' onClick={visitJob}>
									<div className='flex flex-col justify-center items-center w-6 h-6 h-full mr-2 cursor-pointer' style={ {pointerEvents:"none"} }>
										<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" style={ {pointerEvents:"none"} }>
										<path strokeLinecap="round" strokeLinejoin="round" d="M15.59 14.37a6 6 0 0 1-5.84 7.38v-4.8m5.84-2.58a14.98 14.98 0 0 0 6.16-12.12A14.98 14.98 0 0 0 9.631 8.41m5.96 5.96a14.926 14.926 0 0 1-5.841 2.58m-.119-8.54a6 6 0 0 0-7.381 5.84h4.8m2.581-5.84a14.927 14.927 0 0 0-2.58 5.84m2.699 2.7c-.103.021-.207.041-.311.06a15.09 15.09 0 0 1-2.448-2.448 14.9 14.9 0 0 1 .06-.312m-2.24 2.39a4.493 4.493 0 0 0-1.757 4.306 4.493 4.493 0 0 0 4.306-1.758M16.5 9a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z" />
										</svg>
									</div>
									<div className='flex flex-col items-left' style={ {pointerEvents:"none"} }>
										<div className='flex flex-row font-xs'>
											<div className='font-bold'>{item.msg}</div>
											<div className='font-bold'>:</div>
											<div className=''>{item.fields.id}</div>
										</div>
										<div className='font-xs'>{item.fields.name}</div>
										<div className='font-sm font-bold'>{statuses[item.fields.id]}</div>
									</div>
								</div>
							</div>
						)
					default:
						return (
							<div key={i} className='m-4 p-4 border-dotted border-2'>
								New {item.Meta.Class}
							</div>
						)
					}
				})
			}
		</div> }
	</div>
	)
}