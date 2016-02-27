## Demo GAE custom runtime

### Development

- Install packages:

```
go get github.com/chai2010/webp
go get google.golang.org/appengine
```

- Add current workspace to `GOPATH`:

```
export GOPATH=$GOPATH:$(pwd)
```

- Run server in localhost:

```
go run main.go
```

  Try it!:

```
open http://localhost:8080/webp
```

### Deployment

* Set up Google Cloud SDK (if you haven't)

```
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init
```

* Deploy to GAE:

```
go get -u google.golang.org/appengine/cmd/aedeploy
aedeploy gcloud preview app deploy app.yaml --promote --project "gae-custom-vm-try-2"
```

See: https://cloud.google.com/sdk/gcloud/reference/preview/app/deploy

* Try it!:

```
open https://gae-custom-vm-try-2.appspot.com/webp
```
