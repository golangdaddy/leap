import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import Color from '@/inputs/color';
import Checkbox from '@/inputs/checkbox';
import Select from '@/inputs/select';
import CollectionSelect from '@/inputs/collectionSelect';
import Object from '@/inputs/object';

import InputChange from '@/inputs/inputChange';


export function {{titlecase .Object.Name}}Edit(props) {

	console.error("COLLECTION EDIT", props)

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [subject] = useState(localdata.tab.context.object)

	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	const [inputs, setInputs] = useState({
		{{range .Object.Fields}}
		"{{lowercase .ID}}": {
			id: "{{lowercase .ID}}",
			ftype: {{"json .Element}}",
			{{if eq "name" (lowercase .ID)}}
			value: subject.Meta.Name,
			{{else}}
			value: subject.fields.{{lowercase .ID}},
			{{end}}
			required: {{.Required}},
		},
		{{end}}
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			{{range .EditInputs}}
			{{.}}
			<Spacer/>
			{{end}}
		</div>
	);
}
