package main

func segmentFiller(p *powerline) {
	p.appendSegment("filler", segment{
		content:        "                     . . .                                   ",
		foreground:     p.theme.CwdFg,  // FIXME
		background:     p.theme.PathBg, // FIXME
		hideSeparators: true,
	})
}
