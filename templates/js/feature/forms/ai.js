import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import Textarea from '@/inputs/textarea';
import InputChange from '@/inputs/inputChange';
import Checkbox from '@/inputs/checkbox';
import Select from '@/inputs/select';
import CollectionSelect from '@/inputs/collectionSelect';
import Color from '@/inputs/color';

import { {{titlecase .Object.Name}}ChatGPTPOST } from '../_fetch'

export function AI(props) {

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

	console.log("AI SUBJECT", props.subject)

	const [toggle, setToggle] = useState(false)

	function toggleState() {
		setToggle(!toggle)
	}

	const [select, setSelect] = useState('create')
	function updateSelect(e) {
		setSelect(e.target.id)
	}

	function sendPrompt() {
		props.updateList(false)
		const payload = {
			"prompt": document.getElementById("prompt").value,
		}
		switch (select) {
		case "prompt":
			{{titlecase .Object.Name}}ChatGPTPOST(userdata, props.subject.Meta.ID, "create", payload)
			.then((res) => {
				console.log(res)
				props.updateList(true)
			}) 
			.catch((e) => {
				console.error(e)
				props.updateList(true)
			})
			break
		case "create":
			{{titlecase .Object.Name}}ChatGPTPOST(userdata, props.subject.Meta.ID, "create", payload)
			.then((res) => {
				console.log(res)
				props.updateList(true)
			}) 
			.catch((e) => {
				console.error(e)
				props.updateList(true)
			})
			break
		case "modify":
			{{titlecase .Object.Name}}ChatGPTPOST(userdata, props.subject.Meta.ID, props.collection, "modify", payload)
			.then((res) => {
				console.log(res)
				props.updateList(true)
			}) 
			.catch((e) => {
				console.error(e)
				props.updateList(true)
			})
			break
		}
	}

	return (
		<div className='flex flex-col'>
			{
				!toggle && <div className="flex flex-col justify-center rounded-l-lg bg-gray-400" onClick={toggleState}>
					<div id="home" className="flex flex-col justify-center items-center m-4 cursor-pointer" style={ {width:"36px",height:"36px"} }>
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6" style={ {pointerEvents:"none"} }>
						<path strokeLinecap="round" strokeLinejoin="round" d="m21 7.5-2.25-1.313M21 7.5v2.25m0-2.25-2.25 1.313M3 7.5l2.25-1.313M3 7.5l2.25 1.313M3 7.5v2.25m9 3 2.25-1.313M12 12.75l-2.25-1.313M12 12.75V15m0 6.75 2.25-1.313M12 21.75V19.5m0 2.25-2.25-1.313m0-16.875L12 2.25l2.25 1.313M21 14.25v2.25l-2.25 1.313m-13.5 0L3 16.5v-2.25" />
						</svg>
					</div>
				</div>
			}
			{
				toggle && <>
					<div className='flex flex-row'>
						<button id="prompt" onClick={updateSelect} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Prompt</button>
						<button id="create"  onClick={updateSelect} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Create</button>
						<button id="modify" onClick={updateSelect} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Modify</button>
					</div>
					<textarea id='prompt' placeholder={"your "+select+" prompt..."} className='border p-2'></textarea>
					<div>
						<button onClick={sendPrompt} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Send</button>
					</div>
				</>
			}
		</div>
	);
}
