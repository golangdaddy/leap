import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import { GoBack } from '../interfaces'
import Loading from '@/app/loading'

{{range .Object.Children}}import { {{titlecase .Name}}List } from '@/features/{{lowercase .Name}}s/shared/{{lowercase .Name}}List'
{{end}}

import { {{titlecase .Object.Name}}ObjectGET } from './_fetch'

export function {{titlecase .Object.Name}}(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [jdata, setJdata] = useState(localdata.tab.context.object)
    const [subject, setSubject] = useState(localdata.tab.context.object)
    const [image, setImage] = useState()

	function getObject() {
		{{titlecase .Object.Name}}ObjectGET(userdata, subject.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setSubject(data)
			setJdata(JSON.stringify(data.fields))
			if (data.Meta.URIs?.length > 0) {
				setImage("https://storage.googleapis.com/{{.DatabaseID}}-uploads/" + data.Meta.URIs[data.Meta.URIs.length - 1])
			}
			console.log("IMAGE? src:", image)
		}) 
		.catch((e) => {
            console.error(e)
			setLocaldata(GoBack(localdata))
        })
	}

	useEffect(() => {
		getObject()
	}, [])

    return (
        <>
			{ !subject && <Loading/> }
			{
				subject && <div>
					<div className='text-2xl'>{ subject.fields.name }</div>
					<div className='flex flex-row'>
						{
							image && <div className="m-4" style={ {maxWidth:"40vw"} }>
								<img className='w-full' src={image}/>
							</div>
						}
						<div>
							<table className='m-4'>
								<tbody>{{range .Object.Fields}}
									<tr>
										<td className='font-bold'>{{.Name}}</td>
										<td><div className='px-4'></div></td>
										<td>{ subject.fields["{{lowercase .Name}}"] }</td>
									</tr>
								{{end}}</tbody>
							</table>
						</div>
					</div>
				</div>
			}
            {{range .Object.Children}}
			<{{titlecase .Name}}List title="{{titlecase .Name}}" subject={subject} limit={4} />
			{{end}}
        </>
    )

}