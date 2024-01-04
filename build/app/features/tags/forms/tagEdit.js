import { useState, useEffect } from 'react';
import { useUserContext } from '@/context/user';
import { useLocalContext } from '@/context/local';

import Spacer from '@/inputs/spacer';
import Select from '@/inputs/select';
import Submit from '@/inputs/submit';
import Input from '@/inputs/input';
import InputChange from '@/inputs/inputChange';

export function TagEdit(props) {

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
		},"foreground_color": {
			id: "foreground_color",
			type: "string",
			value: subject.foreground_color,
			required: false,
		},"background_color": {
			id: "background_color",
			type: "string",
			value: subject.background_color,
			required: false,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="tag name" placeholder="tag name" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Input id="foreground_color" type='text' required={ false } title="tag foreground_color" placeholder="tag foreground_color" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Input id="background_color" type='text' required={ false } title="tag background_color" placeholder="tag background_color" inputChange={handleInputChange}/>
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name"]}/>
			<Spacer/>
			
		</div>
	);
}
