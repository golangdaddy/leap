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

export function {{titlecase .Object.Name}}sListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/{{lowercase .Object.Name}}s?function=list&parent="+parentID+"&limit="+limit)
}

export function {{titlecase .Object.Name}}sCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/{{lowercase .Object.Name}}s?function=count&collection="+collectionID)
}

export function {{titlecase .Object.Name}}DELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/{{lowercase .Object.Name}}?id="+id)
}

export function {{titlecase .Object.Name}}FunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/{{lowercase .Object.Name}}?function="+func+"&id="+id)
}

export function {{titlecase .Object.Name}}FileUpload(user, id, formData) {
    return AxiosPOST(user, "api/{{lowercase .Object.Name}}?function=upload&id="+id, formData)
}