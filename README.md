# iamfeelingcody
Golang toy project. This app is to give a psuedo-random Github user on request.

Makes use of:
* httprouter (TODO -> fasthttp/router)
* Sqlc (for SQLite)
* Htmx and Templ

Currently scripting is the backbone of development.
```
// TODO towards the end of project
defer("go gen")
defer("openapi3")
```

## Test it!
Workdir is always `src`
```sh
git clone https://github.com/cbugk/iamfeelingcody
cd iamfeelingcody/src
```
---
* Run:
    ```sh
    export IAMFEELINGCODY_GOPATH_PARENT="${HOME}/iamfeelingcody-go"
    
    chmod u+x ./main.sh
    ./main.sh install
    ./main.sh runbin
    ```
* `./main.sh cleanrunbin` runs with clean directory structure.
* `./main.sh build` creates the executable `../bin/iamfeelingcody`.
* `public` directory beside the binary (i.e. `../bin/public`) is served under `/public/` route.
* `pkg/embed/static` directory is embedded into executable and its content is served under `/static/` route.

## VSCode
To get:
* go and dependencies installed.
* environment variables set (e.g. PATH, GOPATH)

Run:
```sh
(cd ~/prj/iamfeelingcody/src && ./main.sh install && source ./env.sh && code ..)
```
