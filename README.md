# Vecty VisibilityChange demo

Example project to test "visibilitychange" event listener in Vecty.

See [https://developer.mozilla.org/en-US/docs/Web/API/Page_Visibility_API](https://developer.mozilla.org/en-US/docs/Web/API/Page_Visibility_API) for details.

### How to test

```
go get github.com/divan/vecty-visibility-demo
gopherjs build // or use js stored in git if you don't have gopherjs
open index.html
```
then open Developers Tools / Console and switch between tabs - it should print something to console.

To check that it's not a browser issue, uncomment raw JS event listener code in index.html.
