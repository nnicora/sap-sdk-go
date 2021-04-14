package request

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/saperr"
	"github.com/nnicora/sap-sdk-go/internal/times"
	"github.com/nnicora/sap-sdk-go/internal/types"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	fieldTagDest     = "dest"
	fieldTagDestName = "dest-name"

	fieldTagSrc     = "src"
	fieldTagSrcName = "src-name"
)

var noEscape [256]bool

var errValueNotSet = fmt.Errorf("value not set")

var byteSliceType = reflect.TypeOf([]byte{})

func init() {
	for i := 0; i < len(noEscape); i++ {
		// expects every character except these to be escaped
		noEscape[i] = (i >= 'A' && i <= 'Z') ||
			(i >= 'a' && i <= 'z') ||
			(i >= '0' && i <= '9') ||
			i == '-' ||
			i == '.' ||
			i == '_' ||
			i == '~'
	}
}

func (r *Request) readFromHttpResponseTo(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Kind()

		if fieldType == reflect.Struct {
			nestObj := field.Addr().Interface()
			r.readFromHttpResponseTo(nestObj)
		} else {
			structField := v.Type().Field(i)
			r.readFromHttpResponseToStruct(field, structField)
		}

		if r.Error != nil {
			return
		}
	}
}
func (r *Request) readFromHttpResponseToStruct(value reflect.Value, structField reflect.StructField) {
	if n := structField.Name; n[0:1] == strings.ToLower(n[0:1]) {
		return
	}

	if value.IsValid() {
		name := structField.Tag.Get(fieldTagSrcName)
		if name == "" {
			name = structField.Name
		}
		if kind := value.Kind(); kind == reflect.Ptr {
			value = value.Elem()
		} else if kind == reflect.Interface {
			if !value.Elem().IsValid() {
				return
			}
		}
		if !value.IsValid() {
			return
		}
		if structField.Tag.Get("ignore") != "" {
			return
		}

		var err error
		switch structField.Tag.Get(fieldTagSrc) {
		case "header":
			err = updateFromHeader(&r.HTTPResponse.Header, value, name)
		case "body":
			err = updateFromBody(r.ResponseBody, value, name)
		case "status":
			err = updateFromStatus(r.HTTPResponse, value, name)
		default:
			// ignore
		}
		r.Error = err
	}
}
func updateFromHeader(header *http.Header, v reflect.Value, name string) error {
	var err error = nil
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("field %s expected to be of type string; %v", name, errRecover)
		}
	}()

	name = strings.TrimSpace(name)

	headerValue := header.Get(name)
	switch v.Interface().(type) {
	case string:
		v.SetString(headerValue)
	case []byte:
		v.SetBytes([]byte(headerValue))
	default:
		return fmt.Errorf("unsupported value for field %v (%s)", name, v.Type())
	}
	return err
}
func updateFromBody(body []byte, v reflect.Value, name string) error {
	var err error = nil
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("field %s expected to be of type string; %v", name, errRecover)
		}
	}()

	if body != nil && len(body) > 0 {
		switch v.Interface().(type) {
		case string:
			v.SetString(string(body))
		case []byte:
			v.SetBytes(body)
		default:
			return fmt.Errorf("unsupported value for field %v (%s)", name, v.Type())
		}
	}
	return err
}
func updateFromStatus(resp *http.Response, v reflect.Value, name string) error {
	var err error = nil
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("field %s expected to be of type string; %v", name, errRecover)
		}
	}()

	if resp != nil {
		switch v.Interface().(type) {
		case string:
			v.SetString(resp.Status)
		case int, int8, int16, int32, int64:
			v.SetInt(int64(resp.StatusCode))
		case uint, uint8, uint16, uint32, uint64:
			v.SetUint(uint64(resp.StatusCode))
		default:
			return fmt.Errorf("unsupported value for field %v (%s)", name, v.Type())
		}
	}
	return err
}

func (r *Request) writeToHttpRequestFrom(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()

	query := r.HTTPRequest.URL.Query()

	// Setup the raw path to match the base path pattern. This is needed
	// so that when the path is mutated a custom escaped version can be
	// stored in RawPath that will be used by the Go client.
	r.HTTPRequest.URL.RawPath = r.HTTPRequest.URL.Path

	for i := 0; i < v.NumField(); i++ {
		m := v.Field(i)
		if n := v.Type().Field(i).Name; n[0:1] == strings.ToLower(n[0:1]) {
			continue
		}

		if m.IsValid() {
			field := v.Type().Field(i)
			name := field.Tag.Get(fieldTagDestName)
			if name == "" {
				name = field.Name
			}
			if kind := m.Kind(); kind == reflect.Ptr {
				m = m.Elem()
			} else if kind == reflect.Interface {
				if !m.Elem().IsValid() {
					continue
				}
			}
			if !m.IsValid() {
				continue
			}
			if field.Tag.Get("ignore") != "" {
				continue
			}

			var err error
			switch field.Tag.Get(fieldTagDest) {
			case "headers": // header maps
				err = buildHeaderMap(&r.HTTPRequest.Header, m, field.Tag)
			case "header":
				err = buildHeader(&r.HTTPRequest.Header, m, name, field.Tag)
			case "uri":
				err = buildURI(r.HTTPRequest.URL, m, name, field.Tag)
			case "querystring":
				err = buildQueryString(query, m, name, field.Tag)
			default:
				// ignore
			}
			r.Error = err
		}
		if r.Error != nil {
			return
		}
	}

	r.HTTPRequest.URL.RawQuery = query.Encode()
	cleanPath(r.HTTPRequest.URL)
}

func buildHeader(header *http.Header, v reflect.Value, name string, tag reflect.StructTag) error {
	str, err := convertType(v, tag)
	if err == errValueNotSet {
		return nil
	} else if err != nil {
		return saperr.New(saperr.Serialization, "failed to encode REST request", err)
	}

	name = strings.TrimSpace(name)
	str = strings.TrimSpace(str)

	header.Add(name, str)

	return nil
}

func buildHeaderMap(header *http.Header, v reflect.Value, tag reflect.StructTag) error {
	prefix := tag.Get(fieldTagDestName)
	for _, key := range v.MapKeys() {
		str, err := convertType(v.MapIndex(key), tag)
		if err == errValueNotSet {
			continue
		} else if err != nil {
			return saperr.New(saperr.Serialization, "failed to encode REST request", err)

		}
		keyStr := strings.TrimSpace(key.String())
		str = strings.TrimSpace(str)

		header.Add(prefix+keyStr, str)
	}
	return nil
}

func buildURI(u *url.URL, v reflect.Value, name string, tag reflect.StructTag) error {
	value, err := convertType(v, tag)
	if err == errValueNotSet {
		return nil
	} else if err != nil {
		return saperr.New(saperr.Serialization, "failed to encode REST request", err)
	}

	u.Path = strings.Replace(u.Path, "{"+name+"}", value, -1)
	u.Path = strings.Replace(u.Path, "{"+name+"+}", value, -1)

	u.RawPath = strings.Replace(u.RawPath, "{"+name+"}", EscapePath(value, true), -1)
	u.RawPath = strings.Replace(u.RawPath, "{"+name+"+}", EscapePath(value, false), -1)

	return nil
}

func buildQueryString(query url.Values, v reflect.Value, name string, tag reflect.StructTag) error {
	switch value := v.Interface().(type) {
	case []*string:
		if len(value) > 0 {
			for _, item := range value {
				if len(*item) > 0 {
					query.Add(name, *item)
				}
			}
		}
	case map[string]*string:
		if len(value) > 0 {
			for key, item := range value {
				if len(*item) > 0 {
					query.Add(key, *item)
				}
			}
		}
	case map[string][]*string:
		if len(value) > 0 {
			for key, items := range value {
				for _, item := range items {
					if len(*item) > 0 {
						query.Add(key, *item)
					}
				}
			}
		}
	case []string:
		if len(value) > 0 {
			query.Add(name, strings.Join(value, ","))
		}
	case map[string]string:
		if len(value) > 0 {
			for key, item := range value {
				query.Add(key, item)
			}
		}
	case map[string][]string:
		if len(value) > 0 {
			for key, items := range value {
				if len(items) > 0 {
					query.Add(key, strings.Join(items, ","))
				}
			}
		}
	case []int:
		if len(value) > 0 {
			items := make([]string, 0)
			for _, item := range value {
				items = append(items, fmt.Sprintf("%d", item))
			}
			query.Add(name, strings.Join(items, ","))
		}
	default:
		str, err := convertType(v, tag)
		if err == errValueNotSet {
			return nil
		} else if err != nil {
			return saperr.New(saperr.Serialization, "failed to encode REST request", err)
		}
		query.Set(name, str)
	}

	return nil
}

func cleanPath(u *url.URL) {
	hasSlash := strings.HasSuffix(u.Path, "/")

	// clean up path, removing duplicate `/`
	u.Path = path.Clean(u.Path)
	u.RawPath = path.Clean(u.RawPath)

	if hasSlash && !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
		u.RawPath += "/"
	}
}

// EscapePath escapes part of a URL path in Amazon style
func EscapePath(path string, encodeSep bool) string {
	var buf bytes.Buffer
	for i := 0; i < len(path); i++ {
		c := path[i]
		if noEscape[c] || (c == '/' && !encodeSep) {
			buf.WriteByte(c)
		} else {
			fmt.Fprintf(&buf, "%%%02X", c)
		}
	}
	return buf.String()
}

func convertType(v reflect.Value, tag reflect.StructTag) (str string, err error) {
	v = reflect.Indirect(v)
	if !v.IsValid() {
		return "", errValueNotSet
	}

	switch value := v.Interface().(type) {
	case string:
		if len(value) <= 0 {
			return "", errValueNotSet
		}
		str = value
	case []byte:
		if len(value) <= 0 {
			return "", errValueNotSet
		}
		str = base64.StdEncoding.EncodeToString(value)
	case bool:
		str = strconv.FormatBool(value)
	case int8, int16, int, int32, int64, uint8, uint16, uint, uint32, uint64:
		str = fmt.Sprintf("%d", value)
	case float64:
		str = strconv.FormatFloat(value, 'f', -1, 64)
	case time.Time:
		format := tag.Get("timestampFormat")
		if len(format) == 0 {
			format = types.RFC822TimeFormatName
			//if tag.Get(fieldTagDest) == "querystring" {
			//	format = types.ISO8601TimeFormatName
			//}
		}
		if value != times.NilTime {
			str = types.FormatTime(format, value)
		} else {
			return "", errValueNotSet
		}
	case types.JSONValue:
		if len(value) == 0 {
			return "", errValueNotSet
		}
		escaping := types.NoEscape
		if tag.Get(fieldTagDest) == "header" {
			escaping = types.Base64Escape
		}
		str, err = types.EncodeJSONValue(value, escaping)
		if err != nil {
			return "", fmt.Errorf("unable to encode JSONValue, %v", err)
		}
	default:
		err := fmt.Errorf("unsupported value for param %v (%s)", v.Interface(), v.Type())
		return "", err
	}
	return str, nil
}
