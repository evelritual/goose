package tiled

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// TileMap represents a Tiled Map
type TileMap struct {
	XMLName xml.Name `xml:"map"`

	Height       int     `xml:"height,attr"`
	Infinite     bool    `xml:"infinite,attr"`
	NextLayerID  int     `xml:"nextlayerid,attr"`
	NextObjectID int     `xml:"nextobjectid,attr"`
	Orientation  string  `xml:"orientation,attr"`
	RenderOrder  string  `xml:"renderorder,attr"`
	TiledVersion string  `xml:"tiledversion,attr"`
	TileHeight   int     `xml:"tileheight,attr"`
	TileWidth    int     `xml:"tilewidth,attr"`
	Version      float32 `xml:"version,attr"`
	Width        int     `xml:"width,attr"`

	EditorSettings *EditorSetting `xml:"editorsettings"`
	TileSets       []*MapTileSet  `xml:"tileset"`
	Layers         []*Layer       `xml:"layer"`
}

// EditorSetting represents settings in a Tiled map
type EditorSetting struct {
	Export *Export `xml:"export"`
}

// Export represents export settings in a Tiled map
type Export struct {
	Format string `xml:"format,attr"`
	Target string `xml:"target,attr"`
}

// MapTileSet represents the used tile sets in a Tiled map
type MapTileSet struct {
	FirstGID int    `xml:"firstgid,attr"`
	Source   string `xml:"source,attr"`
}

// Layer represents a single Tiled map layer
type Layer struct {
	Data *Data `xml:"data"`

	Height int    `xml:"height,attr"`
	ID     int    `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Width  int    `xml:"width,attr"`
}

// Data represents a layers encoded data
type Data struct {
	Encoding string `xml:"encoding,attr"`
	Value    string `xml:",chardata"`
}

// loadTileMap loads a Tiled tmx file.
func loadTileMap(tmxPath string) (*TileMap, error) {
	f, err := os.Open(tmxPath)
	if err != nil {
		return nil, fmt.Errorf("error opening tmx: %v", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading tmx: %v", err)
	}

	var m *TileMap
	err = xml.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling xml: %v", err)
	}

	return m, nil
}
