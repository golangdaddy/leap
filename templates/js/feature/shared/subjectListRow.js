import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import Spacer from "@/inputs/spacer"

import { RowThumbnail } from "@/components/rowThumbnail"
import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"
import { RowOrder } from "@/components/rowOrder"
import { titlecase } from "../_interfaces"

export function {{titlecase .Object.Name}}ListRow(props) {

	const [userdata, setUserdata] = useUserContext()

	function selectItem(e) {
		console.log("SELECT EVENT", props.id)
		return props.select(props.id)
	}
	function deleteItem() {
		props.delete(props.id)
	}

	return (
		<div className='flex flex-row justify-between items-center w-full my-2'>
			{{if .Object.Options.Image}}<RowThumbnail source={'https://storage.googleapis.com/{{.DatabaseID}}-uploads/'+props.item.Meta.URIs[props.item.Meta.URIs.length-1]}/>{{end}}
			<div onClick={selectItem} className='flex flex-row w-full items-center cursor-pointer m-4'>
				{{range $item, $key := .Object.Fields}}<div className='text-xl font-bold' title="{{lowercase $key.Name}}">{ props.item.fields["{{lowercase $key.Name}}"] }</div>
				<div className="px-4"></div>
				<Spacer/>{{end}}
			</div>
			{{if .Object.Options.Order}}<RowOrder id={props.id} listLength={props.listLength} moveUp={props.moveUp} moveDown={props.moveDown}/>{{end}}
			<RowEdit object={props.item} editInterface="edit{{lowercase .Object.Name}}"/>
			<RowDelete id={props.id} delete={deleteItem}/>
		</div>
	)
}