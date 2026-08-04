
<div align="center">
  <a href="https://github.com/Meastro85/Golem">
    <img width="200" height="200" src="https://raw.githubusercontent.com/Meastro85/Meastro85/refs/heads/main/icons/Golem.png">
  </a>
  <br>
  <br>

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Meastro85/Golem/CI.yml?style=plastic)
![GitHub Repo stars](https://img.shields.io/github/stars/Meastro85/Golem?style=plastic)
![GitHub License](https://img.shields.io/github/license/Meastro85/Golem?style=plastic)

  <h1>Golem</h1>
  <p>
    Golem turns plain Go functions into LLM-callable tools. Give a function a
name, a description, and a typed argument struct, and Golem reflects a
JSON-schema parameter description from it and lets you invoke the function
by name with raw JSON — the exact shape a tool-calling loop needs.
 
Includes an adapter for the Anthropic Messages API that turns a registry
into a `tools` array and drives the full tool-use loop.
  </p>
</div>

## Vision
 
Golem's goal is to become a proper Go semantic kernel — not just a function
registry. Concretely, that means growing toward what a kernel does in the
[Semantic Kernel](https://learn.microsoft.com/en-us/semantic-kernel/) sense:
 
- **Multiple AI service backends**, not just Anthropic — the registry and
  schema generation shouldn't care which provider executes a tool call.
- **Semantic functions** — functions whose "body" is a prompt template
  rather than Go code, registered and invoked through the same interface as
  native functions.
- **Memory** — an embeddings-backed store the kernel can query for
  retrieval-augmented context, not just stateless function execution.
- **Shared context across a call chain** — state that threads through a
  sequence of function invocations, rather than each `Execute` call being
  fully isolated.
- **Invocation middleware** — hooks for logging, monitoring, and
  intercepting calls before/after execution.
**Today, Golem is only the first piece of that: a native-function registry
with generated schemas, plus one provider's tool-use loop.** The sections
below document what exists now. See [Roadmap](#roadmap-toward-a-semantic-kernel)
for what's intentionally not built yet and the shape it's expected to take,
so contributions land in a direction consistent with where this is headed.

## Install
 
```bash
Work in progress
```

## Quick start
 
Every function registered with the kernel has the signature
`func(context.Context, ArgsStruct) (ResultType, error)`. The argument struct
carries the schema:
 
```go
Work in progress
```

## Current supported LLM's
 
The plan is to support a bunch of different LLM's. Due to the current early stages there is limited to none support yet, but actively being worked on.

## Roadmap toward a semantic kernel
 
Ordered roughly by how much they'd change existing APIs if added later —
earlier items are closer to drop-in, later items likely mean breaking
changes to `Kernel` or `Tool`. Treat this as the working plan, not a
promise of order or timeline.
 
1. **Provider abstraction.** Extract an interface the `anthropic` package
   implements, so a second backend (OpenAI, etc.) can sit alongside it
   without touching `kernel.go`. Lowest-risk change — the registry and
   schema generation are already provider-agnostic..
2. **Duplicate-registration policy.** Decide whether `RegisterTool` should
   error on a name collision instead of silently overwriting — small change,
   but a breaking one for anything relying on the current overwrite
   behavior.
3. **Invocation middleware.** Pre/post-execute hooks on `Kernel.Execute` for
   logging, metrics, and auth — additive, shouldn't break existing callers.
4. **Shared context across a call chain.** A context/arguments object
   threaded through a sequence of `Execute` calls, instead of each call
   being fully isolated — needed before multi-step planning makes sense.
5. **Semantic (prompt-template) functions.** Functions whose body is a
   prompt rather than Go code, registered and invoked through the same
   `Tool` interface as native ones — likely means `Tool` grows a
   variant or `buildTool` grows a second code path.
6. **Memory / RAG.** An embeddings-backed store the kernel can query,
   probably as its own package rather than a `Kernel` method, so it stays
   optional for consumers who only want the registry.
