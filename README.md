# iamfeelingcody
Golang, Templ, and Htmx toy project.

# Test it!
```sh
export IAMFEELINGCODY_GOPATH_PARENT="${HOME}/iamfeelingcody-go"

chmod u+x ./main.sh
./main.sh install
./main.sh run
```

`./main.sh build` creates the executable `bin/iamfeelingcody`.

## VSCode

Run following to have:
* go and dependencies installed.
* environment variables set (e.g. PATH, GOPATH)

```sh
(cd ~/prj/iamfeelingcody/; ./main.sh install; source ./env.sh; code .)
```
