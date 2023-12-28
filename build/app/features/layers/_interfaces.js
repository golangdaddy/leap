import * as React from 'react'

import { Layers } from './layers'
import { Layer } from './layer'
import { NewLayer } from './newLayer'
import { EditLayer } from './editLayer'

export var LayerInterfaces = {
    "newlayer": {
        level: -1,
        name: "New Layer", 
        component: (<NewLayer />),
    },  
    "editlayer": {
        level: -1,
        name: "Edit Layer", 
        component: (<EditLayer />),
    },  
    "layers": {
        level: 7,
        name: "Manage Layers", 
        component: (<Layers />),
    },
    "layer": {
        level: 8,
        name: "Layer", 
        component: (<Layer />),
    },
}
