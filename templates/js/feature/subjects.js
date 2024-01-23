import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import { AI } from './forms/ai';
import { {{titlecase .Object.Name}}List } from './shared/{{lowercase .Object.Name}}List';

export function {{titlecase .Object.Name}}s(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ subject ] = useState(localdata.tab.context.object)

	const [promptToggle, setPromptToggle] = useState(true)

	function updateList(state) {
		setPromptToggle(state)
	}

	return (
		<div style={ {padding:"30px 60px 30px 60px"} }>
			<AI subject={subject} updateList={updateList} collection="{{lowercase .Object.Name}}s"/>
			{
				promptToggle && <{{titlecase .Object.Name}}List subject={subject} />
			}
		</div>
	)
}