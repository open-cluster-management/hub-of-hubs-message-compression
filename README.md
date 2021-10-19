[comment]: # ( Copyright Contributors to the Open Cluster Management project )

# hub-of-hubs-message-compression

[![Go Report Card](https://goreportcard.com/badge/github.com/open-cluster-management/hub-of-hubs-message-compression)](https://goreportcard.com/report/github.com/open-cluster-management/hub-of-hubs-message-compression)
[![License](https://img.shields.io/github/license/open-cluster-management/hub-of-hubs-message-compression)](/LICENSE)

The message-compression logic component of [Hub-of-Hubs](https://github.com/open-cluster-management/hub-of-hubs).

This repo provides compression logic to be used by different transport components for compression/decompression of synced messages.

## How it works

Code that wants to use a supported Compressor must call NewCompressor function with a string identifier of any supported compression-logic implementation.

#### Supported Compression-Logic (found as constants in compressors_map.go):

- GZip - "gzip" : implements compression based on [compress/gzip](https://pkg.go.dev/compress/gzip) package.