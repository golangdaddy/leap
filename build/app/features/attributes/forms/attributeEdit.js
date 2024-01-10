import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Select from '@/inputs/select';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import InputChange from '@/inputs/inputChange';

export function AttributeEdit(props) {

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
		},"min": {
			id: "min",
			type: "int",
			value: subject.fields.min,
			required: true,
		},"max": {
			id: "max",
			type: "int",
			value: subject.fields.max,
			required: true,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="attribute name" placeholder="attribute name" inputChange={handleInputChange} value={ inputs["name"].value } />
			<Spacer/>
			
			<Input id="min" type='number' required={ true } title="attribute min" placeholder="attribute min" inputChange={handleInputChange} value={ inputs["min"].value } />
			<Spacer/>
			
			<Input id="max" type='number' required={ true } title="attribute max" placeholder="attribute max" inputChange={handleInputChange} value={ inputs["max"].value } />
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name","min","max"]}/>
			<Spacer/>
			
		</div>
	);
}
