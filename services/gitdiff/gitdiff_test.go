	var diff = `diff --git a/newfile2 b/newfile2
	var diff2 = `diff --git "a/A \\ B" "b/A \\ B"
	var diff2a = `diff --git "a/A \\ B" b/A/B
	var diff3 = `diff --git a/README.md b/README.md
	assert.NoError(t, diff.LoadComments(issue, user))
	expected := `		<span class="n">run</span><span class="added-code"><span class="o">(</span><span class="n">db</span></span><span class="o">)</span>`