import { useState } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import { GoBack } from '@/features/interfaces';

import { OverlayssInitPOST } from './_fetch';

import { OverlaysForm } from './forms/overlays';
import { titlecase } from './_interfaces';

export function NewOverlays(props) {

  const [userdata, _] = useUserContext()
  const [localdata, setLocaldata] = useLocalContext()

  const [subject] = useState(localdata.tab.context.object)

	function submitNew(inputs) {
		OverlayssInitPOST(
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
		<OverlaysForm submit={submitNew}/>
	);
}
