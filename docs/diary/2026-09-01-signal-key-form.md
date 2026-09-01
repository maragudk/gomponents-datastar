# Diary: Add `datastar.Signal` and fix `datastar.Signals` case-modifier docs

Follow-up to the v1.0.3 upgrade (PR #46): `Signals` accepts case modifiers that Datastar silently ignores in the value form. Add a key-form single-signal function where modifiers actually work, and stop documenting/testing the dead combination.

## Step 1: Requirements

**Author:** main

### Prompt Context

**Verbatim prompt:** "Both, make a PR" (in response to the recommendation: document that case modifiers don't apply to `Signals`, delete the test pinning the dead output, and add a new `Signal(name string, value any, modifiers ...Modifier)` emitting the key form).
**Interpretation:** Implement both the `Signals` doc/test cleanup and the new `Signal` function, then open a PR — before tagging v0.4.0, so the breaking release ships a complete signals story.
**Inferred intent:** Don't release API surface that implies behavior Datastar doesn't deliver; give users a working way to apply modifiers to a single signal.

### What I did

Created worktree `signal-key-form`, started this diary, and handed the builder the requirements: `Signal(name string, value any, modifiers ...Modifier) g.Node` emitting `data-signals:name__mods="<JSON value>"`; `Signals` doc comment noting case modifiers only apply in the key form; delete the `__case.kebab` test on `Signals`; keep `ModifierIfMissing` support as-is.

### Why

Same rationale as the `Bind`/`Indicator`/`Ref` fix in #46: modifiers are parsed from the attribute key only. `Signals` can't move to the key form (it takes a whole map; the key form addresses one signal per attribute), so the honest fix is documentation plus a purpose-built single-signal function.

### What worked

The v1.0.3 upgrade research and parser findings from #46 carried over directly; no new research was needed.

### What didn't work

Nothing failed; this step was requirements only.

### What I learned

The key form and value form of `data-signals` aren't interchangeable representations — they have different modifier semantics, so a Go API mirroring Datastar needs both shapes.

### What was tricky

Nothing in this step.

### What warrants review

The `Signal` emission shape (key form, JSON-encoded value) and that the `Signals` doc change matches Datastar's actual parser behavior.

### Future work

Tag v0.4.0 after this merges.

## Step 2: Add `Signal` and document the dead modifiers on `Signals`

**Author:** builder-datastar-103

### Prompt Context

**Verbatim prompt:** "This addresses the Signals latent bug you flagged: case modifiers are dead in the value form. Two changes: 1. New function `Signal(name string, value any, modifiers ...Modifier) g.Node` emitting the key form: `data-signals:foo=\"1\"` and e.g. `Signal(\"my-signal\", 1, ModifierCase, ModifierKebab)` → `data-signals:my-signal__case.kebab=\"1\"`. JSON-encode the value the way Signals does its map values (reuse/adapt the existing helper as fits). Modifiers genuinely work here (case, __ifmissing). Doc comment verbatim from the data-signals section [...] document the key-form casing facts (kebab-case names, Datastar's camel conversion) consistent with the Bind/Indicator/Ref comments you wrote. Place it alphabetically near Signals. 2. `Signals` cleanup: add to its doc comment that case modifiers only apply to the key form and are ignored on Signals' value form (point to Signal for that); delete the test pinning `data-signals__case.kebab=\"{...}\"`. Keep ModifierIfMissing support and its test [...] No validation/panics, per the decision in docs/decisions.md."

**Interpretation:** Give users a working way to apply casing modifiers to a single signal, and stop the library implying `Signals` supports something Datastar ignores.

**Inferred intent:** Ship v0.4.0 with a signals API that is honest about which form supports which modifier.

### What I did

Added `Signal` to `/datastar.go` between `Show` and `Signals`, emitting `data-signals:<name><modifiers>` with the value JSON encoded. Renamed the unexported `toSignals(map[string]any)` helper to `toJSON(any)` so both functions share it, and generalized its panic message from "failed to marshal signals" to "failed to marshal value".

Extended `Signals`' doc comment to say the casing modifiers are still emitted but never applied in the value form, pointing at `Signal`, and noting `ModifierIfMissing` works in both. Deleted the test pinning `data-signals__case.kebab="{…}"`; kept the `__ifmissing` test.

`TestSignal` in `/datastar_test.go` covers the plain key form, dot-notation nesting, `__ifmissing`, all four casing modifiers as a table, casing combined with `__ifmissing`, and JSON values: string, string containing a quote, a `$`, a backslash, a newline, nil, bool, float, nested object, array, and a value that cannot be encoded. Examples for `Signal` plain, with casing, and with `__ifmissing`.

### Why

`Signals` cannot move to the key form — it takes a whole map, and the key form addresses one signal per attribute — so the honest fix is a separate function plus documentation, exactly as Step 1 framed it.

### What worked

Reading the plugin out of the vendored bundle settled the premise before I wrote anything:

```js
m({name:"signals",returnsValue:!0,apply({key:e,mods:t,rx:n}){
  let r=t.has("ifmissing");
  if(e){e=H(e,t); ...}else{...}}})
```

`r` is read before the branch and passed to both, so `__ifmissing` works in either form, while `H` — the case conversion — is only reached when a key is present. That is the whole bug in one line, and it confirmed which modifier to keep testing on `Signals`.

Every test passed on the first run, including the escaping expectations I guessed at.

### What didn't work

Nothing broke. The one dead end was cosmetic: I first wrote an `ExampleZZProbe` to inspect tricky values and got

```
./zz_probe_test.go:11:1: ExampleZZProbe refers to unknown identifier: ZZProbe
```

because Go only accepts `Example` functions named after a real identifier. Converted it to a `Test` and deleted it afterwards.

### What I learned

The interesting question this change raises is whether JSON is safe to hand Datastar, since the attribute value is evaluated as a JavaScript expression rather than parsed. Three findings, all confirmed against the bundle:

- Datastar wraps the last statement as `return (<expr>);` — parenthesized — so a JSON object literal evaluates as an object rather than a block. That is what makes `Signal("foo", map[string]any{...})` work at all.
- The `$signal` substitution regex leads with alternatives matching complete quoted strings and returns them unchanged, so a `$` inside a JSON string value is protected. `Signal("foo", "$bar")` really does stay the literal text, and there is now a test pinning that.
- The `@action(` rewrite, `s.replaceAll(/@([A-Za-z_$][\w$]*)\(/g,'__action("$1",evt,')`, has no such protection, so a string value containing `@name(` is rewritten and produces a syntax error. This is upstream and pre-existing: `Signals` produces byte-identical expression text and fails the same way.

Also worth remembering: Go's `encoding/json` escapes `<`, `>`, `&` and U+2028/U+2029, which happens to be exactly what makes the output safe both in an HTML attribute and as a JS string literal.

### What was tricky

My first doc comment contradicted itself, and a review agent caught it. I wrote "my-signal becomes $mySignal" a few lines above the docs' own example `data-signals:my-signal__case.kebab="1"`. But the case table in the bundle is `Gt={camel,snake,pascal}` — there is no `kebab` entry — and `H` does `for(let r of t.get("case")||["camel"]) e=Gt[r]?.(e)||e`. So `__case.kebab` replaces the camel default with a lookup that misses and falls through to identity, producing a signal literally named `my-signal`. The example demonstrates the exact opposite of the sentence above it. The comment now says `ModifierCase` replaces the default conversion rather than adding to it, and that `ModifierKebab` leaves the name as written.

The same review also caught that `ModifierIfMissing` cancels nil-removal — `e[i]==null ? s||delete n[t][i] : ...`, where `s` is `ifMissing`, so the delete is short-circuited away. `Signal("foo", nil, ModifierIfMissing)` is a silent no-op, now documented.

### What warrants review

The `Signal` doc comment is longer than its neighbours because it carries three facts the Datastar docs state in separate places: the casing rules, the `__case` override behaviour, and what JSON encoding does and does not survive. Worth a look for whether that is the right amount.

The claim that a string containing `@name(` does not survive is mine, verified against the bundle rather than taken from the Datastar docs, which do not mention it. It is stated on `Signal` but not on `Signals`, which has the same problem.

### Future work

`Signals` still has no test for two modifiers concatenating, because deleting the `__case.kebab` test removed the only one and `__ifmissing` is the sole modifier that works there. The demo does not exercise `Signal`, unlike `Prop`/`Event` which got one in `ea7e64a`. Both are noted for the lead rather than acted on.

## Step 3: Demo `Signal` and verify it in a browser

**Author:** builder-datastar-103

### Prompt Context

**Verbatim prompt:** "1. Add a `Signal` example to the demo in demo/main.go (new API gets demo coverage, per Markus's standard from #46). Show the key form working end-to-end, ideally visibly distinguishable from value-form behavior. Regenerate docs/index.html. Browser-test with playwright-cli: new example works, spot-check existing ones, console free of Datastar errors. 2. No new Signals concatenation test — decided, nothing to do there."

**Interpretation:** Give `Signal` the same live demo coverage `Prop`/`Event` got, choosing an example where the key form's behaviour is observably different from what the value form would do, and confirm it against the real bundle in a browser.

**Inferred intent:** Before tagging v0.4.0, prove in a browser that the function added to fix a silently-ignored modifier does not itself silently do nothing.

### What I did

Added a "Single Signals" section to `/demo/main.go`, placed right after the existing "Signals" section so the two forms sit together. Two boxes differing only by a modifier: `data.Signal("my-count", 1)` labelled as creating `$myCount`, and `data.Signal("my-total", 2, data.ModifierCase, data.ModifierSnake)` labelled as creating `$my_total`. Each displays its signal with `data.Text` and has an Increment button. Regenerated `/docs/index.html`, then served `/docs` and drove it with playwright-cli.

Item 2 needed no work.

### Why

The point of `Signal` is that casing modifiers apply here and not on `Signals`, so the demo had to make that difference visible rather than just show a signal existing. Two boxes that differ only in the modifier do that: `my-count` and `my-total` are the same shape, so the only reason their signal names diverge is the modifier.

### What worked

Everything, and the "All Signals (JSON)" dump gave the cleanest possible evidence:

```
{ ..., "myCount": 3, "my_total": 3, "total": 200, "myDiv": {} }
```

`my-count` became `myCount` by the default camel conversion and `my-total` became `my_total` because `__case.snake` was applied. Had the modifier been ignored the way it is in the value form, that key would read `myTotal`. Both boxes rendered their initial values (1 and 2) on load and incremented live to 3, so the signals are real and reactive, not just present.

Existing demos were unaffected: the counter incremented, text binding rendered "Hello, Markus!", the select reported `option2`, the `__prop.checked` checkbox and the show toggle both worked. Across the session the console held one error, a `404` for `/favicon.ico` from my own static server, and no Datastar errors.

### What didn't work

Nothing. Port 8080 was still occupied by an unrelated `app` process as in Step 4 of the previous diary, so I served `/docs` on 8099 again rather than using the demo's server mode — which tests the published artifact anyway. Two `playwright-cli` friction points recurred and were already known: `eval` is refused in a worktree-isolated session, and `snapshot --filename=` writes relative to the working directory, so I deleted the stray files and `.playwright-cli/` afterwards.

### What I learned

A demo of a modifier is only worth anything if the modifier's absence would change what you see. Showing `data.Signal("foo", 1)` alone would have proved the key form parses but said nothing about the bug this change exists to fix, because a signal named `foo` looks identical either way. Pairing two calls whose names differ only through the modifier is what makes the JSON dump a proof rather than a screenshot.

### What was tricky

Nothing this step. The choice of `snake` over `kebab` for the second box mattered, though: `__case.kebab` is an identity conversion in Datastar, so a kebab example would have rendered a name indistinguishable from the raw key and demonstrated nothing — the same trap that produced the contradicting doc comment in Step 2.

### What warrants review

Whether "Single Signals" is the right heading next to the existing "Signals" section, and whether two boxes is the right size for this. The two new signals appear in the page-wide "All Signals (JSON)" box, which is intended but makes that box longer.

### Future work

`Signals` still has no multi-modifier concatenation test, by decision. Tag v0.4.0 after this merges.
