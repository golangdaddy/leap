import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import { {{titlecase .Object.Name}}DELETE, {{titlecase .Object.Name}}sListGET } from '../_fetch';

import VisitTab from '@/features/interfaces';

import Loading from '@/app/loading';
import { Preview } from './{{lowercase .Object.Name}}';
import Spacer from '@/inputs/spacer';

export function {{titlecase .Object.Name}}List(props) {

    // set props.limit if you want to limit query results

    const [ userdata, setUserdata] = useUserContext()
    const [ localdata, setLocaldata] = useLocalContext()

    const [ list, setList ] = useState(null)

    function updateList() {
        {{titlecase .Object.Name}}sListGET(userdata, props.subject?.Meta.ID, props.limit)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setList(data)
		})
    }

    useEffect(() => {
		updateList()
    }, [])

    // update tabs handles the updated context and sends the user to a new interface
    function selectItem(id) {
        console.log("SELECT {{titlecase .Object.Name}}", id)
        const next = overlays[parseInt(id)]
        const context = {
            "_": next.name,
            "object": next,
            "deleteFunction": {{titlecase .Object.Name}}DELETE,
        }
        setLocaldata(VisitTab(localdata, "{{lowercase .Object.Name}}", context))
    }

    function deleteItem(id) {
        const object = overlays[parseInt(id)]
        console.log("DELETING", object)
        {{titlecase .Object.Name}}DELETE(userdata, object.Meta.ID)
        .then((res) => console.log(res))
        .then(function () {
            updateList()
        })
        .catch(function (e) {
            console.error("FAILED TO SEND", e)
        })
    }

    return (
    <div className='flex flex-col my-4'>
    {
        props.title && <div className='py-4 my-4 text-2xl font-bold'>{props.title}:</div>
    }
    {
        !list && <Loading/>
    }
    {
        list && list.map(function (item, i) {

            return (
                <div key={i}>
                    <Preview id={i} listLength={list.length} item={item} select={selectItem} delete={deleteItem}/>
                    <Spacer/>
                </div>
            )
        })
    }
    </div>
    )

}