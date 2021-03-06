// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package writers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gardener/docforge/pkg/api"
)

// FSWriter is implementation of Writer interface for writing blobs to the file system
type FSWriter struct {
	Root string
	Ext  string
}

func (f *FSWriter) Write(name, path string, docBlob []byte, node *api.Node) error {
	p := filepath.Join(f.Root, path)
	if len(docBlob) <= 0 {
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}

	if len(f.Ext) > 0 {
		name = fmt.Sprintf("%s.%s", name, f.Ext)
	} else {
		if node != nil && !strings.HasSuffix(name, ".md") {
			name = fmt.Sprintf("%s.md", name)
		}
	}

	filePath := filepath.Join(p, name)

	if err := ioutil.WriteFile(filePath, docBlob, 0644); err != nil {
		return fmt.Errorf("error writing %s: %v", filepath.Join(f.Root, path, name), err)
	}

	return nil
}
