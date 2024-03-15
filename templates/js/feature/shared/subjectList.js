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
import { {{titlecase .Object.Name}}ListRowImage } from './{{lowercase .Object.Name}}ListRowImage';
import { {{titlecase .Object.Name}}DELETE, {{titlecase .Object.Name}}sListGET, {{titlecase .Object.Name}}OrderPOST, {{titlecase .Object.Name}}JobPOST } from '../_fetch';

{{if eq 1 (stringslength .Object.Parents)}}
import { {{firstparenttitle .Object.Parents}}JobPOST } from '@/features/{{firstparent .Object.Parents}}s/_fetch'
{{end}}

export function {{titlecase .Object.Name}}List(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [topics, setTopics] = useState([{{range .Object.Options.Topics}}{"name":"{{.Name}}","topic":"{{.Topic}}"},{{end}}])

	const [ list, setList ] = useState(null)

	var mode = "modified"
	{{if .Object.Options.Order}}mode = "ordered"{{end}}
	{{if .Object.Options.Admin}}mode = "admin"{{end}}
	{{if .Object.Options.EXIF}}mode = "exif"{{end}}

	function updateList() {
		{{titlecase .Object.Name}}sListGET(userdata, props.subject?.Meta.ID, mode, props.limit)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setList(data)
		}).catch((e) => {
			console.error("subjetList.updateList:", e)
		})
	}

	function sendToTopic(e) {
		console.log(e)
		const job = e.target.id
		{{if eq 1 (stringslength .Object.Parents)}}
		{{firstparenttitle .Object.Parents}}JobPOST(userdata, props.subject?.Meta.ID, job)
		.then((res) => console.log(res))
		.catch((e) => {
            console.error(e)
        })
		{{end}}
	}

	useEffect(() => {
		updateList()
	}, [])

	// update tabs handles the updated context and sends the user to a new interface
	function selectItem(id) {
		console.log("SELECT {{titlecase .Object.Name}}", id)
		const next = list[parseInt(id)]
		const context = {
			"_": (next.Meta.Name ? next.fields.name : next.fields.name),
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

	const jobButtonStyle = {
		borderRadius: "20px",
		backgroundColor: "rgb(96, 165, 250)",
		border: "solid 0px",
		color: "white",
		padding: "6px 10px"
	}

	return (
	<div className='flex flex-col w-full'>
		{
			props.title && <div className="flex flex-row justify-between items-center">
				<div className='py-4 my-4 text-xl font-bold'>{props.title}:</div>
				{
					(topics.length > 0) && <div className='flex flex-row'>
					{
						topics.map(function (item, i) {
							return (
								<div key={i} className='flex flex-col justify-center'>
									<button key={i} className='text-sm' id={item.topic} onClick={sendToTopic} style={jobButtonStyle}>
									{item.name}
									</button>
								</div>
							)
						})
					}
					</div>
				}
				{{range .Object.Options.FilterFields}}
					<div>
						<div>{{.Name}}</div>
						<select>
							<option>hello</option>
						</select>
					</div>
				{{end}}
			</div>
		}
		{
			props.title && <hr/>
		}
		{
			!list && <Loading/>
		}
		{{if .Object.Options.Image}}
		<div className='flex flex-wrap'>
		{
			list && list.map(function (item, i) {
				return (
					<div key={i} className='m-2'>
						<{{titlecase .Object.Name}}ListRowImage id={i} listLength={list.length} item={item} select={selectItem} moveUp={moveUp} moveDown={moveDown} delete={deleteItem}/>
					</div>
				)
			})
		}
		</div>
		{{end}}
		{{if eq false .Object.Options.Image}}
		{
			list && list.map(function (item, i) {

				return (
					<div className='{{if .Object.Options.Image}}flex flex-wrap{{end}} py-2 px-4' key={i}>
						{{if eq false .Object.Options.Job}}
							<{{titlecase .Object.Name}}ListRow id={i} listLength={list.length} item={item} select={selectItem} moveUp={moveUp} moveDown={moveDown} delete={deleteItem}/>
						{{end}}
						{{if .Object.Options.Job}}
							<{{titlecase .Object.Name}}ListRowJob id={i} listLength={list.length} item={item} select={selectItem} moveUp={moveUp} moveDown={moveDown} delete={deleteItem}/>
						{{end}}
					</div>
				)
			})
		}
		{{end}}
	</div>
	)

}