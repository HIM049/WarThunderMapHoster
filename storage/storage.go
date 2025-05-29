package storage

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"os"
	"path/filepath"
	"thunder_hoster/config"
)

var ErrDuplicatedName = errors.New("duplicated name")

type MapStorage struct {
	Maps            []MapInformation
	ListMap         map[string]*MapInformation
	StorageFilePath string
}

type MapInformation struct {
	MapName    string `json:"map_name"`
	FilePath   string `json:"file_path"`
	UpdateTime string `json:"update_time"`
}

const STORAGE_FILE_NAME = "storage.json"

var Storage MapStorage

// InitStorage 初始化存储
func InitStorage() {
	err := os.MkdirAll(config.Cfg.MapDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create map directory: %v", err)
	}

	Storage.ListMap = make(map[string]*MapInformation)
	Storage.StorageFilePath = filepath.Join(config.Cfg.MapDir, STORAGE_FILE_NAME)

	err = RefreshStorage()
	if err != nil {
		log.Fatalf("Failed to refresh storage: %v", err)
	}
}

func (m *MapStorage) Add(newMap *MapInformation) error {
	if _, found := m.ListMap[newMap.MapName]; found {
		return ErrDuplicatedName
	}
	m.Maps = append(m.Maps, *newMap)

	err := m.SaveToFile()
	if err != nil {
		return err
	}

	return nil
}

func (m *MapStorage) Remove(name string) error {
	for i, amap := range m.Maps {
		if amap.MapName == name {
			err := os.Remove(amap.FilePath)
			if err != nil {
				return err
			}

			m.Maps = append(m.Maps[:i], m.Maps[i+1:]...)
			break
		}
	}

	err := m.SaveToFile()
	if err != nil {
		return err
	}

	return nil
}

// RefreshStorage 刷新存储
func RefreshStorage() error {
	err := Storage.ReadFromFile()
	if err != nil {
		return err
	}

	Storage.GenerateIndex()

	return nil
}

// SaveToFile 保存地图存储到文件
func (m *MapStorage) SaveToFile() error {
	storageJson, err := json.Marshal(Storage.Maps)
	if err != nil {
		return err
	}

	if err := os.WriteFile(Storage.StorageFilePath, storageJson, 0644); err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}

	return nil
}

// ReadFromFile 从文件中读取存储
func (m *MapStorage) ReadFromFile() error {
	var storageMaps []MapInformation
	// 打开文件
	storageFile, err := os.OpenFile(m.StorageFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer storageFile.Close()

	// 获取文件信息
	info, err := storageFile.Stat()
	if err != nil {
		return err
	}
	// 如果文件为不空则操作
	if info.Size() != 0 {
		storageData, err := io.ReadAll(storageFile)
		if err != nil {
			return err
		}

		err = json.Unmarshal(storageData, &storageMaps)
		if err != nil {
			return err
		}
	}

	m.Maps = storageMaps
	return nil
}

// GenerateIndex 创建地图索引
func (m *MapStorage) GenerateIndex() {
	if len(m.Maps) == 0 {
		return
	}

	for _, mapInfo := range m.Maps {
		m.ListMap[mapInfo.MapName] = &mapInfo
	}
}
