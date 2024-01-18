import * as React from 'react'
import { useUserContext } from '@/context/user'
import { useLocalContext } from '@/context/local'
import { useState, useEffect } from 'react'

import VisitTab from '../interfaces'

import { GoBack } from '../interfaces'
import Loading from '@/app/loading'

import { {{titlecase .Object.Name}}Assets } from './shared/{{lowercase .Object.Name}}Assets'

import { AssetsUser, AssetsWallet } from '@/app/fetch'
import { {{titlecase .Object.Name}}ObjectGET } from './_fetch'

export function Assets(props) {  

	const [userdata, setUserdata] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext() 
	const [subject, setSubject] = useState(localdata.tab.context.object)
	const [assets, setAssets] = useState()

	function getObject() {
		{{titlecase .Object.Name}}ObjectGET(userdata, subject.Meta.ID)
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

	function getAssets() {
		AssetsWallet(userdata, subject.Meta.ID)
		.then((res) => res.json())
		.then((data) => {
			console.log(data)
			setAssets(data)
		}) 
		.catch((e) => {
			console.error(e)
			setLocaldata(GoBack(localdata))
		})
	}

	useEffect(() => {
		getObject()
		getAssets()
	}, [])

	return (
		<div style={ {padding:"30px 60px 30px 60px"} }>
			<{{titlecase .Object.Name}}Assets subject={subject}/>
		</div>
	)

}