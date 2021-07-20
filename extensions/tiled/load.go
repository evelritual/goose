package tiled

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/evelritual/goose"
	"github.com/evelritual/goose/graphics"
)

// Map represents a Tiled map.
type Map struct {
	TileMap  *TileMap
	TileSets []*TileSet

	TextureAtlases []graphics.TextureAtlas
	LayerData      [][][]int
}

// LoadMap loads a Tiled tmx and converts it to an easy-to-use Goose structure.
func LoadMap(tmxPath string) (*Map, error) {
	tmx, err := loadTileMap(tmxPath)
	if err != nil {
		return nil, fmt.Errorf("error loading tile map: %v", err)
	}

	tileSets := []*TileSet{}
	tileSetTextures := []graphics.TextureAtlas{}
	for _, ts := range tmx.TileSets {
		tsx, err := loadTileSet(ts.Source)
		if err != nil {
			return nil, fmt.Errorf("error loading tile set: %v", err)
		}
		tileSets = append(tileSets, tsx)

		ta, err := goose.NewTextureAtlas(tsx.Image.Source, int32(tsx.TileWidth), int32(tsx.TileHeight))
		if err != nil {
			return nil, fmt.Errorf("error loading texture atlas: %v", err)
		}
		tileSetTextures = append(tileSetTextures, ta)
	}

	layerData := [][][]int{}
	for _, l := range tmx.Layers {
		if l.Data.Encoding != "csv" {
			return nil, fmt.Errorf("encoding not supported: %v", err)
		}

		data := strings.Split(strings.TrimSpace(l.Data.Value), "\n")
		layer := [][]int{}
		for _, d := range data {
			dataConvert := []int{}
			dataParts := strings.Split(strings.TrimSuffix(d, ","), ",")
			for _, dc := range dataParts {
				i, err := strconv.Atoi(dc)
				if err != nil {
					return nil, fmt.Errorf("error converting string to int: %v", err)
				}
				dataConvert = append(dataConvert, i-1)
			}
			layer = append(layer, dataConvert)
		}

		layerData = append(layerData, layer)
	}

	return &Map{
		TileMap:        tmx,
		TileSets:       tileSets,
		TextureAtlases: tileSetTextures,
		LayerData:      layerData,
	}, nil
}

// Close closes all loaded texture atlases for the loaded map.
func (m *Map) Close() error {
	for _, t := range m.TextureAtlases {
		t.Close()
	}
	return nil
}
