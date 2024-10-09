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

	{{- $obj := .Object }}
	const [inputs, setInputs] = useState({
		{{range $nn, $field := $obj.Fields}}
		{{if ne nil $field.Element}}
			{{if more $nn 0}},{{end}}
			"{{$field.ID}}": {
				id: "{{$field.ID}}",
				ftype: {{json $field.Element}},
				value: subject.fields.{{$field.ID}},
				required: {{.Required}},
			}
		{{else}}
			{{range $n, $f := $field.Inputs}} 
				{{if more $n 0}},{{else}}{{if more $nn 0}},{{end}}{{end}} 
				"{{$f.ID}}": {
					id: "{{$f.ID}}",
					ftype: {{json $f.Element}},
					value: subject.fields.{{$f.ID}},
					required: {{.Required}},
				}
			{{end}}
		{{end}}
		{{end}}
	});
	

	console.log("NEWINPUTS", inputs)

	return (
		<div className='flex flex-col'>
			{{range .EditInputs}}
			{{.}}
			<Spacer/>
			{{end}}
		</div>
	);
}
