import * as React from 'react'

import { Collections } from './collections'
import { Collection } from './collection'
import { NewCollection } from './newCollection'
import { EditCollection } from './editCollection'
import { UploadCollection } from './uploadCollection'

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
    "uploadcollection": {
        level: 3,
        name: "Upload File", 
        component: (<UploadCollection />),
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
        sublinks: ["editcollection", "uploadcollection"],
        subsublinks: ["attributes","layers",""],
        component: (<Collection />),
    },
}
