[comment]: # ( Copyright Contributors to the Open Cluster Management project )

# hub-of-hubs-message-compression

[![Go Report Card](https://goreportcard.com/badge/github.com/open-cluster-management/hub-of-hubs-message-compression)](https://goreportcard.com/report/github.com/open-cluster-management/hub-of-hubs-message-compression)
[![Go Reference](https://pkg.go.dev/badge/github.com/open-cluster-management/hub-of-hubs-message-compression.svg)](https://pkg.go.dev/github.com/open-cluster-management/hub-of-hubs-message-compression)
[![License](https://img.shields.io/github/license/open-cluster-management/hub-of-hubs-message-compression)](/LICENSE)

The message compression component of [Hub-of-Hubs](https://github.com/open-cluster-management/hub-of-hubs).

This repo provides compression logic to be used by different transport components for compression/decompression of synced messages.

Go to the [Contributing guide](CONTRIBUTING.md) to learn how to get involved.

## Getting Started

In order to use message compression one must call NewCompressor function with the string identifier of any of the supported compression implementations.

#### Supported implementations (found in [compressor types](https://github.com/open-cluster-management/hub-of-hubs-message-compression/blob/main/compressor.go#L18)):

- GZip - "gzip" : implements compression based on [compress/gzip](https://pkg.go.dev/compress/gzip) package.
- NoOp - "no-op" : implements no-op compression based on the No-Op pattern.