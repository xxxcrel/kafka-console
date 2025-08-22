// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package filesystem

// File is a file with its metadata and payload.
type File struct {
	Path     string
	Filename string

	// TrimmedFilename is the filename without the recognized file extension
	TrimmedFilename string

	Payload []byte
}
