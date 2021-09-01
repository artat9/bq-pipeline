package main

import (
	"aurora-backend/lib/functions/lib/common/log"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func main() {
	f, _ := openGltfFile()
	toSingleFile(f, "./")
	saveAsBinary(f)

}

func saveAsBinary(doc *gltf.Document) error {
	toSingleFile(doc, "./")
	return gltf.SaveBinary(doc, "matic.glb")
}

func openGltfFile() (*gltf.Document, error) {
	return gltf.Open("matic.gltf")
}

func toSingleFile(doc *gltf.Document, srcDir string) error {
	path, _ := filepath.Abs("./")
	fmt.Print("current path")
	fmt.Print(path)
	for _, b := range doc.Buffers {
		b.URI = ""
	}
	for _, m := range doc.Images {
		if m.BufferView == nil && m.URI != "" {
			path, _ := filepath.Rel(srcDir, m.URI)
			log.Info("path")
			log.Info(path)
			f, err := os.Open(path)
			if err != nil {
				log.Error("cant open file", err)
				continue
			}
			defer f.Close()
			buf, err := ioutil.ReadAll(f)
			if err != nil {
				log.Error("can not read file", err)
				continue
			}
			if m.MimeType == "" {
				if strings.HasSuffix(strings.ToLower(m.URI), ".png") {
					m.MimeType = "image/png"
				} else {
					m.MimeType = "image/jpeg"
				}
			}
			m.BufferView = gltf.Index(modeler.WriteBufferView(doc, gltf.TargetNone, buf))
			m.URI = ""
		}
	}
	return nil
}
