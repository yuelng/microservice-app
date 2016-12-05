package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

// postgresql数据库存放的就是 16进制数据,所以在golang中可以相互转换.
// 具体实现 需要看 存储的细节.这里只是一个逆向过程
func (p *GeoPoint) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

func (t GeoPoint) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(struct {
		Lng string
		Lat string
	}{
		Lng: strconv.FormatFloat(t.Lng, 'f', 6, 64),
		Lat: strconv.FormatFloat(t.Lat, 'f', 6, 64),
	})
	return v, err
}

type Location struct {
	Base
	Ponit GeoPoint `sql:"type:geometry(Geometry,4326)"`
}
