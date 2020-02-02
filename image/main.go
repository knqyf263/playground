package main

import (
	"compress/gzip"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"

	"github.com/containers/image/v5/image"
	"github.com/containers/image/v5/pkg/blobinfocache"
	"github.com/containers/image/v5/transports/alltransports"
	t "github.com/containers/image/v5/types"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	if err := calcHash("alpine:3.7", true); err != nil {
		return err
	}
	if err := calcHash("alpine:3.7", false); err != nil {
		return err
	}
	return nil
}

func calcHash(imageName string, fromRegistry bool) error {
	if fromRegistry {
		fmt.Printf("%s from Docker Hub\n", imageName)
		imageName = "docker://" + imageName
	} else {
		fmt.Printf("%s from local Docker Engine\n", imageName)
		imageName = "docker-daemon:" + imageName
	}
	ref, err := alltransports.ParseImageName(imageName)
	if err != nil {
		return err
	}
	sys := &t.SystemContext{
		OSChoice: "linux",
	}
	ctx := context.Background()
	rawSource, err := ref.NewImageSource(ctx, sys)
	if err != nil {
		return err
	}
	src, err := image.FromSource(ctx, sys, rawSource)
	if err != nil {
		return err
	}

	fmt.Printf("Image ID: %s\n", src.ConfigInfo().Digest)

	c2 := blobinfocache.DefaultCache(sys)
	for _, layer := range src.LayerInfos() {
		fmt.Printf("Layer ID (Before decompression): %s\n", layer.Digest)
		r, _, err := rawSource.GetBlob(ctx, t.BlobInfo{Digest: layer.Digest, Size: -1}, c2)
		if err != nil {
			return err
		}
		if fromRegistry {
			r, err = gzip.NewReader(r)
			if err != nil {
				return err
			}
		}
		s := sha256.New()
		if _, err = io.Copy(s, r); err != nil {
			return err
		}
		fmt.Printf("Layer ID (After decompression) : sha256:%x\n\n", s.Sum(nil))
	}
	return nil
}
