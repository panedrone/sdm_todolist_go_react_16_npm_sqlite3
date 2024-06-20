import * as React from "react";

import * as shared from './shared.js'
import * as api from "./api"
import {RawHtml} from "./raw_html"


export function fetchWhoIAm() {
    api.getText('api/whoiam', (res) => {
        if (!res) {
            res = '== unknown =='
        }
        res += ", react " + React.version
        shared.render(<RawHtml rawHtml={res}/>, 'whoiam')
    })
}