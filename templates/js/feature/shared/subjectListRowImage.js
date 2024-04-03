import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import Spacer from "@/inputs/spacer"

import { RowThumbnail } from "@/components/rowThumbnail"
import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"
import { RowOrder } from "@/components/rowOrder"
import { titlecase } from "../_interfaces"

export function {{titlecase .Object.Name}}ListRowImage(props) {

	const [userdata, setUserdata] = useUserContext()

	function selectItem(e) {
		console.log("SELECT EVENT", props.id)
		return props.select(props.id)
	}
	function deleteItem() {
		props.delete(props.id)
	}

	return (
		<div className='flex flex-col justify-between items-center w-full'>
			{{if .Object.Options.Image}}
			<div  id={props.item.Meta.ID} onClick={selectItem} className="cursor-pointer"><img src={props.item.Meta.Media.Preview}/></div>
			{{end}}
			<div className='flex flex-col w-full justify-center items-center m-4'>
				<div className="px-4"></div>
				{{range $item, $key := .Object.Fields}}{
					("{{lowercase $key.Name}}" != "name") && !Array.isArray(props.item.fields["{{lowercase $key.Name}}"]) &&  !(typeof props.item.fields["{{lowercase $key.Name}}"] === 'object')  && <>
						<div className='text-base font-bold' title="{{lowercase $key.Name}}">
							"{ props.item.fields["{{lowercase $key.Name}}"] }"
						</div>
						<div className="px-4"></div>
					</>
				}{{end}}
			</div>
			<div className="flex flex-rowc {{if .Object.Options.Image}}my-2{{end}}">
				{{if .Object.Options.Order}}<RowOrder id={props.id} listLength={props.listLength} moveUp={props.moveUp} moveDown={props.moveDown}/>{{end}}
				<RowEdit object={props.item} editInterface="edit{{lowercase .Object.Name}}"/>
				<RowDelete id={props.id} delete={deleteItem}/>
			</div>
		</div>
	)
}