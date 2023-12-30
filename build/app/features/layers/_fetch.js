import { PublicFetch } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function LayersInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/layers?function=init&parent="+parentID, p)
}

export function LayerUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/layer?function=update&id="+id, p)
}

export function LayerObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/layer?function=object&id="+id)
}

export function LayersListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/layers?function=list&parent="+parentID+"&limit="+limit)
}

export function LayersCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/layers?function=count&collection="+collectionID)
}

export function LayerDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/layer?id="+id)
}

export function LayerFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/layer?function="+func+"&id="+id)
}