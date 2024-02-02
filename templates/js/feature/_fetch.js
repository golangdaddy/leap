import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function {{titlecase .Object.Name}}sInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}s?function=init&parent="+parentID, p)
}

export function {{titlecase .Object.Name}}UpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=update&id="+id, p)
}

export function {{titlecase .Object.Name}}ObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/{{lowercase .Object.Name}}?function=object&id="+id)
}

export function {{titlecase .Object.Name}}sListGET(user, parentID, mode, limit) {
    return SessionFetch(user, "GET", "api/{{lowercase .Object.Name}}s?function=list&parent="+parentID+"&mode="+mode+"&limit="+limit)
}

export function {{titlecase .Object.Name}}sCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/{{lowercase .Object.Name}}s?function=count&parent="+parentID)
}

export function {{titlecase .Object.Name}}MoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=up&id="+id)
}

export function {{titlecase .Object.Name}}MoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=down&id="+id)
}

export function {{titlecase .Object.Name}}DELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/{{lowercase .Object.Name}}?id="+id)
}

export function {{titlecase .Object.Name}}FunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function="+func+"&id="+id)
}

export function {{titlecase .Object.Name}}JobPOST(user, id, job) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=job&job="+job+"&id="+id)
}

// file handling

export function {{titlecase .Object.Name}}Upload(user, id, mode, formData) {
    return AxiosPOST(user, "api/{{lowercase .Object.Name}}?function=upload&id="+id+"&mode="+mode, formData)
}

export function {{titlecase .Object.Name}}sUpload(user, parentID, mode, formData) {
    return AxiosPOST(user, "api/{{lowercase .Object.Name}}s?function=upload&parent="+parentID+"&mode="+mode, formData)
}

// misc

export function {{titlecase .Object.Name}}ChatGPTModifyPOST(user, parentID, collectionID, payload) {
    return SessionFetch(user, "POST", "api/openai?function=collectionprompt&collection="+collectionID+"&parent="+parentID, payload)
}

export function {{titlecase .Object.Name}}ChatGPTInitPOST(user, parentID, payload) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}s?function=prompt&parent="+parentID, payload)
}

export function {{titlecase .Object.Name}}ChatGPTPromptPOST(user, id, payload) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=prompt&id="+id, payload)
}

export function {{titlecase .Object.Name}}AdminPOST(user, id, mode, payload) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function=admin&mode="+mode+"&id="+id, payload)
}