import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import { GoBack } from '../interfaces'
import VisitTab from '@/features/interfaces'

import Loading from '@/app/loading'

import { CollectionList } from '@/features/collections/shared/collectionList'
import { FontList } from '@/features/fonts/shared/fontList'


import { ProjectObjectGET } from './_fetch'

export function Project(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [jdata, setJdata] = useState(localdata.tab.context.object)
    const [subject, setSubject] = useState(localdata.tab.context.object)
	function getObject() {
		ProjectObjectGET(userdata, subject.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setSubject(data)
			setJdata(JSON.stringify(data.fields))
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
            
			<CollectionList title="Collection" subject={subject} />
			
			<FontList title="Font" subject={subject} />
			
        </>
    )

}