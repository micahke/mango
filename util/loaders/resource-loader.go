package loaders

import (
	// "fmt"
	"fmt"
	"io/fs"
	"os"
	"sync"

	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/core/settings"
)

var byte_cache map[string][]byte
var cache_mutex sync.Mutex

type file_data struct {
  name string
  data []byte
}

func Init() error {
	byte_cache = make(map[string][]byte)
	err := readFiles()
	if err != nil {
		return err
	}

	return nil
}

func readFiles() error {

	assetsFolder := settings.Settings.ASSET_FOLDER
	// Read the names of all the files in the asset folder
	dir, err := os.ReadDir(assetsFolder)
	if err != nil {
    return fmt.Errorf("No assets folder found!")
	}

	stopwatch := &util.Stopwatch{}
	stopwatch.Start()
  buffer := make(chan *file_data, len(dir))
	for _, file := range dir {
		go readAndSetFile(assetsFolder, file, buffer)
    // readAndSetFileSync(assetsFolder, file)
	}

  for range dir {
    fd := <-buffer
    byte_cache[fd.name] = fd.data
  }
  close(buffer)
	duration := stopwatch.Stop()
	logging.DebugLog("Asset loading took:", duration)

	return nil
}

func readAndSetFileSync(assetsFolder string, file fs.DirEntry) {

	data, err := os.ReadFile(assetsFolder + "/" + file.Name())
	if err != nil {
		logging.DebugLogError("Could not load file:", file.Name())
	}

	byte_cache[file.Name()] = data
}

func readAndSetFile(assetsFolder string, file fs.DirEntry, buffer chan<- *file_data) {

	data, err := os.ReadFile(assetsFolder + "/" + file.Name())
	if err != nil {
		logging.DebugLogError("Could not load file:", file.Name())
	}
  fd := file_data{
    name: file.Name(),
    data: data,
  }
  buffer <- &fd
	// cache_mutex.Lock()
 //  logging.DebugLog(file.Name())
	// byte_cache[file.Name()] = data
	// cache_mutex.Unlock()
}


func GetFileData(filename string) ([]byte, error) {

  data, ok := byte_cache[filename]
  if !ok {
    return nil, fmt.Errorf(fmt.Sprint("Could not read file: ", filename))
  }
  return data, nil
}
