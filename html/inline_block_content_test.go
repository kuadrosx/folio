// Copyright 2026 Carlos Munoz and the Folio Authors
// SPDX-License-Identifier: Apache-2.0

package html_test

import (
	"strings"
	"testing"

	"github.com/carlos7ags/folio/document"
)

// TestInlineBlockKeepsContent is the regression for an inline-block box
// whose content vanished: the paragraph re-planned the box at exactly its
// measured height, float rounding in the padding arithmetic left the inner
// paragraph a fraction of a point short, and the paragraph deferred its only
// line instead of placing it. The box's background drew but its text did not.
func TestInlineBlockKeepsContent(t *testing.T) {
	htmlStr := `<html><head><style>
		body { margin: 0; font-size: 12px; }
		.box { display: inline-block; padding: 10px; background: #d4f1e4; }
	</style></head><body>
		<span class="box">inline block content</span>
	</body></html>`
	_, r := htmlRoundtrip(t, htmlStr, document.PageSizeLetter)
	page, err := r.Page(0)
	if err != nil {
		t.Fatal(err)
	}
	txt, err := page.ExtractText()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(txt, "inline block content") {
		t.Fatalf("inline-block content missing from page text: %q", txt)
	}
}
