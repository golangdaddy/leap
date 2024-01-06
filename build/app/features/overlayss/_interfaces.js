import * as React from 'react'

import { Overlayss } from './overlayss'
import { Overlays } from './overlays'
import { NewOverlays } from './newOverlays'
import { EditOverlays } from './editOverlays'
import { UploadOverlays } from './uploadOverlays'

export var OverlaysInterfaces = {
    "newoverlays": {
        level: 2,
        name: "New Overlays", 
        component: (<NewOverlays />),
    },  
    "editoverlays": {
        level: 3,
        name: "Edit Overlays", 
        component: (<EditOverlays />),
    },  
    "uploadoverlays": {
        level: 3,
        name: "Upload File", 
        component: (<UploadOverlays />),
    },  
    "overlayss": {
        level: 1,
        name: "Manage Overlayss", 
        component: (<Overlayss />),
        subsublinks: ["newoverlays"],
    },
    "overlays": {
        level: 2,
        name: "Overlays",
        sublinks: ["editoverlays", "uploadoverlays"],
        subsublinks: [""],
        component: (<Overlays />),
    },
}
