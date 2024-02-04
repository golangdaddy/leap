import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"
import { format, formatDistance, formatRelative, subDays } from 'date-fns'

import Spacer from "@/inputs/spacer"

import { RowThumbnail } from "@/components/rowThumbnail"
import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"
import { RowOrder } from "@/components/rowOrder"
import { titlecase } from "../_interfaces"

export function {{titlecase .Object.Name}}ListRowJob(props) {

	const [userdata, setUserdata] = useUserContext()

	function selectItem(e) {
		console.log("SELECT EVENT", props.id)
		return props.select(props.id)
	}
	function deleteItem() {
		props.delete(props.id)
	}

	var date = new Date(props.item.Meta.Created * 1000);
	const dateTime = formatRelative(date, new Date())

	// Create a new Date object using the timestamp

	return (
		<div className='flex flex-row justify-between items-center w-full my-2'>
			<div onClick={selectItem} className='flex flex-row justify-between w-full items-center cursor-pointer m-4'>
				<div className='text-xl font-bold' title="">{ props.item.Meta.Name }</div>
				<div className='text-xl' title="">{ props.item.Meta.Context.Status }</div>
				<div className='text-xl' title="">{ dateTime }</div>
			</div>
		</div>
	)
}