import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Select from '@/inputs/select';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import InputChange from '@/inputs/inputChange';

export function ElementEdit(props) {

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
		},"max_mint": {
			id: "max_mint",
			type: "int",
			value: subject.fields.max_mint,
			required: true,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="element name" placeholder="element name" inputChange={handleInputChange} value={ inputs["name"].value } />
			<Spacer/>
			
			<Input id="max_mint" type='number' required={ true } title="element max_mint" placeholder="element max_mint" inputChange={handleInputChange} value={ inputs["max_mint"].value } />
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name","max_mint"]}/>
			<Spacer/>
			
		</div>
	);
}
