import * as React from 'react'

import { Elements } from './elements'
import { Element } from './element'
import { NewElement } from './newElement'
import { EditElement } from './editElement'

export var ElementInterfaces = {
    "newelement": {
        level: -1,
        name: "New Element", 
        component: (<NewElement />),
    },  
    "editelement": {
        level: -1,
        name: "Edit Element", 
        component: (<EditElement />),
    },  
    "elements": {
        level: 7,
        name: "Manage Elements", 
        component: (<Elements />),
    },
    "element": {
        level: 8,
        name: "Element", 
        component: (<Element />),
    },
}
