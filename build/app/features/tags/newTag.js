import { useState } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import { GoBack } from '@/features/interfaces';

import { TagsInitPOST } from './_fetch';

import { TagForm } from './forms/tag';
import { titlecase } from './_interfaces';

export function NewTag(props) {

  const [userdata, _] = useUserContext()
  const [localdata, setLocaldata] = useLocalContext()

  const [subject] = useState(localdata.tab.context.object)

	function submitNew(inputs) {
		TagsInitPOST(
			userdata,
			subject?.Meta.ID,
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
		<TagForm submit={submitNew}/>
	);
}
