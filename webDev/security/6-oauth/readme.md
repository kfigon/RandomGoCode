# oauth 2

uses all features of jwt to provide authenticate by other systems e.g. - login with account google. Use some website to authorize to other website.

2 ways:
* client credentials - provide in API, requires backend
* implicit - show google form, do everything by them

flow:
* provide login data
* google validates and redirects to original website + sends code/token/whatever
* my website exchanges code and secret for access token to goole
* my website uses token to get who the user is on google
* my website can now map googleId with our webId

## oauth2 vs 1
oauth1 is more secure xd but more work required

although oauth2 is good enough and easier
