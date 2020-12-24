package testkit

func (ctx *TestkitContext) AddAutomaticTxEntry(entry AutomaticTxEntry) {
	ctx.m.Lock()

	if entry.StartedAt == 0 {
		entry.StartedAt = ctx.mantle.GetApp().BaseApp.LastBlockHeight() + 1
	}

	ctx.autoTxs = append(ctx.autoTxs, entry)
	ctx.m.Unlock()
}

func (ctx *TestkitContext) AddAutomaticTxPauseEntry(entry AutomaticTxPauseEntry) {
	ctx.m.Lock()
	ctx.autoTxPauses = append(ctx.autoTxPauses, &entry)
	ctx.m.Unlock()
}
