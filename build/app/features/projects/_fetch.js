import { PublicFetch } from '@/app/fetch';
import SessionFetch from '@/app/fetch';
import InputFormat from '@/inputs/inputFormat';

export function ProjectsInitPOST(user, layer, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/projects?function=init&layer="+layer, p)
}

export function ProjectUpdatePOST(user, id, payload) {
    var p = InputFormat(payload)
    return SessionFetch(user, "POST", "api/project?function=update&id="+id, p)
}

export function ProjectObjectGET(user, id) {
    return SessionFetch(user, "GET", "api/project?function=object&id="+id)
}

export function ProjectsListGET(user, layer, limit) {
    return SessionFetch(user, "GET", "api/projects?function=list&layer="+layer+"&limit="+limit)
}

export function ProjectsCountGET(user, collectionID) {
    return SessionFetch(user, "GET", "api/projects?function=count&collection="+collectionID)
}

export function ProjectDELETE(user, id) {
    return SessionFetch(user, "DELETE", "api/project?id="+id)
}

export function ProjectFunctionPOST(user, id, func) {
    return SessionFetch(user, "POST", "api/project?function="+func+"&id="+id)
}