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
			value: subject.name,
			required: true,
		},"min": {
			id: "min",
			type: "int",
			value: subject.min,
			required: true,
		},"max": {
			id: "max",
			type: "int",
			value: subject.max,
			required: true,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="attribute name" placeholder="attribute name" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Input id="min" type='number' required={ true } title="attribute min" placeholder="attribute min" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Input id="max" type='number' required={ true } title="attribute max" placeholder="attribute max" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name","min","max"]}/>
			<Spacer/>
			
		</div>
	);
}
