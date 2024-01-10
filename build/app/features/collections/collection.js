import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import { GoBack } from '../interfaces'
import VisitTab from '@/features/interfaces'

import Loading from '@/app/loading'

import { AttributeList } from '@/features/attributes/shared/attributeList'
import { LayerList } from '@/features/layers/shared/layerList'


import { CollectionObjectGET } from './_fetch'

export function Collection(props) {  

    const [userdata, setUserdata] = useUserContext()
    const [localdata, setLocaldata] = useLocalContext() 

    const [jdata, setJdata] = useState(localdata.tab.context.object)
    const [subject, setSubject] = useState(localdata.tab.context.object)
	function getObject() {
		CollectionObjectGET(userdata, subject.Meta.ID)
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
            
			<AttributeList title="Attribute" subject={subject} />
			
			<LayerList title="Layer" subject={subject} />
			
        </>
    )

}