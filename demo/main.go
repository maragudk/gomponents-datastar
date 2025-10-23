package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	g "maragu.dev/gomponents"
	ds "maragu.dev/gomponents-datastar"
	. "maragu.dev/gomponents/html"
)

func main() {
	generateOnly := flag.Bool("generate", false, "Generate docs/index.html and exit")
	flag.Parse()

	if *generateOnly {
		// Generate mode: write HTML and exit
		if err := writeDocsHTML(); err != nil {
			log.Fatalf("Error: Failed to write docs/index.html: %v\n", err)
		}
		fmt.Println("✓ Generated docs/index.html")
		return
	}

	// Server mode
	http.HandleFunc("/", handleIndex)

	const addr = ":8080"
	fmt.Printf("Demo server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func writeDocsHTML() error {
	// Create docs directory if it doesn't exist
	docsDir := filepath.Join("..", "docs")
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("creating docs directory: %w", err)
	}

	// Create the file
	indexPath := filepath.Join(docsDir, "index.html")
	f, err := os.Create(indexPath)
	if err != nil {
		return fmt.Errorf("creating index.html: %w", err)
	}
	defer f.Close()

	// Write the HTML
	page := buildPage()
	if err := page.Render(f); err != nil {
		return fmt.Errorf("rendering HTML: %w", err)
	}

	return nil
}

func buildPage() g.Node {
	return HTML(
		Lang("en"),
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			TitleEl(g.Text("Datastar Attributes Demo")),
			Script(Type("module"), Src("https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js")),
			StyleEl(Type("text/css"), g.Raw(`
				body {
					font-family: 'Comic Sans MS', 'Comic Neue', system-ui, -apple-system, sans-serif;
					max-width: 1200px;
					margin: 0 auto;
					padding: 20px;
					background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
					min-height: 100vh;
				}
				h1 {
					font-size: 3em;
					text-align: center;
					color: white;
					text-shadow: 3px 3px 0 rgba(0,0,0,0.2);
					margin-bottom: 10px;
				}
				body > p:first-of-type {
					text-align: center;
					color: white;
					font-size: 1.2em;
					margin-bottom: 30px;
				}
				.demo-section {
					background: white;
					border-radius: 20px;
					padding: 25px;
					margin-bottom: 25px;
					box-shadow: 8px 8px 0 rgba(0,0,0,0.2);
					border: 4px solid #333;
				}
				.demo-section h2 {
					margin-top: 0;
					color: #ff6b6b;
					border-bottom: 4px dashed #4ecdc4;
					padding-bottom: 10px;
					font-size: 2em;
				}
				.demo-section h3 {
					color: #5f27cd;
					margin-top: 20px;
					font-size: 1.3em;
				}
				.demo-box {
					border: 3px solid #333;
					padding: 18px;
					margin: 15px 0;
					border-radius: 15px;
					background: linear-gradient(135deg, #ffeaa7 0%, #fdcb6e 100%);
					box-shadow: 4px 4px 0 rgba(0,0,0,0.1);
				}
				button {
					background: linear-gradient(135deg, #00b894, #00cec9);
					color: white;
					border: 3px solid #333;
					padding: 12px 24px;
					border-radius: 12px;
					cursor: pointer;
					margin: 5px;
					font-weight: bold;
					font-size: 1em;
					box-shadow: 4px 4px 0 rgba(0,0,0,0.2);
					transition: transform 0.1s, box-shadow 0.1s;
				}
				button:hover {
					transform: translate(-2px, -2px);
					box-shadow: 6px 6px 0 rgba(0,0,0,0.2);
				}
				button:active {
					transform: translate(2px, 2px);
					box-shadow: 2px 2px 0 rgba(0,0,0,0.2);
				}
				button:disabled {
					background: #dfe6e9;
					cursor: not-allowed;
					transform: none;
				}
				input, select, textarea {
					padding: 10px;
					border: 3px solid #333;
					border-radius: 8px;
					margin: 5px;
					font-size: 1em;
					background: white;
					transition: transform 0.1s;
				}
				input:focus, select:focus, textarea:focus {
					outline: none;
					transform: scale(1.05);
					border-color: #ff6b6b;
				}
				.hidden {
					display: none;
				}
				.highlight {
					background: #ffeb3b;
					animation: wiggle 0.5s ease-in-out;
				}
				@keyframes wiggle {
					0%, 100% { transform: rotate(0deg); }
					25% { transform: rotate(-3deg); }
					75% { transform: rotate(3deg); }
				}
				.bold {
					font-weight: bold;
				}
				pre {
					background: #2d2d2d;
					color: #f8f8f2;
					padding: 15px;
					border-radius: 12px;
					overflow-x: auto;
					border: 3px solid #333;
					box-shadow: 4px 4px 0 rgba(0,0,0,0.3);
				}
				.grid {
					display: grid;
					grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
					gap: 15px;
				}
				.loading-spinner {
					border: 3px solid #f3f3f3;
					border-top: 3px solid #007bff;
					border-radius: 50%;
					width: 20px;
					height: 20px;
					animation: spin 1s linear infinite;
					display: inline-block;
					margin-left: 10px;
				}
				@keyframes spin {
					0% { transform: rotate(0deg); }
					100% { transform: rotate(360deg); }
				}
			`)),
		),
		Body(
			H1(g.Text("Datastar Attributes Demo")),
			P(g.Text("This page demonstrates all the Datastar attributes available in gomponents-datastar.")),

			// Signals Demo
			Div(Class("demo-section"),
				ds.Signals(map[string]any{
					"counter":     0,
					"name":        "World",
					"show":        true,
					"isHighlit":   false,
					"price":       100,
					"quantity":    2,
					"inputValue":  "Hello",
					"selected":    "option1",
					"message":     "",
					"count":       0,
					"intersected": false,
				}),
				H2(g.Text("Signals")),
				P(g.Text("Signals are reactive state variables. Click the buttons to modify them.")),
				Div(Class("demo-box"),
					H3(g.Text("Counter Example")),
					P(ds.Text("$counter"), g.Text(" clicks")),
					Button(
						ds.On("click", "$counter++"),
						g.Text("Increment"),
					),
					Button(
						ds.On("click", "$counter--"),
						g.Text("Decrement"),
					),
					Button(
						ds.On("click", "$counter = 0"),
						g.Text("Reset"),
					),
				),
				Div(Class("demo-box"),
					H3(g.Text("All Signals (JSON)")),
					Pre(ds.JSONSignals(ds.Filter{})),
				),
				Div(Class("demo-box"),
					H3(g.Text("Filtered Signals (counter only)")),
					Pre(ds.JSONSignals(ds.Filter{Include: "/counter/"})),
				),
			),

			// Text Binding Demo
			Div(Class("demo-section"),
				H2(g.Text("Text Binding")),
				P(g.Text("The data-text attribute binds text content to an expression.")),
				Div(Class("demo-box"),
					P(g.Text("Hello, "), Span(ds.Text("$name")), g.Text("!")),
					Input(Type("text"), ds.Bind("name"), Placeholder("Enter your name")),
				),
			),

			// Bind Demo
			Div(Class("demo-section"),
				H2(g.Text("Two-Way Data Binding")),
				P(g.Text("The data-bind attribute creates two-way data binding between signals and form elements.")),
				Div(Class("grid"),
					Div(Class("demo-box"),
						H3(g.Text("Text Input")),
						Input(Type("text"), ds.Bind("inputValue")),
						P(g.Text("Value: "), Span(ds.Text("$inputValue"))),
					),
					Div(Class("demo-box"),
						H3(g.Text("Select")),
						Select(ds.Bind("selected"),
							Option(Value("option1"), g.Text("Option 1")),
							Option(Value("option2"), g.Text("Option 2")),
							Option(Value("option3"), g.Text("Option 3")),
						),
						P(g.Text("Selected: "), Span(ds.Text("$selected"))),
					),
					Div(Class("demo-box"),
						H3(g.Text("Textarea")),
						Textarea(ds.Bind("message"), Placeholder("Type a message...")),
						P(g.Text("Message length: "), Span(ds.Text("$message.length"))),
					),
				),
			),

			// Show Demo
			Div(Class("demo-section"),
				H2(g.Text("Show/Hide Elements")),
				P(g.Text("The data-show attribute shows or hides elements based on an expression.")),
				Div(Class("demo-box"),
					Button(ds.On("click", "$show = !$show"), g.Text("Toggle Visibility")),
					Div(ds.Show("$show"), StyleAttr("margin-top: 10px; padding: 10px; background: #e7f3ff; border-radius: 4px;"),
						g.Text("👋 This element is conditionally visible!"),
					),
					P(g.Text("Visible: "), Span(ds.Text("$show"))),
				),
			),

			// Class Demo
			Div(Class("demo-section"),
				H2(g.Text("Dynamic Classes")),
				P(g.Text("The data-class attribute adds or removes classes based on expressions.")),
				Div(Class("demo-box"),
					Button(ds.On("click", "$isHighlit = !$isHighlit"), g.Text("Toggle Highlight")),
					Div(
						StyleAttr("margin-top: 10px; padding: 10px; border: 1px solid #ddd; border-radius: 4px;"),
						ds.Class("highlight", "$isHighlit", "bold", "$isHighlit"),
						g.Text("This text changes style when highlighted"),
					),
				),
			),

			// Attr Demo
			Div(Class("demo-section"),
				H2(g.Text("Dynamic Attributes")),
				P(g.Text("The data-attr attribute sets attribute values based on expressions.")),
				Div(Class("demo-box"),
					Button(
						ds.On("click", "$counter++"),
						ds.Attr("disabled", "$counter >= 10"),
						g.Text("Click me (max 10)"),
					),
					P(g.Text("Clicks: "), Span(ds.Text("$counter")), g.Text("/10")),
					P(g.Text("Button disabled: "), Span(ds.Text("$counter >= 10"))),
				),
			),

			// Style Demo
			Div(Class("demo-section"),
				H2(g.Text("Dynamic Styles")),
				P(g.Text("The data-style attribute sets inline styles based on expressions.")),
				Div(Class("demo-box"),
					Div(
						ds.Style(
							"color", "$counter % 2 === 0 ? 'blue' : 'red'",
							"fontSize", "$counter > 5 ? '24px' : '16px'",
							"fontWeight", "$counter > 8 ? 'bold' : 'normal'",
						),
						g.Text("This text changes style based on the counter"),
					),
					P(g.Text("Counter: "), Span(ds.Text("$counter"))),
				),
			),

			// Computed Demo
			Div(Class("demo-section"),
				H2(g.Text("Computed Signals")),
				P(g.Text("The data-computed attribute creates read-only signals based on expressions.")),
				Div(Class("demo-box"),
					ds.Computed("total", "$price * $quantity"),
					Div(Class("grid"),
						Div(
							Label(g.Text("Price: $")),
							Input(Type("number"), ds.Bind("price")),
						),
						Div(
							Label(g.Text("Quantity: ")),
							Input(Type("number"), ds.Bind("quantity")),
						),
					),
					P(g.Text("Total: $"), Span(ds.Text("$total"))),
				),
			),

			// Effect Demo
			Div(Class("demo-section"),
				H2(g.Text("Effects")),
				P(g.Text("The data-effect attribute executes expressions when signals change.")),
				Div(Class("demo-box"),
					ds.Effect("console.log('Counter changed to:', $counter)"),
					P(g.Text("Open the browser console and change the counter to see the effect in action.")),
					Button(ds.On("click", "$counter++"), g.Text("Increment (check console)")),
				),
			),

			// Init Demo
			Div(Class("demo-section"),
				H2(g.Text("Initialization")),
				P(g.Text("The data-init attribute runs expressions when elements are loaded.")),
				Div(Class("demo-box"),
					ds.Init("console.log('This section was initialized at:', new Date())"),
					P(g.Text("Check the browser console to see when this section initialized.")),
				),
			),

			// On Events Demo
			Div(Class("demo-section"),
				H2(g.Text("Event Listeners")),
				P(g.Text("The data-on attribute attaches event listeners to elements.")),
				Div(Class("grid"),
					Div(Class("demo-box"),
						H3(g.Text("Click Event")),
						Button(ds.On("click", "$counter++"), g.Text("Click me")),
						P(g.Text("Clicks: "), Span(ds.Text("$counter"))),
					),
					Div(Class("demo-box"),
						H3(g.Text("Mouse Events")),
						Div(
							ds.On("mouseenter", "$show = true"),
							ds.On("mouseleave", "$show = false"),
							StyleAttr("padding: 20px; background: #e7f3ff; border-radius: 4px; cursor: pointer;"),
							g.Text("Hover over me"),
						),
						P(g.Text("Hovering: "), Span(ds.Text("$show"))),
					),
					Div(Class("demo-box"),
						H3(g.Text("Debounced Input")),
						Input(
							Type("text"),
							ds.On("input", "$inputValue = evt.target.value", ds.ModifierDebounce, ds.ModifierDuration(500*time.Millisecond)),
							Placeholder("Type here (debounced 500ms)"),
						),
						P(g.Text("Debounced value: "), Span(ds.Text("$inputValue"))),
					),
				),
			),

			// OnInterval Demo
			Div(Class("demo-section"),
				H2(g.Text("Intervals")),
				P(g.Text("The data-on-interval attribute runs expressions at regular intervals.")),
				Div(Class("demo-box"),
					ds.OnInterval("$count++", ds.ModifierDuration(1*time.Second)),
					P(g.Text("Seconds elapsed: "), Span(ds.Text("$count"))),
				),
			),

			// OnIntersect Demo
			Div(Class("demo-section"),
				H2(g.Text("Intersection Observer")),
				P(g.Text("The data-on-intersect attribute runs expressions when elements enter the viewport.")),
				Div(Class("demo-box"),
					P(g.Text("Scroll down to see the intersection observer in action...")),
					Div(StyleAttr("height: 300px;")),
					Div(
						ds.OnIntersect("$intersected = true", ds.ModifierOnce),
						ds.Class("highlight", "$intersected"),
						StyleAttr("padding: 20px; border: 2px solid #007bff; border-radius: 4px; text-align: center;"),
						g.Text("🎯 I'll highlight when you scroll to me!"),
					),
					P(g.Text("Intersected: "), Span(ds.Text("$intersected"))),
				),
			),

			// Ref Demo
			Div(Class("demo-section"),
				H2(g.Text("Element References")),
				P(g.Text("The data-ref attribute creates signals that reference DOM elements.")),
				Div(Class("demo-box"),
					Div(ds.Ref("myDiv"), StyleAttr("padding: 10px; background: #e7f3ff; border-radius: 4px;"),
						g.Text("I am a referenced element"),
					),
					Button(
						ds.On("click", "$myDiv.style.background = '#ffeb3b'"),
						g.Text("Change Background"),
					),
					Button(
						ds.On("click", "$myDiv.scrollIntoView({ behavior: 'smooth' })"),
						g.Text("Scroll to Element"),
					),
				),
			),

			// PreserveAttr Demo
			Div(Class("demo-section"),
				H2(g.Text("Preserve Attributes")),
				P(g.Text("The data-preserve-attr attribute preserves attribute values during DOM morphing.")),
				Div(Class("demo-box"),
					Details(g.Attr("open"), ds.PreserveAttr("open"),
						Summary(g.Text("Click to expand/collapse")),
						P(g.Text("This content is inside a details element. The 'open' attribute is preserved during DOM updates.")),
					),
				),
			),

			// Ignore Demo
			Div(Class("demo-section"),
				H2(g.Text("Ignore Elements")),
				P(g.Text("The data-ignore attribute tells Datastar to ignore an element and its descendants.")),
				Div(Class("demo-box"),
					Div(ds.Ignore(),
						g.Text("This element is ignored by Datastar. Any data-* attributes here won't work."),
						Button(ds.On("click", "$counter++"), g.Text("This won't work")),
					),
				),
			),

			// IgnoreMorph Demo
			Div(Class("demo-section"),
				H2(g.Text("Ignore Morphing")),
				P(g.Text("The data-ignore-morph attribute prevents elements from being morphed during DOM updates.")),
				Div(Class("demo-box"),
					Div(ds.IgnoreMorph(), StyleAttr("padding: 10px; background: #fff3cd; border-radius: 4px;"),
						g.Text("This element will not be morphed during DOM updates."),
					),
				),
			),

			// OnSignalPatch Demo
			Div(Class("demo-section"),
				H2(g.Text("Signal Patch Events")),
				P(g.Text("The data-on-signal-patch attribute runs expressions when signals are patched.")),
				Div(Class("demo-box"),
					ds.OnSignalPatchFilter(ds.Filter{Include: "/counter/"}),
					ds.OnSignalPatch("console.log('Counter signal patched:', patch)"),
					P(g.Text("Open the console and change the counter to see signal patch events.")),
					Button(ds.On("click", "$counter++"), g.Text("Increment (check console)")),
				),
			),

			// Footer
			Div(Class("demo-section"),
				StyleAttr("text-align: center; background: #f8f9fa;"),
				P(g.Text("Built with "), A(Href("https://www.gomponents.com"), g.Text("gomponents")), g.Text(" and "), A(Href("https://data-star.dev"), g.Text("Datastar"))),
			),
		),
	)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = buildPage().Render(w)
}
