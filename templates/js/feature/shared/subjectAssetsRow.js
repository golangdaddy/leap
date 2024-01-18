import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import Spacer from "@/inputs/spacer"

import { RowThumbnail } from "@/components/rowThumbnail"
import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"
import { RowOrder } from "@/components/rowOrder"
import { titlecase } from "../_interfaces"

export function {{titlecase .Object.Name}}AssetsRow(props) {

	const [userdata, setUserdata] = useUserContext()

	const [expanded, setExpanded] = useState(true)
	function toggleExpand() {
		setExpanded(!expanded)
		var p = document.getElementById("properties")
		if (expanded) {
			p.style.height = 'auto'
			p.style.height = p.scrollHeight + 'px'
		} else {
			p.style.height = 'auto'
		}
	}
	useEffect(() => {
		var p = document.getElementById("properties")
		p.innerHTML = JSON.stringify(props.item, null, 4)
	}, [])

	return (
		<div className='flex flex-col w-full' style={ {borderRadius:"10px", overflow:"hidden"} }>
			<div className="flex flex-row bg-gray-600 p-2 text-white">
				<div className="flex flex-col justify-center items-center mx-4"  style={ {width:"50px",height:"50px"} }>
					<div>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6 pointer-events-none">
						<path strokeLinecap="round" strokeLinejoin="round" d="M16.5 6v.75m0 3v.75m0 3v.75m0 3V18m-9-5.25h5.25M7.5 15h3M3.375 5.25c-.621 0-1.125.504-1.125 1.125v3.026a2.999 2.999 0 0 1 0 5.198v3.026c0 .621.504 1.125 1.125 1.125h17.25c.621 0 1.125-.504 1.125-1.125v-3.026a2.999 2.999 0 0 1 0-5.198V6.375c0-.621-.504-1.125-1.125-1.125H3.375Z" />
						</svg>
					</div>
				</div>
				<div className="flex flex-col px-2">
					<div className="flex flex-row">
						<div className="font-bold">assetId</div>
						<div className="px-2">:</div>
						<div>{props.item.assetId}</div>
					</div>
					<div className="flex flex-row text-sm">
						<div className="font-bold">collectionId</div>
						<div className="px-2">:</div>
						<div>{props.item.collectionId}</div>
					</div>
				</div>
			</div>
			<div className="flex flex-col w-full">
				<div className="w-full p-4 bg-gray-200">
					<textarea id='properties' className="w-full bg-gray-200"></textarea>
				</div>
				<div className="flex flex-row justify-center items-center p-2 bg-gray-600 text-white">
				{
					expanded && <button onClick={toggleExpand}>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6 pointer-events-none">
						<path strokeLinecap="round" strokeLinejoin="round" d="M3.75 3.75v4.5m0-4.5h4.5m-4.5 0L9 9M3.75 20.25v-4.5m0 4.5h4.5m-4.5 0L9 15M20.25 3.75h-4.5m4.5 0v4.5m0-4.5L15 9m5.25 11.25h-4.5m4.5 0v-4.5m0 4.5L15 15" />
						</svg>
					</button>
				}
				{
					!expanded && <button onClick={toggleExpand}>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6 pointer-events-none">
						<path strokeLinecap="round" strokeLinejoin="round" d="M9 9V4.5M9 9H4.5M9 9 3.75 3.75M9 15v4.5M9 15H4.5M9 15l-5.25 5.25M15 9h4.5M15 9V4.5M15 9l5.25-5.25M15 15h4.5M15 15v4.5m0-4.5 5.25 5.25" />
						</svg>
					</button>
				}
				</div>
			</div>
		</div>
	)
}	