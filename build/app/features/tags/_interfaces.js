import * as React from 'react'

import { Tags } from './tags'
import { Tag } from './tag'
import { NewTag } from './newTag'
import { EditTag } from './editTag'
import { UploadTag } from './uploadTag'

export var TagInterfaces = {
    "newtag": {
        level: 2,
        name: "New Tag", 
        component: (<NewTag />),
    },  
    "edittag": {
        level: 3,
        name: "Edit Tag", 
        component: (<EditTag />),
    },  
    "uploadtag": {
        level: 3,
        name: "Upload File", 
        component: (<UploadTag />),
    },  
    "tags": {
        level: 1,
        name: "Manage Tags", 
        component: (<Tags />),
        subsublinks: ["newtag"],
    },
    "tag": {
        level: 2,
        name: "Tag",
        sublinks: ["edittag", "uploadtag"],
        subsublinks: [""],
        component: (<Tag />),
    },
}
