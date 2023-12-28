import { useState } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import { GoBack } from '@/features/interfaces';

import { ProjectsInitPOST } from './_fetch';

import { ProjectForm } from './forms/project';
import { titlecase } from './_interfaces';

export function NewProject(props) {

  const [userdata, _] = useUserContext()
  const [localdata, setLocaldata] = useLocalContext()

  const [subject] = useState(localdata.tab.context.object)

  function submitNew(inputs) {
	ProjectsInitPOST(
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
		<ProjectForm submit={submitNew}/>
	);
}
