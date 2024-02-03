import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import VisitTab from '@/features/interfaces';
import { titlecase } from '../_interfaces';
import Loading from '@/app/loading'
import Spacer from '@/inputs/spacer';

import { {{titlecase .Object.Name}}ListRow } from './{{lowercase .Object.Name}}ListRow';
import { {{titlecase .Object.Name}}ListRowJob } from './{{lowercase .Object.Name}}ListRowJob';
import { {{titlecase .Object.Name}}DELETE, {{titlecase .Object.Name}}sListGET, {{titlecase .Object.Name}}OrderPOST } from '../_fetch';

export function {{titlecase .Object.Name}}List(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ list, setList ] = useState(null)

	var mode = "modified"
	{{if .Object.Options.Order}}mode = "ordered"{{end}}
	{{if .Object.Options.Admin}}mode = "admin"{{end}}

	function updateList() {
		{{titlecase .Object.Name}}sListGET(userdata, props.subject?.Meta.ID, mode, props.limit)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setList(data)
		})
	}

	useEffect(() => {
		updateList()
	}, [])

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

	function selectChild() {
		setLocaldata(VisitTab(localdata, "{{lowercase .Object.Name}}s", context))
	}

	function moveUp(id) {
		const object = list[parseInt(id)]
		console.log("MOVE UP", object)
		{{titlecase .Object.Name}}OrderPOST(userdata, object.Meta.ID, "up")
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
		{{titlecase .Object.Name}}OrderPOST(userdata, object.Meta.ID, "down")
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
	<div className='flex flex-col my-4'>
	{
		props.title && <div className='py-4 my-4 text-xl font-bold cursor-pointer' onClick={selectChild}>{props.title}s:</div>
	}
	{
		props.title && <hr/>
	}
	{
		!list && <Loading/>
	}
	{
		list && list.map(function (item, i) {

			return (
				<div key={i}>
					{{if eq false .Object.Options.Job}}
					<{{titlecase .Object.Name}}ListRow id={i} listLength={list.length} item={item} select={selectItem} moveUp={moveUp} moveDown={moveDown} delete={deleteItem}/>
					{{end}}
					{{if .Object.Options.Job}}
					<{{titlecase .Object.Name}}ListRowJob id={i} listLength={list.length} item={item} select={selectItem} moveUp={moveUp} moveDown={moveDown} delete={deleteItem}/>
					{{end}}
					<Spacer/>
				</div>
			)
		})
	}
	</div>
	)

}