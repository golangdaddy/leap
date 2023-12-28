import * as React from 'react'

import { Collections } from './collections'
import { Collection } from './collection'
import { NewCollection } from './newCollection'
import { EditCollection } from './editCollection'

export var CollectionInterfaces = {
    "newcollection": {
        level: -1,
        name: "New Collection", 
        component: (<NewCollection />),
    },  
    "editcollection": {
        level: -1,
        name: "Edit Collection", 
        component: (<EditCollection />),
    },  
    "collections": {
        level: 7,
        name: "Manage Collections", 
        component: (<Collections />),
    },
    "collection": {
        level: 8,
        name: "Collection", 
        component: (<Collection />),
    },
}
