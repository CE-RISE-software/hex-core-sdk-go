# CE-RISE Go Software Development Kit for Hexagonal Core Service

[![Go Reference](https://pkg.go.dev/badge/github.com/CE-RISE-software/hex-core-sdk-go.svg)](https://pkg.go.dev/github.com/CE-RISE-software/hex-core-sdk-go)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.18953368.svg)](https://doi.org/10.5281/zenodo.18953368)

A Go SDK for the CE-RISE Hex Core Service: https://codeberg.org/CE-RISE-software/hex-core-service.

---

## Package

- Module: `github.com/CE-RISE-software/hex-core-sdk-go`
- Go package name: `hexcoresdk`
- Reference docs: https://pkg.go.dev/github.com/CE-RISE-software/hex-core-sdk-go

## Install

```bash
go get github.com/CE-RISE-software/hex-core-sdk-go@latest
```

## Use the SDK

```go
package main

import (
	"context"
	"fmt"
	"log"

	hexcoresdk "github.com/CE-RISE-software/hex-core-sdk-go"
)

func main() {
	cfg := hexcoresdk.NewConfiguration()
	cfg.Servers = hexcoresdk.ServerConfigurations{
		{
			URL: "https://your-hex-core.example.org",
		},
	}

	client := hexcoresdk.NewAPIClient(cfg)

	health, httpRes, err := client.AdminAPI.Health(context.Background()).Execute()
	if err != nil {
		log.Fatalf("health request failed: %v", err)
	}
	fmt.Printf("health status code: %d\n", httpRes.StatusCode)
	fmt.Printf("health response: %+v\n", health)

	models, _, err := client.DiscoveryAPI.ListModels(context.Background()).Execute()
	if err != nil {
		log.Fatalf("list models failed: %v", err)
	}
	fmt.Printf("models response: %+v\n", models)
}
```

## API Documentation

- pkg.go.dev reference: https://pkg.go.dev/github.com/CE-RISE-software/hex-core-sdk-go
- Generated documentation website: https://ce-rise-software.codeberg.page/hex-core-sdk-go/


## License

Licensed under the [European Union Public Licence v1.2 (EUPL-1.2)](LICENSE).

## Contributing

This repository is maintained on [Codeberg](https://codeberg.org/CE-RISE-software/hex-core-sdk-go) — the canonical source of truth. The GitHub repository is a read mirror used for release archival and Zenodo integration. Issues and pull requests should be opened on Codeberg.

---

<a href="https://europa.eu" target="_blank" rel="noopener noreferrer">
  <img src="https://ce-rise.eu/wp-content/uploads/2023/01/EN-Funded-by-the-EU-PANTONE-e1663585234561-1-1.png" alt="EU emblem" width="200"/>
</a>

Funded by the European Union under Grant Agreement No. 101092281 — CE-RISE.  
Views and opinions expressed are those of the author(s) only and do not necessarily reflect those of the European Union or the granting authority (HADEA).
Neither the European Union nor the granting authority can be held responsible for them.

© 2026 CE-RISE consortium.  
Licensed under the [European Union Public Licence v1.2 (EUPL-1.2)](LICENSE).  
Attribution: CE-RISE project (Grant Agreement No. 101092281) and the individual authors/partners as indicated.

<a href="https://www.nilu.com" target="_blank" rel="noopener noreferrer">
  <img src="https://nilu.no/wp-content/uploads/2023/12/nilu-logo-seagreen-rgb-300px.png" alt="NILU logo" height="20"/>
</a>

Developed by NILU (Riccardo Boero — ribo@nilu.no) within the CE-RISE project.
