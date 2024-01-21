import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import { AI } from './forms/ai';
import { {{titlecase .Object.Name}}List } from './shared/{{lowercase .Object.Name}}List';

export function {{titlecase .Object.Name}}s(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ parent ] = useState(localdata.tab.context.object)

	return (
		<div style={ {padding:"30px 60px 30px 60px"} }>
			<AI parent={parent}/>
			<{{titlecase .Object.Name}}List subject={parent} />
		</div>
	)
}