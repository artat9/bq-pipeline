package generator

import (
	"aurora-backend/lib/functions/lib/common/log"
	"aurora-backend/lib/functions/lib/infrastructure/lambda"
	"aurora-backend/lib/functions/lib/receipt"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
	"golang.org/x/sync/errgroup"
)

const (
	pngFilePrefix = "/tmp/"
	dst           = "/tmp/result.glb"
)

type (
	Generator struct {
		client lambda.Invoker
	}
)

func NewGenerator() Generator {
	return Generator{
		client: lambda.NewInvoker(),
	}
}

func saveTmpImg(img []byte) error {
	decoded, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		log.Error("decode as png failed", err)
		return err
	}
	return saveAsPng(decoded, "project.png")
}

// Generate Generate Receipt model
func (g Generator) Generate(ctx context.Context, param receipt.GenerateParameter) ([]byte, error) {
	eg := errgroup.Group{}
	eg.Go(func() error {
		return g.dateToPng(ctx, param.Info.Issued)
	})
	eg.Go(func() error {
		return g.titleToPng(ctx, param.Info.Post.ProjectName)
	})
	eg.Go(func() error {
		return g.descriptionToPng(ctx, param.Info.Post.Description)
	})
	eg.Go(func() error {
		return g.amtToPng(ctx, param.Amount)
	})
	eg.Go(func() error {
		return g.serialToPng(ctx, param.Context.SerialNo)
	})
	if err := eg.Wait(); err != nil {
		log.Error("unexpected error occured", err)
		return nil, err
	}
	doc, err := openGltfFile(param.Context.SerialNo)
	if err != nil {
		log.Error("open gltf file failed", err)
		return nil, err
	}

	if err = saveAsBinary(doc); err != nil {
		log.Error("save binary failed", err)
	}
	return loadResult()
}

func loadResult() ([]byte, error) {
	return ioutil.ReadFile(dst)
}

func saveAsBinary(doc *gltf.Document) error {
	toSingleFile(doc, "./")
	if err := gltf.SaveBinary(doc, dst); err != nil {
		log.Error("save binary failed", err)
		return err
	}
	return nil
}

func openGltfFile(serialNum int) (*gltf.Document, error) {
	prefix := "./assets/gltf/"
	if !receipt.UltraRare(serialNum) {
		return gltf.Open(prefix + "normal/normalcard.gltf")
	}
	return gltf.Open(prefix + "ultrarare/rarecard.gltf")
}

func saveAsPng(img image.Image, filename string) error {
	f, err := os.Create(pngFilePrefix + filename + ".png")
	if err != nil {
		log.Error("create png failed", err)
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		log.Error("write image failed", err)
	}
	return nil
}

func (g Generator) dateToPng(ctx context.Context, date time.Time) error {
	return g.toPng(func() (image.Image, error) {
		return g.client.DateImage(ctx, datefmt(date))
	}, "date")
}

func (g Generator) amtToPng(ctx context.Context, amt string) error {
	return g.toPng(func() (image.Image, error) {
		return g.client.AmountImage(ctx, amt, "MATIC")
	}, "amount")
}

func (g Generator) descriptionToPng(ctx context.Context, desc string) error {
	return g.toPng(func() (image.Image, error) {
		return g.client.DescriptionImage(ctx, desc)
	}, "description")
}

func (g Generator) titleToPng(ctx context.Context, title string) error {
	return g.toPng(func() (image.Image, error) {
		return g.client.TitleImage(ctx, title)
	}, "title")
}

func toSerialString(serial int) string {
	ser := strconv.Itoa(serial)
	for utf8.RuneCountInString(ser) < 6 {
		ser = "0" + ser
	}
	return ser
}

func (g Generator) serialToPng(ctx context.Context, serial int) error {
	return g.toPng(func() (image.Image, error) {
		return g.client.SerialNoImage(ctx, toSerialString(serial))
	}, "serial")
}

func (g Generator) toPng(toimg func() (image.Image, error), filename string) error {
	img, err := toimg()
	if err != nil {
		return err
	}
	return saveAsPng(img, filename)
}

func datefmt(d time.Time) string {
	return d.Format("2006/01/02")
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
