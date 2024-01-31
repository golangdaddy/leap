export default function Object(props) {

  console.log("SHOW TEXTAREA", props)

  function changeEvent(e) {
    const id = e.target.id
    const value = e.target.value
    props.inputChange(
      {
        "id": id,
        "type": "text",
        "value": value,
        "required": props.required,
      }
    )
  }

  return (
    <div className="flex flex-col">
      <div className="text-l font-bold">{props.title}{props.required && "*"}</div>
      <div className="m-2"></div>
  	  <textarea className="p-2 border" id={props.id} type={props.type} defaultValue={JSON.stringify(props.value)} onChange={changeEvent} placeholder={props.placeholder} style={{"height":"25vh"}}></textarea>
    </div>
  );
}
