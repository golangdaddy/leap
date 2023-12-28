import * as React from 'react'

import { useLocalContext } from '@/context/local'
import VisitTab from '@/app/interfaces'

export default function Username(props) {

	const [localdata, setLocaldata] = useLocalContext()

	function visitProfile() {
		const context = {
			"_": props.user.username,
			"object": props.user
		}
		setLocaldata(VisitTab(localdata, "profile", context))
	}

	return (
		<div className='font-bold cursor-pointer uppercase' onClick={visitProfile}>{props.user.Username}</div>
	)
}
