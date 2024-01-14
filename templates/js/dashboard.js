import * as React from 'react'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'

import { useUserContext } from "@/context/user"
import { useMessagingContext } from "@/context/messaging"
import { useLocalContext } from '@/context/local'

import VisitTab from '@/features/interfaces'

import { AuthLoginPOST } from '@/app/fetch'

import Subsublinks from '@/app/subsublinks'

export default function Dashboard(props) {

	console.log("Dashboard props:", props)

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [messaging, setMessaging] = useMessagingContext()

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
				messaging.socket = new WebSocket('ws://localhost:8080/ws');
				messaging.socket.addEventListener('open', (event) => {
					console.log('Connected to WebSocket server');
				});
				messaging.socket.addEventListener('message', (event) => {
					console.log('Received message:', event.data);

					var msg = JSON.parse(event.data)
					if (messaging[msg.Test] == null) {
						messaging[msg.Test] = []
					}
					messaging[msg.Test].push(event.data)
					setUserdata(user)
				});

			})
			.catch((e) => {
				console.log(e)
				router.push("/otp")
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
		userdata && <div className='flex flex-col'>
			<div className='items-center flex w-full bg-gray-600 text-white px-2 py-2 text-xl shadow-xl'>
				{
					localdata?.tab && <div className='flex flex-col justify-start px-4 pl-6'>
						{ 
							localdata.tab.context.object && <div className='text-sm font-bold' style={{whiteSpace: "nowrap"}}>
								{localdata.tab.context.object.Meta.Class}
							</div>
						}
						<div className='text-2xl' style={{whiteSpace: "nowrap"}}>{localdata.tab.context._ ? localdata.tab.context._ : localdata.tab.name }</div>
					</div>
				}
				{
					localdata?.tab && localdata.tab.hasNewButton && <div className="m-2 flex flex-col justify-center">
						<div id={'new'+localdata.tab.name.substr(0, localdata.tab.name.length-1).toLowerCase()}
						onClick={updateTabEvent} className="flex flex-col justify-center items-center m-2 cursor-pointer" style={{width:"20px",height:"20px"}}>
							<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={{pointerEvents:"none"}}>
							<path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
							</svg>
						</div>
					</div>
				}
				{
					localdata?.tab && localdata.tab.hasEditButton && <div className="m-2 flex flex-col justify-center">
						<div id={'edit'+localdata.tab.context.object.Meta.Class.substr(0, localdata.tab.context.object.Meta.Class.length-1)}
						onClick={updateTabEvent} className="flex flex-col justify-center items-center m-2 cursor-pointer" style={{width:"20px",height:"20px"}}>
							<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-4 h-4" style={{pointerEvents:"none"}}>
							<path strokeLinecap="round" strokeLinejoin="round" d="M11.42 15.17 17.25 21A2.652 2.652 0 0 0 21 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 1 1-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 0 0 4.486-6.336l-3.276 3.277a3.004 3.004 0 0 1-2.25-2.25l3.276-3.276a4.5 4.5 0 0 0-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437 1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008v-.008Z" />
							</svg>
						</div>
					</div>
				}
				{
					localdata?.tab && localdata.tab.hasDeleteButton && <div className="m-2 flex flex-col justify-center">
						<div id={'delete'+localdata.tab.context.object.Meta.Class.substr(0, localdata.tab.context.object.Meta.Class.length-1)}
						onClick={updateTabEvent} className="flex flex-col justify-center items-center m-2 cursor-pointer" style={{width:"20px",height:"20px"}}>
							<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-4 h-4" style={{pointerEvents:"none"}}>
							<path strokeLinecap="round" strokeLinejoin="round" d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5m6 4.125l2.25 2.25m0 0l2.25 2.25M12 13.875l2.25-2.25M12 13.875l-2.25 2.25M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" />
							</svg>
						</div>
					</div>
				}
				{
					localdata?.breadcrumbs && <div className='my-4 w-full bg-gray-800 text-white flex justify-start' style={{borderRadius:"12px",overflow:"hidden"}}>
						<div className="flex flex-col justify-center rounded-l-lg" style={{backgroundColor:"#111111"}}>
							<div id="home" onClick={updateTabEvent} className="flex flex-col justify-center items-center m-4 cursor-pointer" style={{width:"36px",height:"36px"}}>
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={{pointerEvents:"none"}}>
								<path strokeLinecap="round" strokeLinejoin="round" d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
								</svg>
							</div>
						</div>
						{
							(localdata.breadcrumbs.length <= 1) && <div className='text-sm font-bold text-gray-800 flex flex-row w-full' style={{backgroundColor:"#eeeeee"}}>
								<input className='border-none w-full px-4 text-base' style={{backgroundColor:"#eeeeee"}}></input>
							</div>
						}
						{
							(localdata.breadcrumbs.length > 1) && <div className='text-sm font-bold text-gray-800 flex flex-row w-full px-2'  style={{backgroundColor:"#eeeeee"}}>
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
						<div className="flex flex-col justify-center rounded-l-lg" style={{backgroundColor:"rgb(52 185 83)"}}>
							<div id="home" onClick={updateTabEvent} className="flex flex-col justify-center items-center m-4 cursor-pointer" style={{width:"36px",height:"36px"}}>
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
									<path strokeLinecap="round" strokeLinejoin="round" d="m15.75 15.75-2.489-2.489m0 0a3.375 3.375 0 1 0-4.773-4.773 3.375 3.375 0 0 0 4.774 4.774ZM21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
								</svg>
							</div>
						</div>
					</div>
				}
			</div>
			{ 
				localdata?.tab && <div className='flex flex-col min-h-full w-full bg-white'>
					<Subsublinks></Subsublinks>
					<div className='' style={{padding:"30px 60px 30px 60px"}}>
						{ localdata && localdata.tab && (localdata.tab.component) }
					</div>
				</div>
			}
		</div>
	)
}