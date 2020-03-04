/*
 * @Author: dzw
 * @Date: 2020-03-03 19:30:14
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-03 21:30:33
 */

package mylog

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
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

// Config ...
type Config struct {
	ConsoleConfig `ini:"console"`
	FileConfig    `ini:"file"`
}

// LoadConfig ...
func LoadConfig(filename string, conf interface{}) error {
	// 1. 参数校验
	t := reflect.TypeOf(conf)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("param is should be a struct pointer")
	}
	rv := reflect.ValueOf(conf)

	// 2. 读取数据
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// 按行分割数据
	bufSlice := strings.Split(string(buf), "\n")

	// 3. 一行一行读
	structName := ""
	for index, v := range bufSlice {
		// line := strings.ReplaceAll(v, " ", "")    // 去除空格
		line := strings.TrimSpace(v)              // 去除空格
		line = strings.ReplaceAll(line, "\r", "") // 去除\r
		line = strings.ReplaceAll(line, "\n", "") // 去除\n
		// fmt.Println("len line =", len(line))
		if len(line) == 0 { // 空行跳过
			continue
		}
		switch line[0] {
		case '#', ';': // 注释跳过
			continue
		// 4. 如果是[开头的表示是节(section)
		case '[': //section
			if line[len(line)-1] != ']' { // 判断最后一个字符是否是 ]
				return fmt.Errorf("line:%d syntax error", index+1)
			}
			secName := strings.TrimSpace(line[1 : len(line)-1]) // 取出[]中间数据，再去除前后空格
			if len(secName) == 0 {
				return fmt.Errorf("line:%d syntax error", index+1)
			}
			// 根据 secName 去 data 里边根据反射找到对应结构体
			num := t.Elem().NumField()
			for i := 0; i < num; i++ {
				if secName == t.Elem().Field(i).Tag.Get(`ini`) {
					structName = t.Elem().Field(i).Name
					fmt.Println("get", secName, " struct name", structName)
					break
				}
			}
		// 5. 如果不是[开头的表示是=分割的键值对
		default: // key-value
			// 按=分割
			kv := strings.Split(line, "=")
			if len(kv) != 2 {
				return fmt.Errorf("line:%d syntax error", index+1)
			}

			// 去除空格判断是否有效
			key := strings.TrimSpace(kv[0])
			vaule := strings.TrimSpace(kv[1])
			if len(key) == 0 || len(vaule) == 0 {
				return fmt.Errorf("line:%d syntax error", index+1)
			}
			// 赋值到config结构体
			// fmt.Println("len kv ", len(kv))
			structObj := rv.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				return fmt.Errorf("fields %s in data should be struct", structName)
			}

			// structObjType := reflect.TypeOf(structObj)
			structObjType := structObj.Type()
			num := structObjType.NumField()
			for i := 0; i < num; i++ {
				if key == structObjType.Field(i).Tag.Get(`ini`) && structObj.Field(i).CanSet() { // 找到对应节点设置信息
					switch structObj.Field(i).Kind() {
					case reflect.String:
						structObj.Field(i).SetString(vaule)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						if v, err := strconv.ParseInt(vaule, 10, 64); err == nil {
							structObj.Field(i).SetInt(v)
						} else {
							return fmt.Errorf("line:%d param type error", index+1)
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						if v, err := strconv.ParseUint(vaule, 10, 64); err == nil {
							structObj.Field(i).SetUint(v)
						} else {
							return fmt.Errorf("line:%d param type error", index+1)
						}
					case reflect.Bool:
						if v, err := strconv.ParseBool(vaule); err == nil {
							structObj.Field(i).SetBool(v)
						} else {
							return fmt.Errorf("line:%d param type error", index+1)
						}
					}
					break
				}
			}
		}

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
