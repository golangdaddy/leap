import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'
import { format, formatDistance, formatRelative, subDays } from 'date-fns'

import VisitTab from '../interfaces'
import { GetInterfaces } from '@/features/interfaces'
import { GoBack } from '../interfaces'
import Loading from '@/app/loading'
import Spacer from '@/inputs/spacer';

{{range .Object.Children}}import { {{titlecase .Name}}List } from '@/features/{{lowercase .Name}}s/shared/{{lowercase .Name}}List'
{{end}}

import { {{titlecase .Object.Name}}ObjectGET, {{titlecase .Object.Name}}JobPOST } from './_fetch'

export function {{titlecase .Object.Name}}(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [jdata, setJdata] = useState(localdata.tab.context.object)
    const [subject, setSubject] = useState(localdata.tab.context.object)
    const [image, setImage] = useState()

	const interfaces = GetInterfaces()

	var date = new Date(subject.Meta.Modified * 1000);
	const dateTime = formatRelative(date, new Date())

	// update tabs handles the updated context and sends the user to a new interface
	function editData() {
		setLocaldata(VisitTab(localdata, "edit{{lowercase .Object.Name}}", localdata.tab.context))
	}

	function getObject() {
		{{titlecase .Object.Name}}ObjectGET(userdata, subject.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setSubject(data)
			setJdata(JSON.stringify(data.fields))
			{{if .Object.Options.Image}}
			if (data.Meta.URIs?.length > 0) {
				setImage("https://storage.googleapis.com/{{.DatabaseID}}-uploads/" + data.Meta.URIs[data.Meta.URIs.length - 1])
			}{{end}}
			console.log("IMAGE? src:", image)
		}) 
		.catch((e) => {
            console.error(e)
			setLocaldata(GoBack(localdata))
        })
	}

	// update tabs handles the updated context and sends the user to a new interface
	function updateTabEvent(e) {
		console.log("UPDATE TAB EVENT:", e.target.id)
		updateTab(e.target.id)
	}
	function updateTab(tabname) {
		setLocaldata(VisitTab(localdata, tabname, localdata?.tab?.context))
	}

	useEffect(() => {
		getObject()
	}, [])

	const editButtonStyle = {
		borderRadius: "20px",
		backgroundColor: "rgb(52, 211, 153)",
		border: "solid 0px",
		color: "white",
		padding: "6px 10px"
	}

    return (
        <div style={ {padding:"30px 60px 30px 60px"} }>
			{ !subject && <Loading/> }
			{
				subject && <div className='flex flex-col w-full'>
					<div className='flex flex-row justify-between items-center w-full py-4 my-4'>
						<div className='text-base'>
							<span className='uppercase text-sm'>{ subject.Meta.Class }</span> / <span className='font-bold'>{ subject.fields.name }</span>
						</div>
						{
							localdata.tab.subsublinks.map(function (tabname, i) {
								if (tabname.length == 0) { return }
								const tab = interfaces[tabname]
								return (
									<div key={i} className='flex flex-row rounded-md border py-1 px-2 mx-2 bg-white'>
										<div id={tab.id} className='cursor-pointer text-gray-800' onClick={updateTabEvent}>{tab.name}</div>
									</div>
								)
							})
						}
						{{if eq false .Object.Options.ReadOnly}}
						<div className='flex flex-row justify-center items-center'>
							<button className='text-sm' onClick={editData} style={editButtonStyle}>
							Edit {{titlecase .Object.Name}}
							</button>
						</div>
						{{end}}
					</div>
					<hr/>
					<div className='flex flex-row'>
						{
							image && <div className="m-4" style={ {maxWidth:"40vw"} }>
								<img className='w-full' src={image}/>
							</div>
						}
						<div>
							<table className='m-4 w-full'>
								<tbody>
									<tr className='flex flex-row text-sm'>
										<td className='flex flex-col justify-start'>
											<div className='w-full flex flex-row justify-end'>
												<div className=''>Updated</div>
											</div>
										</td>
										<td className='flex flex-col justify-start'>
										</td>
										<td className='flex flex-col justify-start'>
											<div className='w-full flex flex-row justify-end'>
												<div className=''>{ dateTime }</div>
											</div>
										</td>
									</tr>
									<Spacer/>
									{{range .Object.Fields}}
									<tr className='flex flex-row'>
										<td className='flex flex-col justify-start'>
											<div className='w-full flex flex-row justify-end'>
												<div className='font-bold'>{{.Name}}</div>
											</div>
										</td>
										<td className='flex flex-col justify-start'>
											<div className='w-full flex flex-row justify-end'>
												<div className='px-2'>:</div>
											</div>
										</td>
										<td className='flex flex-col justify-start'>
											<div className='w-full flex flex-row justify-end'>
												<div className=''>{ subject.fields["{{lowercase .Name}}"] }</div>
											</div>
										</td>
									</tr>
									<Spacer/>
								{{end}}</tbody>
							</table>
						</div>
					</div>
				</div>
			}



			{{range .Object.Children}}
			{{if .Options.Job}}
			<{{titlecase .Name}}List title="{{titlecase .Name}}" subject={subject} limit={4} />
			{{end}}
			{{end}}

            {{range .Object.Children}}
			{{if eq false .Options.Job}}
			<{{titlecase .Name}}List title="{{titlecase .Name}}" subject={subject} limit={4} />
			{{end}}
			{{end}}
        </div>
    )

}