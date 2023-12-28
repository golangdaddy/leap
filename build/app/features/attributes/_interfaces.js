import * as React from 'react'

import { Attributes } from './attributes'
import { Attribute } from './attribute'
import { NewAttribute } from './newAttribute'
import { EditAttribute } from './editAttribute'

export var AttributeInterfaces = {
    "newattribute": {
        level: -1,
        name: "New Attribute", 
        component: (<NewAttribute />),
    },  
    "editattribute": {
        level: -1,
        name: "Edit Attribute", 
        component: (<EditAttribute />),
    },  
    "attributes": {
        level: 7,
        name: "Manage Attributes", 
        component: (<Attributes />),
    },
    "attribute": {
        level: 8,
        name: "Attribute", 
        component: (<Attribute />),
    },
}
