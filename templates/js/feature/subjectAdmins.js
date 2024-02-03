import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import AutocompleteUsername from '@/inputs/autocompleteUsername'
import Spacer from '@/inputs/spacer'
import VisitTab from '@/features/interfaces'

import { {{titlecase .Object.Name}}ObjectGET, {{titlecase .Object.Name}}AdminPOST } from './_fetch'

import { {{titlecase .Object.Name}}Admin } from './{{lowercase .Object.Name}}Admin';

export function {{titlecase .Object.Name}}Admins(props) {

    const [ userdata, setUserdata] = useUserContext()
    const [ localdata, setLocaldata] = useLocalContext()

    const [project, setProject] = useState(localdata.tab.context.object)

	const [newAdmins, setNewAdmins] = useState()

    console.log("FEATURES >> PROJECTS >> ADMINS", localdata)

	function updateProject() {
		ProjectObjectGET(userdata, project.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setProject(data)
		})
		.catch((e) => {
			console.error(e)
		})
	}

    // update tabs handles the updated context and sends the user to a new interface
    function updateTabEvent(e) {
        const id = e.target.id
        console.log("SELECT PROJECT", id)
        const next = projects[id.split("_")[1]]
        const context = {
            "_": next.name,
            "object": next
        }
        setLocaldata(VisitTab(localdata, "project", context))
    }

	function deleteAdmin(id) {
		const adminID = project.Meta.Moderation.Admins[id]
		{{titlecase .Object.Name}}AdminPOST(userdata, project.Meta.ID, "remove", adminID)
		.then(updateProject)
	}

	function inputChange(obj) {
		console.log("!!!", obj)
		setNewAdmins(obj.value)
	}

	function addAdmins() {
		newAdmins.forEach(function (admin, i) {
			{{titlecase .Object.Name}}AdminPOST(userdata, project.Meta.ID, "add", admin)
			.then(updateProject)
		})
	}

    return (
		<div style={ {padding:"30px 60px 30px 60px"} } className='flex flex-col'>
			<div className='text-xl'>Add Admin</div>
			<AutocompleteUsername inputChange={inputChange} />
			<Spacer/>
			<div>
			{
				newAdmins && <button onClick={addAdmins} className="my-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
					<div className='flex flex-row'>
						<div className='flex flex-col justify-center'>
							<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
							<path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
							</svg>
						</div>
						<div className='flex flex-col justify-center'>
							<div className='text-lg'>Add { newAdmins.join(" & ") }</div>
						</div>
					</div>
				</button>
			}
			</div>
			<Spacer/>
			<div className='text-xl'>Existing Administrators:</div>
			<Spacer/>
			<div className='flex flex-col'>
				{
					project.Meta.Moderation.Admins.map(function (adminID, i) {
						return (
							<{{titlecase .Object.Name}}Admin key={adminID} id={i} admin={adminID} delete={deleteAdmin}/>
						)
					})
				}
			</div>
		</div>
    )

}


