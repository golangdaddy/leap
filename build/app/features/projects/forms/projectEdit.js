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

export function ProjectEdit(props) {

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
		},"description": {
			id: "description",
			type: "string",
			value: subject.description,
			required: true,
		},
	})
	function handleInputChange(obj) {
		InputChange(inputs, setInputs, obj)
	}

  return (
    <div className='flex flex-col'>
		{  
			<>
					<Input id="description" type='text' required={true} value={inputs.description.value} title="Project Description" placeholder="Project description..." inputChange={handleInputChange}/>
					<Spacer/>
					<Select id="type" required={true} type='text' title="Project Type" options={["text", "attribute"]} value={inputs.type.value} inputChange={handleInputChange}/>
					<Spacer/>
					{
						attributes && inputs["type"] && (inputs["type"].value == "attribute") && <>
							<Select id="content" required={true} type='text' title="Project attribute" value={inputs.content.value} options={attributes} inputChange={handleInputChange}/>
							<Spacer/>
						</>
					}
					{
						inputs["type"] && (inputs["type"].value == "text") && <>
							<Input id="content" required={true} title="Project Value" type="text" value={inputs.content.value} placeholder="Project value..." inputChange={handleInputChange} />
							<Spacer/>
						</>
					}
					<Input id="x" required={true} title="Project X Position" type="number" value={inputs.x.value} placeholder="Project X..." inputChange={handleInputChange} />
					<Spacer/>
					<Input id="y" required={true} title="Project Y Position" type="number" value={inputs.y.value} placeholder="Project Y..." inputChange={handleInputChange} />
					<Spacer/>
					{
						fonts && <>
							<Select id="font" required={true} type='text' title="Project Font" options={fonts} value={inputs.font.value} inputChange={handleInputChange}/>
							<Spacer/>
						</>
					}
					<Input id="fontSize" required={true} title="Font Size" type="number"  value={inputs.fontSize.value} inputChange={handleInputChange} />
					<Spacer/>
					<Submit text="Update" inputs={inputs} submit={props.submit} assert={["description", "type", "content", "x", "y", "font", "fontSize"]}/>
			</>
		}
    </div>
  );
}
