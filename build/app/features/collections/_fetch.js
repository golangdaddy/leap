import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function CollectionsInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/collections?function=init&parent="+parentID, p)
}

export function CollectionUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/collection?function=update&id="+id, p)
}

export function CollectionObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/collection?function=object&id="+id)
}

export function CollectionsListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/collections?function=list&parent="+parentID+"&limit="+limit)
}

export function CollectionsCountGET(user, parentID) {
    return SessionFetch(user, "GET", "api/collections?function=count&parent="+parentID)
}

export function CollectionMoveUpPOST(user, id) {
    return SessionFetch(user, "POST", "api/collection?function=up&id="+id)
}

export function CollectionMoveDownPOST(user, id) {
    return SessionFetch(user, "POST", "api/collection?function=down&id="+id)
}

export function CollectionDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/collection?id="+id)
}

export function CollectionFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/collection?function="+func+"&id="+id)
}

export function CollectionFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/collection?function=upload&id="+id, formData)
}
