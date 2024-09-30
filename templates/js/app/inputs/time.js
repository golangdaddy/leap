export default function Time(props) {

    console.log("SHOW INPUT", props)
  
    function changeEventOnload(e) {
      const id = e.target.id
      var value = e.target.value
      if (props.type == "number") {
        value = parseFloat(value)
      }
      const data = {
        "id": id,
        "ftype": props.ftype,
        "value": value,
        "required": props.required,
      }
      console.log("ONLOAD", data)
      props.inputChange(
        data
      )
    }

    function formatTime() {
        var hours = document.getElementById("hour").value;
        var minutes = document.getElementById("minute").value;
        props.inputChange(
            {
              "id": props.id,
              "ftype": props.ftype,
              "value": hours+":"+minutes,
              "required": props.required,
            }
          )
    }
  
    return (
      <>
        <div className="flex flex-col">
            <div className="text-l font-bold">{props.title}{props.required && "*"}</div>
            <div className="m-2"></div>
            <div className="flex flex-row">
                <input min="0" max="23" className="py-2 px-4 border" id="hour" type="number" defaultValue={props.value} onChange={formatTime} onLoad={changeEventOnload} placeholder="Hours" />
                <div className="flex flex-col justify-center">
                    <div className="p-2 font-bold">:</div>
                </div>
                <input min="0" max="59" className="py-2 px-4 border" id="minute" type="number" defaultValue={props.value} onChange={formatTime} onLoad={changeEventOnload} placeholder="Mins" />
            </div>
        </div>
      </>
    );
  }
  