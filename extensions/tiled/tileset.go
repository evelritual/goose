package tiled

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// TileSet represents a Tiled tile set.
type TileSet struct {
	XMLName xml.Name `xml:"tileset"`

	Version      float32 `xml:"version,attr"`
	TiledVersion string  `xml:"tiledversion,attr"`
	Name         string  `xml:"name,attr"`
	TileWidth    int     `xml:"tilewidth,attr"`
	TileHeight   int     `xml:"tileheight,attr"`
	TileCount    int     `xml:"tilecount,attr"`
	Columns      int     `xml:"columns,attr"`

	Image *Image `xml:"image"`
}

// Image represents an image used in a tile set.
type Image struct {
	Source string `xml:"source,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

// loadTileSet loads a Tiled tsx file.
func loadTileSet(tsxPath string) (*TileSet, error) {
	f, err := os.Open(tsxPath)
	if err != nil {
		return nil, fmt.Errorf("error opening tsx: %v", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading tsx: %v", err)
	}

	var ts *TileSet
	err = xml.Unmarshal(b, &ts)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling xml: %v", err)
	}

	return ts, nil
}
