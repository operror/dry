package swarm

import (
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/cli/command/formatter"
	termui "github.com/gizak/termui"
	"github.com/moncho/dry/appui"
	dryformatter "github.com/moncho/dry/docker/formatter"
	drytermui "github.com/moncho/dry/ui/termui"
)

//ServiceRow is a Grid row showing service information
type ServiceRow struct {
	service      swarm.Service
	ID           *drytermui.ParColumn
	Name         *drytermui.ParColumn
	Mode         *drytermui.ParColumn
	Replicas     *drytermui.ParColumn
	Image        *drytermui.ParColumn
	ServicePorts *drytermui.ParColumn

	drytermui.Row
}

//NewServiceRow creats a new ServiceRow widget
func NewServiceRow(service swarm.Service, serviceInfo formatter.ServiceListInfo, table drytermui.Table) *ServiceRow {
	row := &ServiceRow{
		service:      service,
		ID:           drytermui.NewThemedParColumn(appui.DryTheme, service.ID),
		Name:         drytermui.NewThemedParColumn(appui.DryTheme, service.Spec.Name),
		Mode:         drytermui.NewThemedParColumn(appui.DryTheme, serviceInfo.Mode),
		Replicas:     drytermui.NewThemedParColumn(appui.DryTheme, serviceInfo.Replicas),
		Image:        drytermui.NewThemedParColumn(appui.DryTheme, service.Spec.TaskTemplate.ContainerSpec.Image),
		ServicePorts: drytermui.NewThemedParColumn(appui.DryTheme, dryformatter.FormatPorts(service.Spec.EndpointSpec.Ports)),
	}

	row.Height = 1
	row.Table = table
	//Columns are rendered following the slice order
	row.Columns = []termui.GridBufferer{
		row.ID,
		row.Name,
		row.Mode,
		row.Replicas,
		row.ServicePorts,
		row.Image,
	}
	return row

}

//Highlighted marks this rows as being highlighted
func (row *ServiceRow) Highlighted() {
	row.changeTextColor(
		termui.Attribute(appui.DryTheme.Fg),
		termui.Attribute(appui.DryTheme.CursorLineBg))
}

//NotHighlighted marks this rows as being not highlighted
func (row *ServiceRow) NotHighlighted() {
	row.changeTextColor(
		termui.Attribute(appui.DryTheme.ListItem),
		termui.Attribute(appui.DryTheme.Bg))
}

func (row *ServiceRow) changeTextColor(fg, bg termui.Attribute) {
	row.ID.TextFgColor = fg
	row.ID.TextBgColor = bg
}

func serviceMode(service swarm.Service) string {
	if service.Spec.Mode.Global != nil {
		return "global"
	}
	return "replicated"
}
