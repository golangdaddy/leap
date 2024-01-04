import * as React from 'react'

import { Attributes } from './attributes'
import { Attribute } from './attribute'
import { NewAttribute } from './newAttribute'
import { EditAttribute } from './editAttribute'
import { UploadAttribute } from './uploadAttribute'

export var AttributeInterfaces = {
    "newattribute": {
        level: 2,
        name: "New Attribute", 
        component: (<NewAttribute />),
    },  
    "editattribute": {
        level: 3,
        name: "Edit Attribute", 
        component: (<EditAttribute />),
    },  
    "uploadattribute": {
        level: 3,
        name: "Upload File", 
        component: (<UploadAttribute />),
    },  
    "attributes": {
        level: 1,
        name: "Manage Attributes", 
        component: (<Attributes />),
        subsublinks: ["newattribute"],
    },
    "attribute": {
        level: 2,
        name: "Attribute",
        sublinks: ["editattribute", "uploadattribute"],
        subsublinks: [""],
        component: (<Attribute />),
    },
}
