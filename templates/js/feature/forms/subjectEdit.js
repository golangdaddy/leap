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
		"{{.ID}}": {
			id: "{{.ID}}",
			ftype: {{json .Object.Element}},
			{{if ne nil .Object.Element}}
				value: subject.fields.{{.ID}},
			{{else}}
				value: "",
			{{end}}
			required: {{.Required}},
		},
		{{end}}
	})

	console.log("NEWINPUTS", inputs)

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
