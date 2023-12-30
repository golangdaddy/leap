import { PublicFetch } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function AttributesInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/attributes?function=init&parent="+parentID, p)
}

export function AttributeUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/attribute?function=update&id="+id, p)
}

export function AttributeObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/attribute?function=object&id="+id)
}

export function AttributesListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/attributes?function=list&parent="+parentID+"&limit="+limit)
}

export function AttributesCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/attributes?function=count&collection="+collectionID)
}

export function AttributeDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/attribute?id="+id)
}

export function AttributeFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/attribute?function="+func+"&id="+id)
}