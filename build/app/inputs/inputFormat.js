export default function InputFormat(inputs) {
    var payload = {}
    for (var field in inputs) {
        const v = inputs[field]
        payload[v.id] = v.value
    }
    return payload
}
  