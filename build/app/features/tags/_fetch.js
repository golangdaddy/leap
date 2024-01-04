import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function TagsInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/tags?function=init&parent="+parentID, p)
}

export function TagUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/tag?function=update&id="+id, p)
}

export function TagObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/tag?function=object&id="+id)
}

export function TagsListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/tags?function=list&parent="+parentID+"&limit="+limit)
}

export function TagsCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/tags?function=count&collection="+collectionID)
}

export function TagDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/tag?id="+id)
}

export function TagFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/tag?function="+func+"&id="+id)
}

export function TagFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/tag?function=upload&id="+id, formData)
}