import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import VisitTab from '@/features/interfaces'

import Loading from '@/app/loading'

import { ProjectObjectGET } from './_fetch'

export function Project(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [subject, setSubject] = useState(localdata.tab.context.object)
	function getObject() {
		ProjectObjectGET(userdata, subject.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setSubject(data)
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
                subject && <textarea className='w-full'>{JSON.stringify(subject.fields)}</textarea>
            }
        </>
    )

}