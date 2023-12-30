import * as React from 'react'

import { Elements } from './elements'
import { Element } from './element'
import { NewElement } from './newElement'
import { EditElement } from './editElement'

export var ElementInterfaces = {
    "newelement": {
        level: 2,
        name: "New Element", 
        component: (<NewElement />),
    },  
    "editelement": {
        level: 3,
        name: "Edit Element", 
        component: (<EditElement />),
    },  
    "elements": {
        level: 1,
        name: "Manage Elements", 
        component: (<Elements />),
        subsublinks: ["newelement"],
    },
    "element": {
        level: 2,
        name: "Element", 
        subsublinks: [""],
        component: (<Element />),
    },
}
