package grip

import "github.com/tychoish/grip/level"

func (self *Journaler) ComposeDefault(m MessageComposer) {
	self.send(self.sender.GetDefaultLevel(), m)
}
func ComposeDefault(m MessageComposer) {
	std.ComposeDefault(m)
}

func (self *Journaler) ComposeEmergency(m MessageComposer) {
	self.send(level.Emergency, m)
}

func ComposeEmergency(m MessageComposer) {
	std.ComposeEmergency(m)
}

func (self *Journaler) ComposeAlert(m MessageComposer) {
	self.send(level.Alert, m)
}

func ComposeAlert(m MessageComposer) {
	std.ComposeAlert(m)
}

func (self *Journaler) ComposeCritical(m MessageComposer) {
	self.send(level.Critical, m)
}

func ComposeCritical(m MessageComposer) {
	std.ComposeCritical(m)
}

func (self *Journaler) ComposeError(m MessageComposer) {
	self.send(level.Error, m)
}

func ComposeError(m MessageComposer) {
	std.ComposeError(m)
}

func (self *Journaler) ComposeWarning(m MessageComposer) {
	self.send(level.Warning, m)
}

func ComposeWarning(m MessageComposer) {
	std.ComposeWarning(m)
}

func (self *Journaler) ComposeNotice(m MessageComposer) {
	self.send(level.Notice, m)
}

func ComposeNotice(m MessageComposer) {
	std.ComposeNotice(m)
}

func (self *Journaler) ComposeInfo(m MessageComposer) {
	self.send(level.Info, m)
}

func ComposeInfo(m MessageComposer) {
	std.ComposeInfo(m)
}

func (self *Journaler) ComposeDebug(m MessageComposer) {
	self.send(level.Debug, m)
}

func ComposeDebug(m MessageComposer) {
	std.ComposeDebug(m)
}
