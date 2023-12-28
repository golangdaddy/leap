import { PublicFetch } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function CollectionsInitPOST(user, layer, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/collections?function=init&layer="+layer, p)
}

export function CollectionUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/collection?function=update&id="+id, p)
}

export function CollectionObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/collection?function=object&id="+id)
}

export function CollectionsListGET(user, layer, limit) {
    return SessionFetch(user, "GET", "api/collections?function=list&layer="+layer+"&limit="+limit)
}

export function CollectionsCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/collections?function=count&collection="+collectionID)
}

export function CollectionDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/collection?id="+id)
}

export function CollectionFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/collection?function="+func+"&id="+id)
}