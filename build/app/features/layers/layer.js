import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import { GoBack } from '../interfaces'
import VisitTab from '@/features/interfaces'

import Loading from '@/app/loading'

import { OverlayList } from '@/features/overlays/shared/overlayList'
import { ElementList } from '@/features/elements/shared/elementList'
import { TagList } from '@/features/tags/shared/tagList'


import { LayerObjectGET } from './_fetch'

export function Layer(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [jdata, setJdata] = useState(localdata.tab.context.object)
    const [subject, setSubject] = useState(localdata.tab.context.object)
	function getObject() {
		LayerObjectGET(userdata, subject.Meta.ID)
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
            
			<OverlayList title="Overlay" subject={subject} />
			
			<ElementList title="Element" subject={subject} />
			
			<TagList title="Tag" subject={subject} />
			
        </>
    )

}