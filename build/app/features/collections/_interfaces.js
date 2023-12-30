import * as React from 'react'

import { Collections } from './collections'
import { Collection } from './collection'
import { NewCollection } from './newCollection'
import { EditCollection } from './editCollection'

export var CollectionInterfaces = {
    "newcollection": {
        level: 2,
        name: "New Collection", 
        component: (<NewCollection />),
    },  
    "editcollection": {
        level: 3,
        name: "Edit Collection", 
        component: (<EditCollection />),
    },  
    "collections": {
        level: 1,
        name: "Manage Collections", 
        component: (<Collections />),
        subsublinks: ["newcollection"],
    },
    "collection": {
        level: 2,
        name: "Collection", 
        subsublinks: ["layers","attributes",""],
        component: (<Collection />),
    },
}
