export default function InputChange(inputs, setInputs, obj) {
    var newData = Object.assign({}, inputs)
    newData[obj.id] = obj
    setInputs(newData)
    console.log("NEW INPUT CHANGE", obj.id, inputs)
}
  