// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
  // Hacky: need to import this package to let `aedeploy` know and copy to Docker
  _ "github.com/chai2010/webp"

  "google.golang.org/appengine"
  "webp_demo"
)

func main() {
  webp_demo.InitHttpHandlers()
  // https://godoc.org/google.golang.org/appengine#Main
  appengine.Main()
}
