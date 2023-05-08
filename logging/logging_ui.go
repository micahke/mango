package logging

import (
	"fmt"
	"strings"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/util"
)

type LogPanel struct {
	showPanel bool

	width  int
	height int

	panelWidth  int
	panelHeight int

	filter string

	showGameLogs   bool
	showEngineLogs bool

	freeScroll bool

	errorColor imgui.Vec4
}

func InitLogPanel(width, height int) *LogPanel {

	panel := new(LogPanel)

	panel.width, panel.height = width, height

	panel.panelWidth = 400
	panel.panelHeight = height / 2
	panel.showGameLogs = true
	panel.showEngineLogs = true

	panel.errorColor = imgui.Vec4{
    X: 1.0,
    Y: 0.0,
    Z: 0.0,
    W: 1.0,
  }

	return panel

}

func (panel *LogPanel) RenderPanel() {

	imgui.SetNextWindowSizeV(imgui.Vec2{
		X: float32(panel.panelWidth),
		Y: float32(panel.panelHeight),
	}, imgui.ConditionOnce)
	imgui.BeginV("Mango Log", util.ImguiPanelStatus("logPanel"), 0)

	{

		if imgui.Button("Clear") {
			_log = []*LogItem{}
		}

		imgui.SameLine()

		imgui.SameLine()

		imgui.InputText("Filter", &panel.filter)

		imgui.Checkbox("Game Logs", &panel.showGameLogs)
		imgui.SameLine()
		imgui.Checkbox("Engine Logs", &panel.showEngineLogs)
		imgui.SameLine()
		imgui.Checkbox("Free Scroll", &panel.freeScroll)
	}

	imgui.Separator()

	imgui.BeginChild("Logs")

	{
		iterator := _log
		if len(iterator) > 500 {
			iterator = iterator[len(iterator)-500:]
		}

		for _, logItem := range iterator {
			logStr := constructLogString(logItem)
			if len(panel.filter) > 0 {
				if !panel.containsSearchTerm(logStr) {
					continue
				}
			}
			if panel.showGameLogs && logItem.source == APP {
				panel.DrawLogText(logStr, logItem.logType)
			}
			if panel.showEngineLogs && logItem.source == ENGINE {
				panel.DrawLogText(logStr, logItem.logType)
			}
		}

		if !panel.freeScroll {
			imgui.SetScrollHereY(1.0)
		}

	}

	imgui.EndChild()

	imgui.End()

}

func (panel *LogPanel) DrawLogText(text string, logType LogType) {

	if logType == LOG_ERROR {
		imgui.PushStyleColor(imgui.StyleColorText, panel.errorColor)
		imgui.Text(text)
		imgui.PopStyleColor()
	} else {
		imgui.Text(text)
	}

}

func (panel *LogPanel) containsSearchTerm(log string) bool {

	return strings.Contains(log, panel.filter)

}

func constructLogString(item *LogItem) string {

	str := "["

  if item.source == ENGINE {
    str += "ENGINE "
  }

  if item.source == APP {
    str += "APPLICATION "
  }

	// add log type
	if item.logType == LOG_LOG {
		str += "LOG"
	}
	if item.logType == LOG_ERROR {
		str += "ERROR"
	}
	str += "] "

	for _, arg := range item.content {
		str += fmt.Sprint(arg) + "  "
	}

	return str

}
