package logging

import (
	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

// Internal helpers to manage sending interaction

func (g *Grip) conditionalSend(priority level.Priority, conditional bool, m message.Composer) {
	if !conditional {
		return
	}

	g.sender.Send(priority, m)
	return
}

func (g *Grip) conditionalSendPanic(priority level.Priority, conditional bool, m message.Composer) {
	if !conditional {
		return
	}

	g.sendPanic(priority, m)
}

func (g *Grip) conditionalSendFatal(priority level.Priority, conditional bool, m message.Composer) {
	if !conditional {
		return
	}

	g.sendFatal(priority, m)
}

/////////////

func (g *Grip) SendWhen(conditional bool, l level.Priority, m interface{}) {
	g.conditionalSend(l, conditional, message.ConvertToComposer(m))
}
func (g *Grip) SendWhenln(conditional bool, l level.Priority, msg ...interface{}) {
	g.SendWhen(conditional, l, message.NewLinesMessage(msg...))
}
func (g *Grip) SendWhenf(conditional bool, l level.Priority, msg string, args ...interface{}) {
	g.SendWhen(conditional, l, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) DefaultWhen(conditional bool, m interface{}) {
	g.conditionalSend(g.sender.DefaultLevel(), conditional, message.ConvertToComposer(m))
}
func (g *Grip) DefaultWhenln(conditional bool, msg ...interface{}) {
	g.DefaultWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	g.DefaultWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) EmergencyWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Emergency, conditional, message.ConvertToComposer(m))
}
func (g *Grip) EmergencyWhenln(conditional bool, msg ...interface{}) {
	g.EmergencyWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	g.EmergencyWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) AlertWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Alert, conditional, message.ConvertToComposer(m))
}
func (g *Grip) AlertWhenln(conditional bool, msg ...interface{}) {
	g.AlertWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) AlertWhenf(conditional bool, msg string, args ...interface{}) {
	g.AlertWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) CriticalWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Critical, conditional, message.ConvertToComposer(m))
}
func (g *Grip) CriticalWhenln(conditional bool, msg ...interface{}) {
	g.CriticalWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	g.CriticalWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) ErrorWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Error, conditional, message.ConvertToComposer(m))
}
func (g *Grip) ErrorWhenln(conditional bool, msg ...interface{}) {
	g.ErrorWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	g.ErrorWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) WarningWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Warning, conditional, message.ConvertToComposer(m))
}
func (g *Grip) WarningWhenln(conditional bool, msg ...interface{}) {
	g.WarningWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) WarningWhenf(conditional bool, msg string, args ...interface{}) {
	g.WarningWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) NoticeWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Notice, conditional, message.ConvertToComposer(m))
}
func (g *Grip) NoticeWhenln(conditional bool, msg ...interface{}) {
	g.NoticeWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	g.NoticeWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) InfoWhen(conditional bool, msg interface{}) {
	g.conditionalSend(level.Info, conditional, message.ConvertToComposer(msg))
}
func (g *Grip) InfoWhenln(conditional bool, msg ...interface{}) {
	g.InfoWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) InfoWhenf(conditional bool, msg string, args ...interface{}) {
	g.InfoWhen(conditional, message.NewFormatedMessage(msg, args...))
}

/////////////

func (g *Grip) DebugWhen(conditional bool, m interface{}) {
	g.conditionalSend(level.Debug, conditional, message.ConvertToComposer(m))
}
func (g *Grip) DebugWhenln(conditional bool, msg ...interface{}) {
	g.DebugWhen(conditional, message.NewLinesMessage(msg...))
}
func (g *Grip) DebugWhenf(conditional bool, msg string, args ...interface{}) {
	g.DebugWhen(conditional, message.NewFormatedMessage(msg, args...))
}
