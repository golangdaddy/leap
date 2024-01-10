import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Select from '@/inputs/select';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import InputChange from '@/inputs/inputChange';

export function LayerEdit(props) {

	console.log("COLLECTION EDIT", props)

	const [userdata, _] = useUserContext()
	const [localdata, setLocaldata] = useLocalContext()
	const [subject] = useState(localdata.tab.context.object)

	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	const [inputs, setInputs] = useState({
		"name": {
			id: "name",
			type: "string",
			value: subject.fields.name,
			required: true,
		},"type": {
			id: "type",
			type: "string",
			value: subject.fields.type,
			required: true,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="layer name" placeholder="layer name" inputChange={handleInputChange} value={ inputs["name"].value } />
			<Spacer/>
			
			<Select id="type" type='text' required={ true } title="layer type" options={ ["foreground","element","background"] } placeholder="layer type" inputChange={handleInputChange} value={ inputs["type"].value } />
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name","type"]}/>
			<Spacer/>
			
		</div>
	);
}
