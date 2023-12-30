import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import Spacer from "@/inputs/spacer"
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

	return (
		<div className='flex flex-row justify-between items-center w-full my-2 px-4'>

			<div onClick={selectItem} className='flex flex-row w-full items-center cursor-pointer'>
				{{range $item, $key := .Object.Fields}}<div className='text-xl font-bold' title="{{lowercase $key.Name}}">{ props.item.fields["{{lowercase $key.Name}}"] }</div>
				<div className="px-4"></div>
				<Spacer/>{{end}}
			</div>
			<RowEdit object={props.item} editInterface="edit{{lowercase .Object.Name}}"/>
			<RowDelete id={props.id} delete={deleteItem}/>
		</div>
	)
}