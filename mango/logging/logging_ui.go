package logging

import (
	"fmt"
	"strings"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/infinite-universe/mango/util"
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
}

func InitLogPanel(width, height int) *LogPanel {

	panel := new(LogPanel)

	panel.width, panel.height = width, height

	panel.panelWidth = width
	panel.panelHeight = int(0.25 * float64(panel.height))
	panel.showGameLogs = true
	panel.showEngineLogs = true

	return panel

}

func (panel *LogPanel) RenderPanel() {


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
			// if imgui.TreeNode(logStr) {
			//   imgui.Text(logStr)
			//   imgui.TreePop()
			// }
			if len(panel.filter) > 0 {
				if !panel.containsSearchTerm(logStr) {
					continue
				}
			}
			if panel.showGameLogs && logItem.source == APP {
				imgui.Text(logStr)
			}
			if panel.showEngineLogs && logItem.source == ENGINE {
				imgui.Text(logStr)
			}
		}

		if !panel.freeScroll {
			imgui.SetScrollHereY(1.0)
		}

	}

	imgui.EndChild()

	imgui.End()

}

func (panel *LogPanel) containsSearchTerm(log string) bool {

	return strings.Contains(log, panel.filter)

}

func constructLogString(item *LogItem) string {

	str := "["

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
