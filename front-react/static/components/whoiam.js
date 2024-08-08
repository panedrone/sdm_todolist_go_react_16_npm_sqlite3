import * as React from "react";

import * as shared from './shared.js'
import * as api from "./api"
import {RawHtml} from "./raw_html"


export function fetchWhoIAm() {
    api.getText('api/whoiam', (text) => {
        if (!text) {
            return
        }
        if (text.includes('sqlx')) {
            text += ', <a target="_blank" href="swagger/index.html">swagger</a>'
        }
        text += ", npm, react " + React.version
        shared.render(<RawHtml rawHtml={text}/>, 'whoiam')
    })
}