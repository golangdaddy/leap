import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import { GoBack } from '@/features/interfaces';

import { {{titlecase .Object.Name}}ObjectGET, {{titlecase .Object.Name}}UpdatePOST } from './_fetch';

import { {{titlecase .Object.Name}}Edit } from './forms/{{lowercase .Object.Name}}Edit';

export function Edit{{titlecase .Object.Name}}(props) {

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

    // make sure the object is current
    const [subject, setSubject] = useState(localdata.tab.context.object)
    useEffect(() => {
        {{titlecase .Object.Name}}ObjectGET(userdata, subject?.Meta.ID)
        .then((res) => res.json())
		.then((data) => {
			console.log("UPDATED OBJECT",data)
			setSubject(data)
		})
		.catch((e) => {
			console.log(e)
			setLocaldata(GoBack(localdata))
		})
	}, [])

	function submitEdit(inputs) {
		{{titlecase .Object.Name}}UpdatePOST(
			userdata,
			subject.Meta.ID,
			inputs
		)
		.then((res) => console.log(res))
		.then(function () {
			// return to previous interface
			setLocaldata(GoBack(localdata))
		})
		.catch(function (e) {
			console.error("FAILED TO SEND", e)
		})
	}

	return (
		<>
		{
			subject && <{{titlecase .Object.Name}}Edit subject={subject} submit={submitEdit}/>
		}
		</>
  	);
}
