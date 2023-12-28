import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"

export function Preview(props) {

	const [userdata, setUserdata] = useUserContext()

	function selectItem(e) {
		console.log("SELECT EVENT", props.id)
		return props.select(props.id)
	}
	function deleteItem() {
		props.delete(props.id)
	}
	function moveUp() {
		props.moveUp(props.id)
	}
	function moveDown() {
		props.moveDown(props.id)
	}

	const iconStyle = {width:"30px",height:"30px"}
	const rowStyle = {}
	switch (props.item.type) {
	case "background": 
		rowStyle["backgroundColor"] = "rgba(0,0,255,0.3)"
		break
	case "element": 
		rowStyle["backgroundColor"] = "yellow"
		break
	case "foreground": 
		rowStyle["backgroundColor"] = "rgba(0,255,0,0.3)"
		break
	}

	return (
		<div className='flex flex-row justify-between items-center w-full my-2 px-4' style={rowStyle}>

			<div onClick={selectItem} className='flex flex-row w-full items-center'>
				<div className='text-xl font-bold'>
				{ props.overlay.description }
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				{ props.overlay.type }
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				"{ props.overlay.content }"
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				X:{ props.overlay.x }px
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				Y:{ props.overlay.y }px
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				{ props.overlay.font }
				</div>
				<div className="px-4"></div>
				<div className='text-l'>
				{ props.overlay.fontSize }pts
				</div>
			</div>
			<RowEdit object={props.overlay} editInterface="editcollection"/>
			<RowDelete id={props.id} delete={deleteItem}/>
		</div>
	)
}