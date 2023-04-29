package res

import (
	"embed"

	"github.com/micahke/mango/logging"
)

type engine_resource struct {
	filename string
	data     []byte
}

//go:embed *
var fs embed.FS

// Load a specific resource from the engine filesystem
// THIS SHOULD ONLY BE USED FROM INSIDE THE ENGINE
// TODO: honestly rethink this system becuase I see it going to badly
func LoadEngineResource(name string) ([]byte, error) {

  logging.DebugLog("Loading resource:", name)

  content, err := fs.ReadFile(name)
  if err != nil {
    logging.DebugLogError(err)
    return nil, err
  }

  // If no error, return the content
  return content, nil

}


