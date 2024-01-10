import { useState } from 'react';
import Calendar from 'react-calendar'
import 'react-calendar/dist/Calendar.css';

export default function DatePicker(props) {

    const [date, setDate] = useState(new Date())

    function change(date) {
        console.log(date)
        props.inputChange(
			{
				"id": props.id,
				"type": "date",
				"value": date,
				"required": props.required,
			}
        )
    }

    return (
        <Calendar onChange={change}/>
    )

}