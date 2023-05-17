package core

type EngineSettings struct {
	SCENE_EDITOR_STARTUP bool
	CONSOLE_ON_STARTUP   bool
  SHADER_EDITOR_ON_STARTUP bool
}

// Defaults
var Settings *EngineSettings = &EngineSettings{
	SCENE_EDITOR_STARTUP: false, // Decides whether the scene editor is displaye
  CONSOLE_ON_STARTUP: false,
  SHADER_EDITOR_ON_STARTUP: false,
}
