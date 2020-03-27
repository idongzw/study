/*
 * @Author: dzw
 * @Date: 2020-03-03 19:30:14
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-03 21:30:33
 */

package chatlog

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

// ConsoleConfig ...
type ConsoleConfig struct {
	Flag  bool   `ini:"flag"`
	Level string `ini:"level"`
}

// FileConfig ...
type FileConfig struct {
	Flag     bool   `ini:"flag"`
	Level    string `ini:"level"`
	Filepath string `ini:"filepath"`
	Maxsize  string `ini:"maxsize"`
}

// Buffer ...
type Buffer struct {
	BufSize int `ini:"bufsize"`
}

// Config ...
type Config struct {
	ConsoleConfig `ini:"console"`
	FileConfig    `ini:"file"`
	Buffer        `ini:"buffer"`
}

// LoadConfig ...
func LoadConfig(filename string, conf *Config) error {
	// load log config from ini file
	file, err := ini.Load(filename)
	if err != nil {
		return err
	}
	// map config to struct
	err = file.MapTo(conf)
	if err != nil {
		return err
	}

	return nil
}

// GetMaxBytesSize ...
func GetMaxBytesSize(maxStr string) (max int64, err error) {
	num := len(maxStr)
	maxStr = strings.ToUpper(maxStr)

	// KB
	if strings.HasSuffix(maxStr, "KB") && num >= 3 {
		maxStr = maxStr[:num-2]
		max, err = strconv.ParseInt(maxStr, 10, 64)
		max *= 1024
		return
	}

	// MB
	if strings.HasSuffix(maxStr, "MB") && num >= 3 {
		maxStr = maxStr[:num-2]
		max, err = strconv.ParseInt(maxStr, 10, 64)
		max *= 1024 * 1024
		return
	}

	// GB
	if strings.HasSuffix(maxStr, "GB") && num >= 3 {
		maxStr = maxStr[:num-2]
		max, err = strconv.ParseInt(maxStr, 10, 64)
		max *= 1024 * 1024 * 1024
		return
	}

	if strings.HasSuffix(maxStr, "B") && num >= 2 {
		maxStr = maxStr[:num-1]
		max, err = strconv.ParseInt(maxStr, 10, 64)
		return
	}

	return 0, fmt.Errorf("%s is invalid", maxStr)
}
