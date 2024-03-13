import * as React from 'react'
import { useEffect, useState } from 'react'
import Pusher from 'pusher-js';
import { useRouter } from 'next/router'

import { useUserContext } from "@/context/user"
import { useMessagingContext } from "@/context/messaging"
import { useLocalContext } from '@/context/local'

import Sidebar from '@/components/sidebar'

import VisitTab from '@/features/interfaces'

import { AuthLoginPOST } from '@/app/fetch'


import Subsublinks from '@/app/subsublinks'
import Controller from './controller'

export default function Dashboard(props) {

	console.log("Dashboard props:", props)

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [messaging, setMessaging] = useMessagingContext()

	const [feed, setFeed] = useState([])

	const router = useRouter();
	console.log("USERDATA", userdata)

	useEffect(() => {
		if (userdata === null) {
			updateTab("home")
			AuthLoginPOST(props.otp)
			.then((res) => res.json())
			.then((data) => {
				console.log("DOWNLOADED USERDATA", user)
				var user = data.user
				user.headers = {"Authorization": data.secret}
				setUserdata(user)
				// init websocket
				messaging.pusher = new Pusher('818e55ca022763d940aa', {
					cluster: 'eu',
					encrypted: true
				});
				setMessaging(messaging)
			})
			.catch((e) => {
				console.log(e)
//				router.push("/login")
			})
		}
	}, [])

	// update tabs handles the updated context and sends the user to a new interface
	function updateTabEvent(e) {
		console.log("UPDATE TAB EVENT:", e.target.id)
		updateTab(e.target.id)
	}
	function updateTab(tabname) {
		setLocaldata(VisitTab(localdata, tabname, localdata?.tab?.context))
	}

	// breadcrumbs
    function updateBackEvent(e) {
        updateBack(parseInt(e.target.id.split("crumb_")[1]))
    }
    function updateBack(crumbIndex) {
		console.log("C-INDEX", crumbIndex)
		const tab = localdata.breadcrumbs[crumbIndex]
        setLocaldata(VisitTab(localdata, tab.id, tab.context))
    }

	return (
		userdata && <div className='flex flex-col w-full'>
			<Controller/>
			{ 
				localdata?.tab && <div className='flex flex-row min-h-full w-full bg-white'>
					<div className='flex flex-col min-h-full w-full bg-white'>		
					{ 
						localdata && localdata.tab && localdata.tab.component
					}
					</div>
					{{if .Options.Sidebar}}<Sidebar/>{{end}}
				</div>
			}
		</div>
	)
}