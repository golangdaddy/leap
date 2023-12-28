import * as React from 'react'

import { Fonts } from './fonts'
import { Font } from './font'
import { NewFont } from './newFont'
import { EditFont } from './editFont'

export var FontInterfaces = {
    "newfont": {
        level: -1,
        name: "New Font", 
        component: (<NewFont />),
    },  
    "editfont": {
        level: -1,
        name: "Edit Font", 
        component: (<EditFont />),
    },  
    "fonts": {
        level: 7,
        name: "Manage Fonts", 
        component: (<Fonts />),
    },
    "font": {
        level: 8,
        name: "Font", 
        component: (<Font />),
    },
}
