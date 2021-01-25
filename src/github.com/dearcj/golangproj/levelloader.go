package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/dearcj/golangproj/bitmask"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type TileObject struct {
	Coord        IntVec2
	ShortCutType bitmask.Bitmask
}

type Map struct {
	offsetX    float32
	offsetY    float32
	cols       int
	rows       int
	tileWidth  float32
	tileHeight float32
	Objects    []*TileObject
}

func (m *Map) roundCoord(wp *Vec2) Vec2 {
	t := m.getTileFromWorldPos(wp)
	x := (float32(t.Coord[0]) + 0.5) * m.tileWidth
	y := (float32(t.Coord[1]) + 0.5) * m.tileHeight

	return Vec2{x + m.offsetX, y + m.offsetY}
}

func (m *Map) getWorldPosFromTile(tile *IntVec2) *Vec2 {
	return &Vec2{float32(tile[0])*m.tileWidth + m.offsetX, float32(tile[1])*m.tileHeight + m.offsetY}
}

func (m *Map) getTileFromWorldPos(wp *Vec2) *TileObject {
	x := int(math.Floor(float64((wp[0] - m.offsetX) / m.tileWidth)))
	y := int(math.Floor(float64((wp[1] - m.offsetY) / m.tileHeight)))

	return m.getObject(x, y)
}

func (m *Map) getObject(x int, y int) *TileObject {
	inx := y*m.cols + x
	if inx < 0 || inx > len(m.Objects) {
		return nil
	}
	return m.Objects[inx]
}

func (m *Map) getXY(inx int) (int, int) {
	x := inx % m.cols
	y := inx / m.cols
	return x, y
}

func CreateMap(r int, c int, tw int, th int) *Map {
	m := &Map{cols: c, rows: r, tileWidth: float32(tw), tileHeight: float32(th)}
	m.Objects = make([]*TileObject, r*c)
	return m
}

type Level struct {
	objects     []TiledObject
	af          *ActorF
	spawnPoints []Vec2
}

func (l *Level) addSpawnPoint(v *Vec2) {
	l.spawnPoints = append(l.spawnPoints, *v)
}

func (l *Level) hasChildrenComponent(obj *TiledObject, classUID string) *TiledObject {
	/*	for _, component := range obj.Components {
			if l.objects[component.Id].Type == classUID {
				return &l.objects[component.Id]
			}
		}
	*/
	return nil

}

func (l *Level) getCenterPos(obj *TiledObject) *Vec2 {
	var offsetVec Vec2
	if obj.Gid > 0 {
		offsetVec = Vec2{obj.Width / 2, -obj.Height / 2}
	} else {
		offsetVec = Vec2{obj.Width / 2, obj.Height / 2}
	}
	rot := math.Pi * (obj.Rotation / 180)

	offsetVec = offsetVec.rotate(rot)
	return &Vec2{obj.X + offsetVec[0], obj.Y + offsetVec[1]}
}

func (l *Level) createStatic() {
	for _, obj := range l.objects {
		var _type = obj.GroupType

		if obj.Properties.find("type") != "" {
			_type = obj.Properties.find("type")
		}
		if obj.Type != "" {
			_type = obj.Type
		}

		if _type != "" {
			_type = strings.ToLower(_type)

			if _type == "spawnpoint" {
				l.addSpawnPoint(&Vec2{obj.X, obj.Y})
			}
		}
	}
}

type TiledObject struct {
	GroupType  string
	Type       string          `xml:"type,attr"`
	Name       string          `xml:"name,attr"`
	Id         int             `xml:"id,attr"`
	X          float32         `xml:"x,attr"`
	Y          float32         `xml:"y,attr"`
	Gid        int             `xml:"gid,attr"`
	Width      float32         `xml:"width,attr"`
	Height     float32         `xml:"height,attr"`
	Rotation   float32         `xml:"rotation,attr"`
	Properties TiledProperties `xml:"properties,omitempty"`
}

type TiledProperty struct {
	Value string `xml:"value,attr"`
	Name  string `xml:"name,attr"`
}

type TiledProperties struct {
	Property []TiledProperty `xml:"property,omitempty"`
}

func (properties *TiledProperties) find(s string) string {
	for _, prop := range properties.Property {
		if strings.ToLower(prop.Name) == strings.ToLower(s) {
			return prop.Value
		}
	}

	return ""
}

type ObjectGroup struct {
	Properties TiledProperties `xml:"properties,omitempty"`
	OffsetX    float32         `xml:"offsetx,attr"`
	OffsetY    float32         `xml:"offsety,attr"`
	Name       string          `xml:"name,attr"`
	Objects    []TiledObject   `xml:"object,omitempty"`
}

//<map version="1.0" orientation="orthogonal" renderorder="right-down" width="40" height="40" tilewidth="128" tileheight="128" nextobjectid="139">
type TiledMap struct {
	Columns      int           `xml:"width,attr"`
	Rows         int           `xml:"height,attr"`
	TileWidth    int           `xml:"tilewidth,attr"`
	TileHeight   int           `xml:"tileheight,attr"`
	Objectgroups []ObjectGroup `xml:"objectgroup,omitempty"`
	Layers       []TileLayer   `xml:"layer"`
}

type TileData struct {
}

// <layer name="bg" width="40" height="40" offsetx="-64" offsety="-64">
//<data encoding="csv">
type TileLayer struct {
	OffsetX    float32         `xml:"offsetx,attr"`
	OffsetY    float32         `xml:"offsety,attr"`
	Data       string          `xml:"data"`
	Properties TiledProperties `xml:"properties,omitempty"`
}

type TiledLevel struct {
	Maps []TiledMap `xml:"map"`
}

func (l *Level) loadLevel(level string) {
	filename := dirPath + "\\front\\levels\\" + level
	file, err := os.Open(filename)
	if err != nil {
		server.logger.Error("Errpr opening file", zap.String("file", filename), zap.Error(err))
		return
	}

	defer file.Close()

	b, _ := ioutil.ReadAll(file)

	tiledLev := TiledMap{}

	xml.Unmarshal(b, &tiledLev)
	var objects []TiledObject

	var replacer = strings.NewReplacer("\n", "")

	Map := CreateMap(tiledLev.Rows, tiledLev.Columns, tiledLev.TileWidth, tiledLev.TileHeight)

	for i := 0; i < len(Map.Objects); i++ {
		x, y := Map.getXY(i)
		Map.Objects[i] = &TileObject{Coord: IntVec2{x, y}}
	}

	for _, tilelayer := range tiledLev.Layers {
		tilelayer.Data = replacer.Replace(tilelayer.Data)
		Map.offsetX = tilelayer.OffsetX
		Map.offsetY = tilelayer.OffsetY
		//get offsetX / Y from last layer IDC
	}

	for _, objgroup := range tiledLev.Objectgroups {
		for _, obj := range objgroup.Objects {
			grouptype := strings.ToLower(objgroup.Properties.find("type"))
			for _, prop := range objgroup.Properties.Property {
				obj.Properties.Property = append(obj.Properties.Property, prop)
			}
			obj.GroupType = grouptype
			obj.X += objgroup.OffsetX
			obj.Y += objgroup.OffsetY
			objects = append(objects, obj)
		}
	}

	l.objects = objects
	l.createStatic()
	//	l.factory.initPathFinding()
}

type JSONRoom struct {
	Id     int
	Prefab string
}

type WorldObject struct {
	Rooms []JSONRoom `json:"runs,omitempty"`
}

func loadWorld(fileName string) *WorldObject {
	file, err := os.Open(dirPath + fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	b, _ := ioutil.ReadAll(file)

	objects := WorldObject{}

	json.Unmarshal(b, &objects)
	return &objects
}
