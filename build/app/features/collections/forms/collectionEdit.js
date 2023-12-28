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

export function CollectionEdit(props) {

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
					<Input id="description" type='text' required={true} value={inputs.description.value} title="Collection Description" placeholder="Collection description..." inputChange={handleInputChange}/>
					<Spacer/>
					<Select id="type" required={true} type='text' title="Collection Type" options={["text", "attribute"]} value={inputs.type.value} inputChange={handleInputChange}/>
					<Spacer/>
					{
						attributes && inputs["type"] && (inputs["type"].value == "attribute") && <>
							<Select id="content" required={true} type='text' title="Collection attribute" value={inputs.content.value} options={attributes} inputChange={handleInputChange}/>
							<Spacer/>
						</>
					}
					{
						inputs["type"] && (inputs["type"].value == "text") && <>
							<Input id="content" required={true} title="Collection Value" type="text" value={inputs.content.value} placeholder="Collection value..." inputChange={handleInputChange} />
							<Spacer/>
						</>
					}
					<Input id="x" required={true} title="Collection X Position" type="number" value={inputs.x.value} placeholder="Collection X..." inputChange={handleInputChange} />
					<Spacer/>
					<Input id="y" required={true} title="Collection Y Position" type="number" value={inputs.y.value} placeholder="Collection Y..." inputChange={handleInputChange} />
					<Spacer/>
					{
						fonts && <>
							<Select id="font" required={true} type='text' title="Collection Font" options={fonts} value={inputs.font.value} inputChange={handleInputChange}/>
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
