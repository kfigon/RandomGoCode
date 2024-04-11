# basics

* hx-METHOD - method and url
* hx-swap
    * outer - replace entire target element
    * inner - put inside target
* hx-target - target element (css selector) to put response
    * we can also use some commands, like `closest li`, to get closest relative selector
* hx-trigger - event types that trigger requests
* hx-select - extract part of response (css selector), not use the full one. Useful for redirects


* `hx-on` hooks for standard events or htmx specific
* most htmx elements are inherited to children, so `hx-swap` could be set on a list, instead of every list element
* by default only 200 status is rendered. 4xx and 5xx are ignored