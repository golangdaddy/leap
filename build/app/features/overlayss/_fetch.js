import { PublicFetch, AxiosPOST } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function OverlayssInitPOST(user, parentID, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/overlayss?function=init&parent="+parentID, p)
}

export function OverlaysUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/overlays?function=update&id="+id, p)
}

export function OverlaysObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/overlays?function=object&id="+id)
}

export function OverlayssListGET(user, parentID, limit) {
    return SessionFetch(user, "GET", "api/overlayss?function=list&parent="+parentID+"&limit="+limit)
}

export function OverlayssCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/overlayss?function=count&collection="+collectionID)
}

export function OverlaysDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/overlays?id="+id)
}

export function OverlaysFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/overlays?function="+func+"&id="+id)
}

export function OverlaysFileUpload(user, id, formData) {
    return AxiosPOST(user, "api/overlays?function=upload&id="+id, formData)
}