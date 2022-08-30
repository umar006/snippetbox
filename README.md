# Snippetbox

A web application called Snippetbox, which lets people paste and share snippets of text — a bit like Pastebin or GitHub’s Gists.

## Routes

| Method | Pattern | Action |
| ------ | ------- | ------ |
| GET | / | Display the home page |
| GET | /snippet/view/:id	| Display a specific snippet |
| GET | /snippet/create	| Display a HTML form for creating a new snippet |
| POST | /snippet/create | Create a new snippet |
| GET | /user/signup | Display a HTML form for signing up a new user |
| POST | /user/signup | Create a new user |
| GET | /user/login | Display a HTML form for logging in a user |
| POST | /user/login | Authenticate and login the user |
| GET | /user/logout | Logout the user |
| GET | /static/*filepath | Serve a specific static file |
