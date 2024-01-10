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
			value: subject.fields.name,
			required: true,
		},"foreground_color": {
			id: "foreground_color",
			type: "string",
			value: subject.fields.foreground_color,
			required: false,
		},"background_color": {
			id: "background_color",
			type: "string",
			value: subject.fields.background_color,
			required: false,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

	return (
		<div className='flex flex-col'>
			
			<Input id="name" type='text' required={ true } title="tag name" placeholder="tag name" inputChange={handleInputChange} value={ inputs["name"].value } />
			<Spacer/>
			
			<Input id="foreground_color" type='text' required={ false } title="tag foreground_color" placeholder="tag foreground_color" inputChange={handleInputChange} value={ inputs["foreground_color"].value } />
			<Spacer/>
			
			<Input id="background_color" type='text' required={ false } title="tag background_color" placeholder="tag background_color" inputChange={handleInputChange} value={ inputs["background_color"].value } />
			<Spacer/>
			
			<Submit text="Save" inputs={inputs} submit={props.submit} assert={["name"]}/>
			<Spacer/>
			
		</div>
	);
}
