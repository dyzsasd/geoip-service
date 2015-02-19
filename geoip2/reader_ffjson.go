// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: reader.go
// DO NOT EDIT!

package geoip2

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"strconv"
)

func (mj *City) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(512)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *City) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"City":`)

	{
		err = mj.City.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Continent":`)

	{
		err = mj.Continent.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Country":`)

	{
		err = mj.Country.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Location":`)

	{
		err = mj.Location.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Postal":`)

	{
		err = mj.Postal.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"RegisteredCountry":`)

	{
		err = mj.RegisteredCountry.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"RepresentedCountry":`)

	{
		err = mj.RepresentedCountry.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Subdivisions":`)
	if mj.Subdivisions != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Subdivisions {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{
				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"Traits":`)

	{
		err = mj.Traits.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteByte('}')
	return nil
}

func (mj *ConnectionType) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(64)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ConnectionType) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"ConnectionType":`)
	fflib.WriteJsonString(buf, string(mj.ConnectionType))
	buf.WriteByte('}')
	return nil
}

func (mj *Continent) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Continent) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"Code":`)
	fflib.WriteJsonString(buf, string(mj.Code))
	buf.WriteString(`,"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

func (mj *Country) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(512)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Country) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Continent":`)

	{
		err = mj.Continent.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Country":`)

	{
		err = mj.Country.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"RegisteredCountry":`)

	{
		err = mj.RegisteredCountry.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"RepresentedCountry":`)

	{
		err = mj.RepresentedCountry.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`,"Subdivisions":`)
	if mj.Subdivisions != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Subdivisions {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{
				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"Traits":`)

	{
		err = mj.Traits.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}
	}

	buf.WriteByte('}')
	return nil
}

func (mj *Domain) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(64)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Domain) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Domain":`)
	fflib.WriteJsonString(buf, string(mj.Domain))
	buf.WriteByte('}')
	return nil
}

func (mj *ISP) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ISP) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"AutonomousSystemNumber":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.AutonomousSystemNumber), 10, false)
	buf.WriteString(`,"AutonomousSystemOrganization":`)
	fflib.WriteJsonString(buf, string(mj.AutonomousSystemOrganization))
	buf.WriteString(`,"ISP":`)
	fflib.WriteJsonString(buf, string(mj.ISP))
	buf.WriteString(`,"Organization":`)
	fflib.WriteJsonString(buf, string(mj.Organization))
	buf.WriteByte('}')
	return nil
}

func (mj *Location) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(128)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Location) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"Latitude":`)
	buf.Write(strconv.AppendFloat([]byte{}, mj.Latitude, 'g', -1, 64))
	buf.WriteString(`,"Longitude":`)
	buf.Write(strconv.AppendFloat([]byte{}, mj.Longitude, 'g', -1, 64))
	buf.WriteString(`,"MetroCode":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.MetroCode), 10, false)
	buf.WriteString(`,"TimeZone":`)
	fflib.WriteJsonString(buf, string(mj.TimeZone))
	buf.WriteByte('}')
	return nil
}

func (mj *Postal) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(64)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Postal) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Code":`)
	fflib.WriteJsonString(buf, string(mj.Code))
	buf.WriteByte('}')
	return nil
}

func (mj *Reader) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(8)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Reader) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{}`)
	return nil
}

func (mj *RegisteredCountry) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *RegisteredCountry) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	buf.WriteString(`,"IsoCode":`)
	fflib.WriteJsonString(buf, string(mj.IsoCode))
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

func (mj *RepresentedCountry) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *RepresentedCountry) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	buf.WriteString(`,"IsoCode":`)
	fflib.WriteJsonString(buf, string(mj.IsoCode))
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteString(`,"Type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteByte('}')
	return nil
}

func (mj *Subdivision) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Subdivision) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	buf.WriteString(`,"IsoCode":`)
	fflib.WriteJsonString(buf, string(mj.IsoCode))
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

func (mj *TheCity) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *TheCity) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

func (mj *TheCountry) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *TheCountry) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"GeoNameID":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.GeoNameID), 10, false)
	buf.WriteString(`,"IsoCode":`)
	fflib.WriteJsonString(buf, string(mj.IsoCode))
	if mj.Names == nil {
		buf.WriteString(`,"Names":null`)
	} else {
		buf.WriteString(`,"Names":{ `)
		for key, value := range mj.Names {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

func (mj *Traits) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(16)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Traits) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	if mj.IsAnonymousProxy {
		buf.WriteString(`{"IsAnonymousProxy":true`)
	} else {
		buf.WriteString(`{"IsAnonymousProxy":false`)
	}
	if mj.IsSatelliteProvider {
		buf.WriteString(`,"IsSatelliteProvider":true`)
	} else {
		buf.WriteString(`,"IsSatelliteProvider":false`)
	}
	buf.WriteByte('}')
	return nil
}
