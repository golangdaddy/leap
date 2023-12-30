import * as React from 'react'

import { Fonts } from './fonts'
import { Font } from './font'
import { NewFont } from './newFont'
import { EditFont } from './editFont'

export var FontInterfaces = {
    "newfont": {
        level: 2,
        name: "New Font", 
        component: (<NewFont />),
    },  
    "editfont": {
        level: 3,
        name: "Edit Font", 
        component: (<EditFont />),
    },  
    "fonts": {
        level: 1,
        name: "Manage Fonts", 
        component: (<Fonts />),
        subsublinks: ["newfont"],
    },
    "font": {
        level: 2,
        name: "Font", 
        subsublinks: [""],
        component: (<Font />),
    },
}
