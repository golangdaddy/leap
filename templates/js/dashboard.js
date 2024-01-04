import * as React from 'react'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'

import { useUserContext } from "@/context/user"
import { useLocalContext } from '@/context/local'

import VisitTab from "@/features/interfaces"

import { AuthLoginPOST } from '@/app/fetch'

import Back from '@/app/back'
import Sublinks from '@/app/sublinks'
import Subsublinks from '@/app/subsublinks'

export default function Dashboard(props) {

		console.log("Dashboard props:", props)

		const [userdata, setUserdata] = useUserContext()
		const [localdata, setLocaldata] = useLocalContext()
		const [region, setRegion] = useState("UK")

		const router = useRouter();
		console.log("USERDATA", userdata)

		useEffect(() => {
			if (userdata === null) {
				updateTab("home")
				AuthLoginPOST(props.otp)
				.then((res) => res.json())
				.then((data) => {
					var user = data.user
					user.headers = {"Authorization": data.secret}
					setUserdata(user)
					console.log("DOWNLOADED USERDATA", user)
				})
				.catch((e) => {
					console.log(e)
					router("/otp")
				})
			}
		}, [])

		// update tabs handles the updated context and sends the user to a new interface
		function updateTabEvent(e) {
			updateTab(e.target.id)
		}

		// update tabs handles the updated context and sends the user to a new interface
		function updateTab(tabname) {
			const newState = VisitTab(localdata, tabname)
			console.log("DASHBOARD UPDATE STATE", tabname, newState)
			setLocaldata(newState)
		}

		// update tabs handles the updated context and sends the user to a new interface
		function updateProfile() {
			const tabname = "profile"
			const context = {
				"_": userdata.username,
				"object": userdata,
			}
			const newState = VisitTab(localdata, tabname, context)
			console.log("DASHBOARD UPDATE STATE", tabname, newState)
			setLocaldata(newState)
		}

		return (
			userdata && 
			<div className='flex flex-col justify-between bg-gray-800'>
				{ localdata && <Back/> }
				<div className='flex flex-row pos-absolute'>
					{ localdata?.tab && 
					<div className='flex flex-col w-full bg-white'>
						<Sublinks></Sublinks>
						<div className='flex flex-row justify-center'>
							<div className='m-10 text-3xl font-bold'>{localdata.tab.context._ ? localdata.tab.context._ : localdata.tab.name }</div>
						</div>
						<hr/>
						<Subsublinks></Subsublinks>
						<div className='flex flex-col justify-start w-full p-8'>
							{ localdata && localdata.tab && (localdata.tab.component) }
						</div>
					</div>
					}
				</div>
			</div>
		)
}
