import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Select from '@/inputs/select';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import InputChange from '@/inputs/inputChange';

import { FontsGET } from '@/features/fonts/_fetch';
import { AttributesGET } from '@/features/attributes/_fetch';

export function {{titlecase .Object.Name}}Edit(props) {

	console.log("COLLECTION EDIT", props)

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [subject] = useState(localdata.tab.context.object)

	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	const [inputs, setInputs] = useState({
		{{range .Object.Fields}}"{{lowercase .Name}}": {
			id: "{{lowercase .Name}}",
			type: "{{lowercase .Type}}",
			value: subject.{{lowercase .Name}},
			required: {{.Required}},
		},{{end}}
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			{{range .Inputs}}
			{{.}}
			<Spacer/>
			{{end}}
		</div>
	);
}
