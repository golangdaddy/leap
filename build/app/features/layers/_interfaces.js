import * as React from 'react'

import { Layers } from './layers'
import { Layer } from './layer'
import { NewLayer } from './newLayer'
import { EditLayer } from './editLayer'
import { UploadLayer } from './uploadLayer'

export var LayerInterfaces = {
    "newlayer": {
        level: 2,
        name: "New Layer", 
        component: (<NewLayer />),
    },  
    "editlayer": {
        level: 3,
        name: "Edit Layer", 
        component: (<EditLayer />),
    },  
    "uploadlayer": {
        level: 3,
        name: "Upload File", 
        component: (<UploadLayer />),
    },  
    "layers": {
        level: 1,
        name: "Manage Layers", 
        component: (<Layers />),
        subsublinks: ["newlayer"],
    },
    "layer": {
        level: 2,
        name: "Layer",
        sublinks: ["editlayer", "uploadlayer"],
        subsublinks: ["elements","tags",""],
        component: (<Layer />),
    },
}
