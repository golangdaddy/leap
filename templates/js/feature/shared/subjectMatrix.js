import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import VisitTab from '@/features/interfaces';
import { titlecase } from '../_interfaces';

import Loading from '@/app/loading';
import Spacer from '@/inputs/spacer';

import { {{titlecase .Object.Name}}MatrixRow } from './{{lowercase .Object.Name}}MatrixRow';
import {
	{{titlecase .Object.Name}}DELETE,
	{{titlecase .Object.Name}}sListGET,
	{{titlecase .Object.Name}}MoveUpPOST,
	{{titlecase .Object.Name}}MoveDownPOST,
} from '../_fetch';
import ObjectPATCH from '@/app/fetch'


export function {{titlecase .Object.Name}}Matrix(props) {

	// set props.limit if you want to limit query results

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ list, setList ] = useState(null)

	function updateList() {
		{{titlecase .Object.Name}}sListGET(userdata, props.subject?.Meta.ID, props.limit)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setList(data)
		})
	}

	useEffect(() => {
		updateList()
	}, [])

	function saveUpdate(rowID, fieldID, value) {
		const row = list[rowID]
		console.log("SAVEUPDATE", row, fieldID, value)
		row.fields[fieldID] = value
		ObjectPATCH(userdata, row, fieldID, value)
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
		setList(list)
	}

	// update tabs handles the updated context and sends the user to a new interface
	function selectItem(id) {
		console.log("SELECT {{titlecase .Object.Name}}", id)
		const next = list[parseInt(id)]
		const context = {
			"_": next.fields.name,
			"object": next,
		}
		setLocaldata(VisitTab(localdata, "{{lowercase .Object.Name}}", context))
	}

	function moveUp(id) {
		const object = list[parseInt(id)]
		console.log("MOVE UP", object)
		{{titlecase .Object.Name}}MoveUpPOST(userdata, object.Meta.ID)
		.then((res) => console.log(res))
		.then(function () {
			updateList()
		})
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	function moveDown(id) {
		const object = list[parseInt(id)]
		console.log("MOVE DOWN", object)
		{{titlecase .Object.Name}}MoveDownPOST(userdata, object.Meta.ID)
		.then((res) => console.log(res))
		.then(function () {
			updateList()
		})
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	function deleteItem(id) {
		const object = list[parseInt(id)]
		console.log("DELETING", object)
		{{titlecase .Object.Name}}DELETE(userdata, object.Meta.ID)
		.then((res) => console.log(res))
		.then(function () {
			updateList()
		})
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	return (
	<>
	{
		!list && <Loading/>
	}
		<table className='w-full' style={ {border:"1px solid"} }><tbody>
			<tr>
				{{range .Object.Fields}}<td className='font-bold px-2' style={ {border:"1px solid"} }>{{lowercase .Name}}</td>{{end}}
			</tr>
			{
				list && list.map(function (row, i) {
					return (
						<{{titlecase .Object.Name}}MatrixRow id={i} key={i} row={row} save={saveUpdate}/>
					)
				})
			}
		</tbody></table>
		<Spacer/>

	</>
	)

}