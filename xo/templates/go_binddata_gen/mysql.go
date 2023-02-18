package tplbin

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _templates_mysql_enum_go_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x3f\x4f\xc3\x30\x10\xc5\xf7\x48\xf9\x0e\x4f\x51\x87\x76\x88\x25\x16\x86\x4a\x9d\x3a\xb0\x75\x01\xc1\x88\x8c\x31\x8d\x21\xb1\x2b\xff\x29\x42\x96\xbf\x3b\x3a\xc7\x51\xe8\x1f\x24\x48\x16\x9f\xef\xee\xf7\xde\x73\x8c\x2d\x16\x9a\x0f\x12\xeb\x0d\xde\x8d\xd2\x4f\xca\x77\x68\x9e\x1b\xb0\x07\xfe\xd2\xcb\x1d\xb5\xd8\xd6\xf4\x61\xd0\xf9\xdc\xa6\x54\x57\x31\x62\x21\xf8\x20\xfb\x5d\xd9\xcc\xc5\x96\x3b\x59\x60\xe3\x50\x8b\x85\xeb\x8c\xf5\xd3\xd4\x5c\xfc\xd8\xce\xc0\x03\x17\x1f\x7c\x2f\x21\x75\x18\xea\xaa\xae\xfc\xd7\x41\xe2\x54\x25\x25\x04\xa5\xfd\xcd\x2d\xf5\x85\xd1\xce\x63\x99\x9d\x58\xae\xf7\x12\xec\x91\xf7\x41\xba\xac\x0c\xe0\x6c\xb9\x4d\x89\xec\xcc\x36\x19\xf1\x2e\x04\x36\x50\xc6\x73\x64\xac\xd4\xaf\x19\xb6\x22\xbd\xb7\xa0\x05\x96\x34\x3f\x67\xb8\x02\x58\xe1\xde\x5b\xa5\xf7\xcb\x15\x5c\x3e\x20\x8e\x6e\x8e\xdc\xe2\x48\x06\xcb\x3d\x31\xe9\xde\x7d\x2a\x2f\x3a\x5c\x82\xc7\x76\x09\x72\x3d\x21\x7d\x82\xb2\xfc\x21\xea\x7a\x5e\x19\xed\x90\x95\x0d\x9a\x18\x59\x4a\x0d\x4e\xd4\xa6\xe0\x54\xa6\xc9\xa8\x95\x3e\x58\x3d\x2e\xd6\x55\xfa\xd7\x9b\xdc\x99\x5f\x5e\xa5\x40\xcf\x11\x6c\x1a\x2f\x42\xf4\x7f\x07\x00\x00\xff\xff\x58\xcf\xac\xaa\xa8\x02\x00\x00")

func templates_mysql_enum_go_tpl() ([]byte, error) {
	return bindata_read(
		_templates_mysql_enum_go_tpl,
		"templates/mysql/enum.go.tpl",
	)
}

var _templates_mysql_repo_go_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xdd\x4f\xe3\x48\x12\x7f\x06\x89\xff\xa1\xce\x5a\x8d\x6c\x94\x6d\xf6\xa4\xd3\x3d\xe4\x94\x93\x08\x81\x39\x4e\x0c\x3b\xbb\xb0\xb7\x0f\xa3\xd1\xd2\x89\x2b\xa1\x75\xc6\x8e\xdb\x1d\x3e\x2e\xe3\xff\xfd\xd4\xdd\xfe\x8c\xdb\x5f\x09\x59\x18\x1e\x20\x76\x77\x57\xfd\xaa\xba\xbe\xc3\x7a\xfd\x23\xfc\x20\xe8\xd4\xc3\x6b\xfa\x80\x67\xf4\x01\x3d\x18\x8e\x60\x26\x3f\x9c\xd1\x08\x81\xdc\xca\x45\xfd\x5b\x6e\x81\x1f\xe3\xf8\xe8\x50\x1d\x8b\xee\x03\x2e\xd4\xbb\xe1\x08\xf2\x87\x4d\x72\xea\x40\x72\x84\xb9\xb7\x2f\x4b\xb5\xdf\x62\xbe\xb0\x20\xa5\xc5\xa9\xbf\xc8\x78\x9d\x05\xde\xea\xc1\x8f\xd4\x22\x00\x80\xdc\xc0\xe6\x80\x21\x24\x4b\x8a\x8d\xc5\x5c\x2b\xdb\x92\x6e\x4b\x19\x8c\x80\x7c\x0c\xd4\xa7\x22\x11\xf4\xdd\x94\x21\xfa\x2e\x68\x5c\x4b\x3a\xfb\x2f\x5d\x20\x70\x5c\x06\xf2\x99\x3d\x2c\x03\x2e\xc0\xd6\xc7\xa2\x10\xac\x05\x13\xf7\xab\x29\x99\x05\x0f\x27\xe8\x2d\x38\x8b\x4e\xa2\x90\x7b\x96\xde\x50\x5c\x5d\x04\xc1\xc2\xc3\x93\x27\xc6\xd1\x3a\x3a\x74\x24\x35\x21\x31\x5c\xae\xd7\x15\xad\xc4\xf1\xaf\xb8\x0c\x22\x26\x02\xfe\x02\xcc\x17\xc8\xe7\x74\x86\xb0\xd6\x54\xdb\x4e\xfc\xb2\x42\xfe\x32\x5e\x31\xcf\x45\xae\x4f\x24\xe7\xfc\x08\xb9\x30\x1e\xb6\x67\xe2\x19\x66\x81\x2f\xf0\x59\x90\x33\xfd\x77\x00\x72\x6b\x7e\x75\x71\x0c\xea\x1c\x31\x52\x38\xe3\x48\x05\x3a\x60\x1f\x37\x6c\x1a\x00\x72\x1e\x70\xa7\x15\xcf\xef\x4c\xdc\xdf\xac\xe6\x73\xf6\xfc\x5a\xc8\x06\x10\x29\x7a\x10\x85\xe4\x26\xf4\xd8\xff\x90\xbf\x16\xd8\xcb\xc9\xaf\x18\xad\x3c\xb1\x5f\xa8\xcc\x17\x7f\xff\x5b\x8e\x49\xa3\xfa\x6d\xe9\x52\x81\x46\x5a\xe3\x97\x0b\x86\x9e\x1b\x99\x51\x31\x57\x01\x4b\x3c\x42\x0a\xdb\x07\xa7\x66\xdb\x4b\x7f\x0d\x48\x77\xd7\x5b\x2f\x24\xfa\xf7\x04\x3d\xdc\x23\x1e\xc5\xb0\x95\xd3\xf8\xe5\x72\xd2\xed\x7e\x1c\xb0\xa7\x41\xe0\x99\x24\xb9\x60\xbe\x7b\xea\x79\xbb\x88\xd2\xa4\xbc\x0b\xe6\x09\xe4\x03\x58\xd2\x05\xf3\xa9\x60\x81\x0f\xc7\x2a\x24\xf9\xd4\x23\x9f\xb3\x97\x0e\xd8\x9a\xc8\x15\x8b\xcc\x6e\x52\xc6\xde\x84\xba\xc5\xfb\xf7\x01\x3f\xf5\x39\x8c\x80\x10\x52\x72\xbc\xce\x62\x1d\x1d\xaa\x9c\xd1\x31\xae\x17\xa3\x74\x35\xc6\x37\xa9\x67\x4c\x23\x54\xa7\xcd\xb7\x3b\x57\x12\x77\x52\xca\x5c\x45\x08\x88\x04\x67\xfe\x22\x0b\x3b\x9b\x0a\x38\x96\x4f\xe8\xe1\x4c\x24\x70\xcb\x37\x79\xea\xba\xb9\x1e\xcd\x90\xc2\x29\x18\x88\xb4\x5b\x54\x03\xe7\x38\x53\x74\x5b\xfe\x8c\x04\x5f\xcd\x44\xaa\xd8\xc9\x18\x32\x5e\x97\x93\xa9\x7e\x59\xba\x8b\x7e\xe9\xb5\x33\x8e\x12\x8f\x1c\x93\x3a\xff\x48\x39\x5c\xe3\x53\x9b\x24\x23\x90\xd5\x03\xb9\xc6\xa7\x1b\x14\x49\x0d\xa2\xde\xdc\x28\x72\xb6\x8f\x4f\x76\x0b\x0d\x67\x00\xd6\xb1\xe5\x0c\xb6\x3a\x5c\x14\xc1\x40\x68\xcc\x7c\x57\x91\x69\xd3\xa0\x33\x80\x2e\x50\xb7\x23\xbe\x01\xb2\xaf\x58\x92\xa9\xf2\xe5\xf9\xca\x9f\x81\xbd\x19\x6b\x38\x1c\xb7\xc1\x7e\x47\x75\x56\x6a\xf4\x1c\xc5\x8a\xfb\x15\x3e\x9c\x74\x2e\xc1\xaa\x20\x07\xe0\x33\x2f\x75\xc4\xbd\x28\xeb\x3d\x15\x81\xa9\x26\xa5\xa7\x22\xe7\xc5\xfc\xae\x7f\x33\x57\x6d\x95\xdd\x4b\x2f\x35\x17\x8b\x47\x93\x92\x35\xe0\x24\xd4\xca\x46\x87\x73\xf8\xcb\x48\xea\x3e\x85\x54\xb8\x60\x9f\xe9\x0a\x41\x2f\x24\xcd\x8d\xaf\x03\x4b\x49\x53\xc3\x51\x93\xb2\xd6\xc9\xc9\x70\xaa\x9a\xb7\x34\x06\xdb\xd2\xe1\xc9\x05\x0f\x1e\xec\xbb\xf5\xba\xda\xfc\xc5\xf1\x5d\x56\x98\x86\x53\xf2\xfb\x3d\x72\xb4\xa3\x90\x9c\x87\x6b\xeb\x8e\xb9\x77\xd6\x10\x98\x1b\x27\xa2\x48\x39\x4c\xaa\x9a\x8c\xc9\x47\x4c\xb4\xf1\xc1\x00\x5d\x26\x93\x8c\xcb\x56\xfa\x48\xd6\xcc\xc4\x7d\xe6\xed\xd3\xa4\xdf\xa0\x55\xa8\xb1\xdc\x8d\x1b\xd6\x98\x6d\xab\xee\x62\x2d\x27\xed\xbc\xed\x72\x53\xdd\xd8\x9c\xa7\x3f\x9a\x6c\xa1\x3f\x97\x14\x07\x65\x42\x69\xd3\x2d\x9f\x1d\xf2\x1f\xea\xad\x70\x2b\x5e\x9b\x4a\x94\xea\x2b\x4c\x2c\x4a\x20\x9a\x20\x64\x16\x96\x68\xb8\x6a\x64\x7a\x41\x65\x8f\xf4\x1e\x4e\xf9\x22\x4a\x3f\x9f\xeb\x68\xa0\x1f\xc8\x6d\x70\x13\x7a\xb6\x93\x9f\xce\x28\x9f\x1b\x2d\xb8\x60\xa9\x3f\x15\x28\xe6\x1b\x0a\x62\x87\x53\x92\xc4\xc9\x1a\x48\x84\x10\x27\xf3\x01\xfd\xe1\xe4\x04\xf8\xca\x87\x50\x6e\x3d\x3a\x3c\xe0\x18\x35\xc4\xaf\xc9\x98\x9c\x3f\xe3\x4c\x7b\xa5\xf2\xbf\x83\x8a\xeb\x1d\x1c\xe4\x70\x95\xcb\x1d\x94\x78\xa1\xe0\x0c\x1f\x11\x98\x2b\x0f\xe7\xc1\x92\x63\x44\xae\x68\x24\xb4\x01\x5e\xba\x76\x5f\xe2\xc9\x82\x24\xa9\x9d\x17\x76\xf3\xde\xef\xa1\xa5\x36\xfa\x34\xa4\x1a\x59\x29\x52\x9f\xe8\x52\xea\xf7\x81\x2e\xbf\xe8\xfa\xfe\x6b\xd6\x64\xac\xb3\x08\xdf\xc9\xab\xd8\xbc\xaa\xc9\x26\xaf\x4a\xee\xcd\xd9\x34\xe7\x0c\xd7\x17\x63\x3c\xf8\x0a\x23\x7d\x3d\x5d\xf9\x54\x7c\xa1\x3c\xbb\x2b\xc5\x37\xad\xde\xfa\xbc\x45\x6e\x50\x7c\xa2\x4b\x3b\xc3\xe8\x34\xe4\x2f\x4d\xfb\x8f\x41\x53\x12\xdb\xf0\x97\xad\xb2\x55\x32\x65\x54\xc9\xf7\x97\xda\x5c\xac\x94\xf9\x83\x39\x68\x57\x68\x8c\xb2\x8f\xad\xf2\x71\x95\xa9\x3a\x96\x0a\x9d\xd2\xb9\xa6\x38\xc8\x20\xec\xa6\x97\x34\x8f\xa7\x54\x95\xf7\x67\x3d\xf8\x5e\x02\xc0\x9f\x36\xa9\x6a\x4c\xda\x27\x27\x10\x85\x5e\x1a\xba\x4d\x66\xde\x94\xc6\x13\x43\xaf\x89\x0b\x3d\xc2\x42\x32\x79\xf7\xb1\x79\xf2\x6e\xf4\xf5\xe1\x6b\xa5\xe9\xcd\xe7\xd8\xec\xb7\x15\x6e\x97\x93\xdc\xd0\xcb\xa9\xf0\xdd\xb9\x76\x9d\x67\xf7\xf1\xeb\x1a\xf9\x21\x6b\x5f\xbe\x4b\x67\xdf\xd9\xd5\xf7\x3c\x04\xb6\xcb\xce\xfc\x47\x53\x85\xd5\x3e\x24\xae\x72\x27\x97\x13\xa7\xa4\x1e\xa5\xb7\x7d\x6a\x65\xeb\x81\x75\xb7\x2e\xa4\x5b\xf8\xb2\x2d\x3a\x13\xec\x11\xad\x01\xcc\xa9\x17\x61\xa1\xe9\x84\x51\x73\xe7\xb9\x27\xf7\x56\x30\x8a\x66\x0b\xa5\x6b\x11\x7c\x85\xaf\xd3\x5a\xbe\xe1\x58\xd8\x30\x19\xaf\x1d\xcf\xb6\xce\x9d\x8a\x23\x37\xd2\x59\xa8\x54\x88\x14\x61\x0e\x4d\x37\x38\x3b\x29\xb7\x34\x05\xfc\xde\x14\x6d\x18\x4b\x6d\x0e\x6f\x34\xaf\xae\x89\x85\xcd\x53\x39\xaa\x46\xdf\xa3\x15\x37\x7e\x37\x9f\x78\x6f\x65\x7b\x89\x2f\x39\x55\xbb\x60\x64\xec\x8c\x0b\xdb\xc3\x69\xea\xcd\xd9\x5c\xff\xd4\x75\xb5\x86\x6d\xb9\x58\x5f\x1f\x93\x3b\x0d\xe5\xce\x1a\xe4\x87\xf5\xc9\x9f\x7d\xd5\xea\xad\x61\x9d\x2d\x9c\x87\x43\xe5\xc9\x31\xc4\xce\x3f\x6a\xe2\xc1\xe6\x4f\xe2\x00\x09\x46\xf3\x3e\x83\x16\x62\x40\x2f\xc2\xfd\x0b\x5d\x52\xf5\xde\x85\xaa\x1a\x87\x92\xb2\xc6\x0c\x76\x12\xd1\x34\x6e\x4a\xa5\x6d\x28\x34\x3b\xa9\xa0\x4d\x7c\x93\x9c\x85\x0a\xb5\xee\x9d\x59\xde\x53\xd7\x65\x82\x05\x3e\xf5\x0a\x92\x27\x82\xa8\x34\x17\x65\x8f\xff\x0e\x98\x9f\x3f\x5d\xe1\x5c\x94\xdf\x7c\xe4\xc1\x6a\x39\x7e\xc9\x5f\xfc\x8b\x3e\x32\x7f\x11\x95\x67\x52\xf5\xf2\xd7\xc9\x1d\x97\xa6\xab\x7d\x5a\x87\xfa\xc0\x60\xb4\xff\xf7\xea\xec\x4d\x06\x51\x1d\x52\x98\xee\xbf\x3a\xb8\x98\x07\x5c\x96\x29\xc9\xec\x71\x38\x4a\x74\x9a\x65\x88\x02\x8c\x50\x4f\xfb\xa8\x1a\x3d\x62\xeb\xd0\x71\xeb\x0b\x86\xd2\xa4\xb1\xc0\xb5\x34\x5d\xdc\x20\xf4\x1a\x45\xcf\x5b\x7c\xc1\xdc\xaf\x82\xa9\x20\x1c\x28\xe1\x73\xfe\xcd\xc5\x49\xcf\xda\xe4\x2d\xf5\x91\x54\x2c\xc3\x11\x7c\xf9\xaa\xcb\x96\x8e\x95\x41\xd6\x69\xa6\x3f\x56\x25\x46\x37\x7d\x23\x50\xb6\xac\xa2\xf3\x17\x54\x51\xd6\xf9\x40\xb3\xa8\x86\x02\x2b\x2d\xbc\x76\xff\xf6\xb2\xff\xff\xde\xf4\xa8\x0e\xdb\xaf\xcb\x63\x91\x80\x2e\xff\x9f\xd2\xd5\xaa\xbb\xff\x57\x4e\x5e\x8b\x77\xb5\xf2\x5d\xf5\xd9\xf6\x75\xf0\xeb\x69\xb6\xa1\x12\xdf\x56\xe3\x69\xe2\x32\x36\xff\x5b\x34\x40\xb5\x05\x3c\x39\xb6\x36\xdb\xa2\xc6\xc0\x9f\xb6\xa8\xcd\x22\xad\xe3\xea\xf7\xa8\x79\x2e\xae\x4a\xd4\x25\x20\xee\x01\x56\xeb\x78\x2a\xe9\x85\xf4\x84\x4a\xde\x25\x99\x50\x41\x7b\x7d\xa7\x2c\x4f\x99\xb8\xb2\x79\xd1\xac\x92\xb6\xe5\xdb\xb7\xc2\x4b\xf2\x19\xf9\x67\xba\xc0\x9a\xc5\xc2\x4a\x81\xab\x02\x79\x1b\x08\xea\x9d\x05\x2b\x5f\xc0\x08\x3c\xf4\xed\x0c\xba\x53\x83\x4f\xe5\xdd\x12\x3e\xd9\x22\xca\xb5\x4f\x28\x68\x1e\x40\xaf\x92\x37\x2e\x15\x34\x93\xa3\xe9\x66\xb7\xb1\xd5\xb3\x9f\x7f\xbb\xbe\xb5\xff\xea\xc0\xe9\x0d\xcc\xa4\x14\x56\x6d\x4d\xb5\xbd\x39\x56\x3a\xd6\x0f\x1f\x94\xaa\x36\xaa\x5f\x07\xfe\x09\x3f\x95\xca\x27\x35\x3c\xcd\x87\xaf\x06\xb0\xaa\x67\x4e\xd6\x55\x91\x49\xad\x72\xc1\xd3\x69\x20\x9a\x2a\xbf\x64\x6d\x86\xeb\x4d\xb6\x11\xf5\x62\x63\xfe\x59\x30\xbe\xf8\xe8\xf0\xff\x01\x00\x00\xff\xff\x35\xae\x67\x66\xa0\x2f\x00\x00")

func templates_mysql_repo_go_tpl() ([]byte, error) {
	return bindata_read(
		_templates_mysql_repo_go_tpl,
		"templates/mysql/repo.go.tpl",
	)
}

var _templates_mysql_table_go_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xc1\x6e\xe3\x36\x10\xbd\x1b\xf0\x3f\xbc\x1a\x45\x2b\x2d\x1c\xf9\x1e\xc0\x87\xad\x77\x93\x4d\x91\x3a\x87\x66\xbb\x87\x20\x68\x68\x6b\x64\x2b\xa1\x49\x99\xa4\x9c\xa6\x82\xfe\xa2\xf7\x5e\xfa\x81\xfd\x84\x82\x94\x64\x2b\x8e\x2c\xdb\x49\x0e\xf5\xc5\xc2\x70\x38\xf3\xe6\x71\x66\x38\xcc\xb2\x13\x7c\x6f\xd8\x84\xd3\x98\x2d\x68\xc4\x16\xc4\x71\x3a\xc4\xd4\x7e\x8c\x98\x26\x04\xd7\xd5\x22\x4e\xf2\xbc\xdb\x49\xd8\xf4\x81\xcd\x08\x6e\x4f\xb7\xd3\xed\xc4\x8b\x44\x2a\x03\xaf\xdb\x01\x00\xbd\x44\x6f\x16\x9b\x79\x3a\x09\xa6\x72\x31\x20\x3e\x53\xb1\x1e\xe8\xa5\xe2\xbd\x42\xa1\xbe\x9a\x3c\xcc\x06\xa4\x94\x54\xba\xd7\xed\xf8\xd6\x98\x79\x4a\x08\x59\xf6\x02\x52\x9e\x43\x1b\x95\x4e\x0d\xb2\x6e\xc7\x62\x56\x4c\xcc\x08\xc1\x48\xf2\x74\x21\xb4\x05\x66\xa5\x71\x04\x26\x42\x04\x17\xfa\xb3\x48\x17\xf0\x68\x89\x60\x2c\xcd\x38\xe5\xdc\xda\x43\xc4\xb8\x26\x1f\x56\xdd\x62\xc9\xb2\x7a\xa0\x85\x2d\x17\x69\x9e\xe3\x43\x96\x21\x38\x97\xd7\x16\x50\x9e\xe3\xee\x5e\x4b\x71\xda\xcb\xb2\x9a\x5a\x9e\xf7\x10\x4e\x5e\x0a\xef\x2c\x18\x10\xd7\x74\x90\xa7\x37\x3a\x3a\x01\x89\x10\x15\x03\xd5\x77\xde\xce\xe6\x59\xcc\x0d\xa9\xbd\x9c\xee\x83\x1e\x0b\x43\x4a\x30\x1e\x14\xf6\xae\xc4\x59\x4c\x3c\xc4\x73\x28\xd6\xc8\xb7\x39\x29\xd2\xb8\xb9\xd5\xcb\xe0\xd7\x25\x8f\xff\x24\x55\x2c\xfc\x2c\x63\xd1\x24\xbf\xa4\xc8\xec\x5a\x3b\x57\x32\x4d\x7e\x7a\x72\x4b\x46\xc5\x62\x56\x88\xbf\xb0\x55\x2c\x66\xdb\x1b\x1c\x11\x51\x2a\xa6\xf0\x22\x77\xa8\xbb\xc8\xf0\x31\xa6\xc7\xe2\xd3\xf3\x8b\xc8\x22\x36\xa5\x2c\xb7\xf4\x58\xfb\x71\x84\x08\xc3\x21\x44\xcc\x2b\x91\xfd\x29\x32\xa9\x12\xf8\xa1\xc5\x74\x56\xb2\x50\xfe\x95\x3b\xa2\xa3\xc0\xad\xcb\xd0\xf3\x51\x44\x5d\x81\x28\xcd\xf5\xee\x6c\x22\x6d\xaa\x35\xcf\xef\x7a\x47\x79\xf8\x45\x86\xe9\x1e\x17\xdb\x1e\x8e\x73\x70\xa1\xc7\x31\xf7\x7c\x4c\xa4\xe4\x5b\x96\x2b\x62\x4b\x7b\xcd\xd9\x78\x98\x97\x8f\x61\xd8\x92\xb2\x5e\xe4\xd4\x5c\xb5\x6d\x65\xaf\x95\xf5\xb1\xaa\x1f\xbd\x5f\xc1\x8c\x82\x2c\xdb\x59\x05\x43\xb0\x24\x21\x11\x7a\x6d\x5a\x7d\x2c\x58\x72\xd3\xe0\xf2\xb6\xe6\x2f\xdb\xa0\x3b\xc5\x2a\xf7\x1d\x1d\xf5\x5a\x3a\x94\x04\x57\x6f\xde\x0a\x9b\x4a\xf0\x5b\x77\x6c\xe2\x2c\x2b\xb5\x16\x53\x21\xe9\x63\xe5\x37\xa6\xef\xa1\x90\x6c\x35\x7b\xf7\xaf\x40\x54\xb4\x81\x1a\x20\x27\xe8\xe3\xbe\x19\xcf\x61\x68\xaa\xfe\xf2\x2a\x44\x9b\xe6\x54\x43\xb5\x16\xbe\x11\x59\xd9\xdd\xbc\xd9\xa4\x2c\xc2\x43\x51\xad\xdb\x62\x0d\x54\x25\xeb\x63\x36\x79\x0b\xa8\xa2\xb7\x7a\xf3\x57\x90\x55\xb5\xe5\x1a\xaa\x52\xd4\xc7\xbc\x19\xd3\x60\x80\x2f\x4c\xcf\x6d\x03\x8a\x35\x52\x4d\x21\x22\xa9\x30\x65\x53\x2b\x0b\x82\x00\x8f\x84\x50\x8a\x1f\x0d\x04\x51\x88\xd8\xc0\xe6\x68\x50\xa4\xe3\x60\x80\xeb\xab\x4f\x57\xa7\x18\x29\x62\x86\x3e\x0b\x13\x9b\x27\x6b\xc8\xda\x98\x29\x96\xcc\x97\x1c\x17\x30\xf3\x58\x3c\xe0\xdf\xbf\xff\xf9\xab\xed\xa2\x2c\x6c\xbc\xf9\xa2\x7c\xb7\x3b\x3e\x7f\x16\xe1\x37\xc2\x9c\xad\x08\x46\x82\xfe\x98\xf2\x34\x24\x7c\x4c\x8d\x3c\x27\x41\x8a\x19\x4b\x9b\xbd\x8f\xb5\xdb\x71\x26\x15\x84\x7c\xc4\x05\xd8\x02\x0f\x44\x89\xa3\xd7\x20\x16\x7d\x30\x0d\x21\x0d\x74\xaa\x08\x73\xf9\x68\xc5\x2c\x8a\x68\x6a\x74\x1b\x37\x5f\x93\xf0\x3d\xb8\xd9\x9a\xb4\x06\x03\x6c\x71\xd0\xc0\x80\x8b\x68\x4e\x3c\x21\x05\x9b\xc1\x26\x96\x42\x57\xc9\x9c\xee\xc8\xcc\x02\xaf\x8f\x6b\xd9\x72\xd4\x9e\x0f\xcf\x36\xc0\x16\x95\x3e\x48\x29\xb8\xc1\xd5\xdf\x3b\x8f\xda\x11\xf4\x42\xd7\x4e\x84\xd5\xe7\xc2\x38\x42\x1a\xb4\x91\xf3\x5d\xc3\xbc\xa1\x5b\x77\x0c\xf1\xa1\xd5\xe4\xb3\x51\xa4\x06\xb2\x3e\x27\x1b\x95\xba\x71\x1f\x95\x56\x0f\xbd\x3c\x77\x13\xed\xcb\xc9\xc7\x5d\x0e\xc5\x18\x1f\x8c\xe9\xd1\xeb\xfd\xc6\x78\x4a\x18\x31\xe1\x92\x6a\x42\x18\x7f\xbd\xbc\xec\x95\xa5\xbe\xb1\x59\x1c\xe9\x49\xfd\x7c\x4f\xb6\x0f\x7b\xdb\xcd\x66\x44\x70\x69\x79\x19\x6b\xb3\xef\xb5\x60\x4d\x5c\x4b\xc3\xf8\x48\xa6\xc2\xa6\xbb\x29\x64\x9f\x98\x61\xb8\xb9\x6d\xdc\xbe\x73\x0c\xa9\x0d\xf4\x47\xbd\x36\x8a\xc4\xe4\xf8\xb0\x13\xb1\x8f\x73\x32\x1f\x39\x6f\x9b\x5d\x7c\xdc\xdc\x6e\x55\x4b\x19\xe0\x8a\x29\x14\x03\xf6\xf3\xf5\xb2\x09\x4b\x85\xdf\xfb\x88\x0d\x2d\xec\xe3\xae\x08\x89\x07\x8e\x81\xe7\x99\xb5\xe9\xd2\x8e\x6e\xbb\xa3\x2d\x97\xfc\xa6\xb9\x56\x91\x76\xfc\xad\x4f\xb9\x4a\xf8\xf7\x23\x61\x0f\x07\xff\x3f\x0a\x1a\x5e\x68\x87\xf1\x51\x3e\x48\x22\x14\xca\x0e\xff\x0e\x55\x3b\x56\x97\xcd\xab\xc5\x60\x76\x14\x1f\xf6\xcd\xe3\x9c\xfa\x75\x69\xd5\x86\x9c\x76\x9d\x2e\x27\x29\xac\xfa\x1b\xf5\xed\xd7\x8f\x0e\x6a\xf5\x38\x04\x27\xb1\xde\xeb\x37\xd2\x78\x28\x55\x76\xb8\x38\x8e\xa8\x46\x8d\x3e\x22\x99\x8a\xb0\xd4\x7b\x3f\xc2\x5c\x4c\x76\xad\xef\x1a\xec\x6e\x82\x9c\x62\x23\xb4\xcc\x82\xb3\x6d\x65\x7d\x0f\x76\x3b\xff\x05\x00\x00\xff\xff\xc1\x89\x58\x88\xc3\x11\x00\x00")

func templates_mysql_table_go_tpl() ([]byte, error) {
	return bindata_read(
		_templates_mysql_table_go_tpl,
		"templates/mysql/table.go.tpl",
	)
}

var _templates_mysql_xo_wire_go_tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\xc1\x6a\xc3\x30\x10\x44\xcf\x15\xe8\x1f\x06\x9f\x6c\x68\xe5\x2f\xe8\xa9\x77\x1f\xea\xde\xcb\xda\x5d\x54\x11\xcb\x2b\x64\x25\x4e\x10\xfa\xf7\xa0\x98\x90\x3d\xec\x61\xe6\xf1\x26\xd0\x7c\x22\xcb\xb8\xca\xef\xee\x22\x6b\xa5\x95\xf3\x41\x62\x42\xab\xd5\x5b\x63\x5d\xfa\x3f\x4f\x66\x16\xdf\x5b\x11\xbb\x70\x5f\xa9\x46\xab\xae\x92\x17\x8a\xf8\xe6\x20\x9b\x4b\x12\x6f\x23\x27\x7c\xa2\xf6\x66\xe0\x7d\xe4\xd4\x6a\x05\x00\x39\x23\xd2\x6a\x19\x06\xa5\x1c\x51\xbd\xc8\x41\x2a\x98\x33\x66\xf2\xbc\x7c\xd1\xc6\x30\x3f\x34\x2d\x7c\xfc\x81\x3c\xa3\x94\xd7\xc0\xfb\xd3\xf7\x01\x5e\xff\x1e\xb2\xee\x1e\x00\x00\xff\xff\x12\x75\x18\xab\xc0\x00\x00\x00")

func templates_mysql_xo_wire_go_tpl() ([]byte, error) {
	return bindata_read(
		_templates_mysql_xo_wire_go_tpl,
		"templates/mysql/xo_wire.go.tpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/mysql/enum.go.tpl": templates_mysql_enum_go_tpl,
	"templates/mysql/repo.go.tpl": templates_mysql_repo_go_tpl,
	"templates/mysql/table.go.tpl": templates_mysql_table_go_tpl,
	"templates/mysql/xo_wire.go.tpl": templates_mysql_xo_wire_go_tpl,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates/mysql/enum.go.tpl": &_bintree_t{templates_mysql_enum_go_tpl, map[string]*_bintree_t{
	}},
	"templates/mysql/repo.go.tpl": &_bintree_t{templates_mysql_repo_go_tpl, map[string]*_bintree_t{
	}},
	"templates/mysql/table.go.tpl": &_bintree_t{templates_mysql_table_go_tpl, map[string]*_bintree_t{
	}},
	"templates/mysql/xo_wire.go.tpl": &_bintree_t{templates_mysql_xo_wire_go_tpl, map[string]*_bintree_t{
	}},
}}
