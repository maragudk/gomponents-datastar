# Project Decisions

This document records significant architectural and design decisions made throughout the project's development.

## 2026-09-01: Panic on programmer bugs, don't validate HTML

Decided during the Datastar v1.0.3 upgrade, when a self-review proposed validating attribute key names (panicking on uppercase, empty, or HTML-breaking characters in names passed to `Bind`, `Indicator`, `Ref`, `On`).

Context: modifiers only work in Datastar's attribute key form (`data-bind:foo__prop.checked`), which places caller-supplied names into the HTML attribute name, where gomponents applies no escaping. Validation would have caught malformed names at render time but turned previously-rendering calls into panics.

Decision: this library does not validate attribute names, same as gomponents — names are written verbatim and are the caller's responsibility. Panics are reserved for unambiguous programmer bugs, not for HTML validation: `Nonce("")` panics because an empty nonce makes Datastar throw and never initialize. Facts callers need (kebab-case requirement for key-form names, Datastar's camel conversion) are documented in doc comments, not enforced.
