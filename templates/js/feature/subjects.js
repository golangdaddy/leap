import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import Loading from '@/app/loading'

import { AI } from './forms/ai';
import { {{titlecase .Object.Name}}List } from './shared/{{lowercase .Object.Name}}List';

import VisitTab from '../interfaces';

export function {{titlecase .Object.Name}}s(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ subject ] = useState(localdata.tab.context.object)

	const [promptToggle, setPromptToggle] = useState(true)

	function updateList(state) {
		setPromptToggle(state)
	}

	// update tabs handles the updated context and sends the user to a new interface
	function updateTabEvent(e) {
		console.log("UPDATE TAB EVENT:", e.target.id)
		updateTab(e.target.id)
	}
	function updateTab(tabname) {
		setLocaldata(VisitTab(localdata, tabname, localdata?.tab?.context))
	}

	const buttonStyle = {
		borderRadius: "12px",
		backgroundColor: "rgb(96, 165, 250)",
		border: "solid 0px",
		color: "white",
		padding: "6px 10px"
	}

	return (
		<div style={ {padding:"30px 60px 30px 60px"} } className='flex flex-col w-full'>
			<div className='flex flex-row justify-between w-full'>
				<div className='flex flex-row'>
					<button id={'new{{lowercase .Object.Name}}'} onClick={updateTabEvent} className="flex flex-col justify-center items-center m-2 cursor-pointer" style={buttonStyle}>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={ {pointerEvents:"none"} }>
						<path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
						</svg>
					</button>
					{{if .Object.Options.File}}
					<button id={'initupload{{lowercase .Object.Name}}'} onClick={updateTabEvent} className="flex flex-col justify-center items-center m-2 cursor-pointer" style={buttonStyle}>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6"  style={ {pointerEvents:"none"} }>
						<path strokeLinecap="round" strokeLinejoin="round" d="M7.5 7.5h-.75A2.25 2.25 0 0 0 4.5 9.75v7.5a2.25 2.25 0 0 0 2.25 2.25h7.5a2.25 2.25 0 0 0 2.25-2.25v-7.5a2.25 2.25 0 0 0-2.25-2.25h-.75m0-3-3-3m0 0-3 3m3-3v11.25m6-2.25h.75a2.25 2.25 0 0 1 2.25 2.25v7.5a2.25 2.25 0 0 1-2.25 2.25h-7.5a2.25 2.25 0 0 1-2.25-2.25v-.75" />
						</svg>
					</button>
					{{end}}
				</div>
				<AI subject={subject} updateList={updateList} collection="{{lowercase .Object.Name}}s"/>
			</div>
			{
				!promptToggle && <Loading/>
			}
			{
				promptToggle && <{{titlecase .Object.Name}}List title="{{titlecase .Object.Plural}}" subject={subject} native={true} />
			}
		</div>
	)
}