import { useState, useEffect } from "react"
import { useUserContext } from "@/context/user"

import Spacer from "@/inputs/spacer"

import { RowThumbnail } from "@/components/rowThumbnail"
import { RowDelete } from "@/components/rowDelete"
import { RowEdit } from "@/components/rowEdit"
import { RowOrder } from "@/components/rowOrder"
import { RowPay } from "@/components/rowPay"
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
		<div className='flex flex-row justify-between py-2 items-center w-full'>
			{{if .Object.Options.Image}}
			<RowThumbnail source={'https://storage.googleapis.com/{{.ProjectName}}-uploads/'+props.item.Meta.Media.URIs[props.item.Meta.Media.URIs.length-1]}/>
			{{end}}
			<div onClick={selectItem} className='flex flex-row w-full items-center cursor-pointer px-4'>
				{
					props.item.Meta.Name?.length && <>
						<div className='text-lg font-bold' title="Name">{ props.item.Meta.Name }</div>
					</>
				}
				<div className="px-4"></div>
				{{range $item, $key := .Object.Fields}}{
					("{{lowercase $key.Name}}" != "name") && !Array.isArray(props.item.fields["{{lowercase $key.Name}}"]) &&  !(typeof props.item.fields["{{lowercase $key.Name}}"] === 'object')  && <>
						<div className='text-sm font-bold' title="{{lowercase $key.Name}}">
							{ props.item.fields["{{lowercase $key.Name}}"] }
						</div>
						<div className="px-4"></div>
					</>
				}{{end}}
			</div>
			<RowPay id={props.id}/>

			{{if .Object.Options.Order}}<RowOrder id={props.id} listLength={props.listLength} moveUp={props.moveUp} moveDown={props.moveDown}/>{{end}}
			<RowEdit object={props.item} editInterface="edit{{lowercase .Object.Name}}"/>
			<RowDelete id={props.id} delete={deleteItem}/>
		</div>
	)
}