import * as React from 'react'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'

import { useUserContext } from "@/context/user"
import { useMessagingContext } from "@/context/messaging"
import { useLocalContext } from '@/context/local'

import VisitTab from '@/features/interfaces'

export default function Controller(props) {

	console.log("Controller props:", props)

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [messaging, setMessaging] = useMessagingContext()

	const [feed, setFeed] = useState([])

	const router = useRouter();


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
			<div className='items-center flex flex-row w-full bg-gray-600 text-white px-4 py-2 text-base shadow-xl'>
				{
					localdata?.breadcrumbs && <div className='my-4 w-full bg-gray-800 text-white flex justify-start' style={ {borderRadius:"12px",overflow:"hidden"} }>
						<div className="flex flex-col justify-center rounded-l-lg" style={ {backgroundColor:"#111111"} }>
							<div id="home" onClick={updateTabEvent} className="flex flex-col justify-center items-center m-4 cursor-pointer" style={ {width:"36px",height:"36px"} }>
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={ {pointerEvents:"none"} }>
								<path strokeLinecap="round" strokeLinejoin="round" d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
								</svg>
							</div>
						</div>
						{
							(localdata.breadcrumbs.length <= 1) && <div className='text-sm font-bold text-gray-800 flex flex-row w-full' style={ {backgroundColor:"#eeeeee"} }>
								<input className='border-none w-full px-4 text-base' style={ {backgroundColor:"#eeeeee"} }></input>
							</div>
						}
						{
							(localdata.breadcrumbs.length > 1) && <div className='text-sm font-bold text-gray-800 flex flex-row w-full px-2'  style={ {backgroundColor:"#eeeeee"} }>
							{
								localdata.breadcrumbs.map(function (tab, i) {

									if (i == 0) {
										return
									}

									if (!tab) {
										console.log("skipping missed tab")
										return
									}
									var displayName = tab.name

									if (tab.context?._) {
										displayName = tab.context._
									}

									return (
										<div key={i} className='flex flex-row items-center'>
											<div className='m-2'>/</div>
											<div className="ml-2 cursor-pointer" id={"crumb_"+i} onClick={ (i == (localdata.breadcrumbs.length - 1) ) ? function () {} : updateBackEvent}>
												{displayName}
											</div>
										</div>
									)
								})
							}
							<div className='px-2'/>
							</div>
						}
						<div className="flex flex-col justify-center rounded-l-lg" style={ {backgroundColor:"rgb(52 185 83)"} }>
							<div id="home" onClick={updateTabEvent} className="flex flex-col justify-center items-center m-4 cursor-pointer" style={ {width:"36px",height:"36px"} }>
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
									<path strokeLinecap="round" strokeLinejoin="round" d="m15.75 15.75-2.489-2.489m0 0a3.375 3.375 0 1 0-4.773-4.773 3.375 3.375 0 0 0 4.774 4.774ZM21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
								</svg>
							</div>
						</div>
					</div>
				}
			</div>
		</div>
	)
}