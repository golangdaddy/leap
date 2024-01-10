import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import { GoBack } from '@/features/interfaces';

import { TagObjectGET, TagUpdatePOST } from './_fetch';

import { TagEdit } from './forms/tagEdit';

export function EditTag(props) {

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()

    // make sure the object is current
    const [subject, setSubject] = useState(localdata.tab.context.object)
    useEffect(() => {
        TagObjectGET(userdata, subject?.Meta.ID)
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
		TagUpdatePOST(
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
			subject && <TagEdit subject={subject} submit={submitEdit}/>
		}
		</>
  	);
}
