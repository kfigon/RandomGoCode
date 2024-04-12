# basics

* hx-METHOD - method and url
* hx-swap
    * outer - replace entire target element
    * inner - put inside target
* hx-target - target element (css selector) to put response
    * we can also use some commands, like `closest li`, to get closest relative selector
* hx-trigger - event types that trigger requests
* hx-select - extract part of response (css selector), not use the full one. Useful for redirects
* `hx-disabled-elt` - disable element (e.g. button) when request is issued
* add more data:
```
hx-vals='{"id": "{{.Id}}" }'
```
* `hx-params` to include data from different form fields
* `hx-headers` additional headers e.g. for csrf
* `hx-sync` - for syncing multiple requests (cancel other request when other element sent something)
* redirection: `r.Header.Add("HX-Redirect", "/path")`
* `hx-on` hooks for standard events or htmx specific
* `hx-boost` on body - all `a href` push data to url (so back button works), replace the data to body
    * works also on forms - translates regular form attributes `method` and `action` to `hx-method`and update entire page - useful for redirects
* most htmx elements are inherited to children, so `hx-swap` could be set on a list, instead of every list element
* by default only 200 status is rendered. 4xx and 5xx are ignored
* out of bands - we can send updates for 2 or more elements within single response
    * on the second response, just return `hx-swap-oob` (similar to hx-swap) and id to target. Or use `hx-swap-oob="true:#modal"` to get the same. Need to wrap in div, as the content will be swapped
```
<h2 id="modal" hx-swap-oob="true">
or
<h2 hx-swap-oob="true:#modal">
```