import * as React from 'react'
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';
import { useState, useEffect } from 'react';

import { {{titlecase .Object.Name}}Matrix } from './shared/{{lowercase .Object.Name}}Matrix';

export function {{titlecase .Object.Name}}s(props) {

	const [ userdata, setUserdata] = useUserContext()
	const [ localdata, setLocaldata] = useLocalContext()

	const [ parent ] = useState(localdata.tab.context.object)

	return (
		<{{titlecase .Object.Name}}Matrix subject={parent} />
	)
}