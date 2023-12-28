import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import VisitTab from '@/features/interfaces'

import Loading from '@/app/loading'
import { Preview } from './shared/collection'

import { CollectionObjectGET } from './_fetch'

export function Collection(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [subject, setSubject] = useState(localdata.tab.context.object)
	function getObject() {
		CollectionObjectGET(userdata, subject.Meta.ID)
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
            {
                subject && <Preview subject={subject} />
            }
        </>
    )

}