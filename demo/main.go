package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	data "maragu.dev/gomponents-datastar"
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

func buildPage() Node {
	return HTML(
		Lang("en"),
		Head(
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			TitleEl(Text("Datastar Attributes Demo")),
			Script(Type("module"), Src("https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js")),
			StyleEl(Type("text/css"), Raw(`
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
			H1(Text("Datastar Attributes Demo")),
			P(Text("This page demonstrates all the Datastar attributes available in gomponents-datastar.")),

			// Signals Demo
			Div(Class("demo-section"),
				data.Signals(map[string]any{
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
				H2(Text("Signals")),
				P(Text("Signals are reactive state variables. Click the buttons to modify them.")),
				Div(Class("demo-box"),
					H3(Text("Counter Example")),
					P(data.Text("$counter"), Text(" clicks")),
					Button(
						data.On("click", "$counter++"),
						Text("Increment"),
					),
					Button(
						data.On("click", "$counter--"),
						Text("Decrement"),
					),
					Button(
						data.On("click", "$counter = 0"),
						Text("Reset"),
					),
				),
				Div(Class("demo-box"),
					H3(Text("All Signals (JSON)")),
					Pre(data.JSONSignals(data.Filter{})),
				),
				Div(Class("demo-box"),
					H3(Text("Filtered Signals (counter only)")),
					Pre(data.JSONSignals(data.Filter{Include: "/counter/"})),
				),
			),

			// Text Binding Demo
			Div(Class("demo-section"),
				H2(Text("Text Binding")),
				P(Text("The data-text attribute binds text content to an expression.")),
				Div(Class("demo-box"),
					P(Text("Hello, "), Span(data.Text("$name")), Text("!")),
					Input(Type("text"), data.Bind("name"), Placeholder("Enter your name")),
				),
			),

			// Bind Demo
			Div(Class("demo-section"),
				H2(Text("Two-Way Data Binding")),
				P(Text("The data-bind attribute creates two-way data binding between signals and form elements.")),
				Div(Class("grid"),
					Div(Class("demo-box"),
						H3(Text("Text Input")),
						Input(Type("text"), data.Bind("inputValue")),
						P(Text("Value: "), Span(data.Text("$inputValue"))),
					),
					Div(Class("demo-box"),
						H3(Text("Select")),
						Select(data.Bind("selected"),
							Option(Value("option1"), Text("Option 1")),
							Option(Value("option2"), Text("Option 2")),
							Option(Value("option3"), Text("Option 3")),
						),
						P(Text("Selected: "), Span(data.Text("$selected"))),
					),
					Div(Class("demo-box"),
						H3(Text("Textarea")),
						Textarea(data.Bind("message"), Placeholder("Type a message...")),
						P(Text("Message length: "), Span(data.Text("$message.length"))),
					),
				),
			),

			// Show Demo
			Div(Class("demo-section"),
				H2(Text("Show/Hide Elements")),
				P(Text("The data-show attribute shows or hides elements based on an expression.")),
				Div(Class("demo-box"),
					Button(data.On("click", "$show = !$show"), Text("Toggle Visibility")),
					Div(data.Show("$show"), StyleAttr("margin-top: 10px; padding: 10px; background: #e7f3ff; border-radius: 4px;"),
						Text("👋 This element is conditionally visible!"),
					),
					P(Text("Visible: "), Span(data.Text("$show"))),
				),
			),

			// Class Demo
			Div(Class("demo-section"),
				H2(Text("Dynamic Classes")),
				P(Text("The data-class attribute adds or removes classes based on expressions.")),
				Div(Class("demo-box"),
					Button(data.On("click", "$isHighlit = !$isHighlit"), Text("Toggle Highlight")),
					Div(
						StyleAttr("margin-top: 10px; padding: 10px; border: 1px solid #ddd; border-radius: 4px;"),
						data.Class("highlight", "$isHighlit", "bold", "$isHighlit"),
						Text("This text changes style when highlighted"),
					),
				),
			),

			// Attr Demo
			Div(Class("demo-section"),
				H2(Text("Dynamic Attributes")),
				P(Text("The data-attr attribute sets attribute values based on expressions.")),
				Div(Class("demo-box"),
					Button(
						data.On("click", "$counter++"),
						data.Attr("disabled", "$counter >= 10"),
						Text("Click me (max 10)"),
					),
					P(Text("Clicks: "), Span(data.Text("$counter")), Text("/10")),
					P(Text("Button disabled: "), Span(data.Text("$counter >= 10"))),
				),
			),

			// Style Demo
			Div(Class("demo-section"),
				H2(Text("Dynamic Styles")),
				P(Text("The data-style attribute sets inline styles based on expressions.")),
				Div(Class("demo-box"),
					Div(
						data.Style(
							"color", "$counter % 2 === 0 ? 'blue' : 'red'",
							"fontSize", "$counter > 5 ? '24px' : '16px'",
							"fontWeight", "$counter > 8 ? 'bold' : 'normal'",
						),
						Text("This text changes style based on the counter"),
					),
					P(Text("Counter: "), Span(data.Text("$counter"))),
				),
			),

			// Computed Demo
			Div(Class("demo-section"),
				H2(Text("Computed Signals")),
				P(Text("The data-computed attribute creates read-only signals based on expressions.")),
				Div(Class("demo-box"),
					data.Computed("total", "$price * $quantity"),
					Div(Class("grid"),
						Div(
							Label(Text("Price: $")),
							Input(Type("number"), data.Bind("price")),
						),
						Div(
							Label(Text("Quantity: ")),
							Input(Type("number"), data.Bind("quantity")),
						),
					),
					P(Text("Total: $"), Span(data.Text("$total"))),
				),
			),

			// Effect Demo
			Div(Class("demo-section"),
				H2(Text("Effects")),
				P(Text("The data-effect attribute executes expressions when signals change.")),
				Div(Class("demo-box"),
					data.Effect("console.log('Counter changed to:', $counter)"),
					P(Text("Open the browser console and change the counter to see the effect in action.")),
					Button(data.On("click", "$counter++"), Text("Increment (check console)")),
				),
			),

			// Init Demo
			Div(Class("demo-section"),
				H2(Text("Initialization")),
				P(Text("The data-init attribute runs expressions when elements are loaded.")),
				Div(Class("demo-box"),
					data.Init("console.log('This section was initialized at:', new Date())"),
					P(Text("Check the browser console to see when this section initialized.")),
				),
			),

			// On Events Demo
			Div(Class("demo-section"),
				H2(Text("Event Listeners")),
				P(Text("The data-on attribute attaches event listeners to elements.")),
				Div(Class("grid"),
					Div(Class("demo-box"),
						H3(Text("Click Event")),
						Button(data.On("click", "$counter++"), Text("Click me")),
						P(Text("Clicks: "), Span(data.Text("$counter"))),
					),
					Div(Class("demo-box"),
						H3(Text("Mouse Events")),
						Div(
							data.On("mouseenter", "$show = true"),
							data.On("mouseleave", "$show = false"),
							StyleAttr("padding: 20px; background: #e7f3ff; border-radius: 4px; cursor: pointer;"),
							Text("Hover over me"),
						),
						P(Text("Hovering: "), Span(data.Text("$show"))),
					),
					Div(Class("demo-box"),
						H3(Text("Debounced Input")),
						Input(
							Type("text"),
							data.On("input", "$inputValue = evt.target.value", data.ModifierDebounce, data.Duration(500*time.Millisecond)),
							Placeholder("Type here (debounced 500ms)"),
						),
						P(Text("Debounced value: "), Span(data.Text("$inputValue"))),
					),
				),
			),

			// OnInterval Demo
			Div(Class("demo-section"),
				H2(Text("Intervals")),
				P(Text("The data-on-interval attribute runs expressions at regular intervals.")),
				Div(Class("demo-box"),
					data.OnInterval("$count++", data.ModifierDuration, data.Duration(1*time.Second)),
					P(Text("Seconds elapsed: "), Span(data.Text("$count"))),
				),
			),

			// OnIntersect Demo
			Div(Class("demo-section"),
				H2(Text("Intersection Observer")),
				P(Text("The data-on-intersect attribute runs expressions when elements enter the viewport.")),
				Div(Class("demo-box"),
					P(Text("Scroll down to see the intersection observer in action...")),
					Div(StyleAttr("height: 300px;")),
					Div(
						data.OnIntersect("$intersected = true", data.ModifierOnce),
						data.Class("highlight", "$intersected"),
						StyleAttr("padding: 20px; border: 2px solid #007bff; border-radius: 4px; text-align: center;"),
						Text("🎯 I'll highlight when you scroll to me!"),
					),
					P(Text("Intersected: "), Span(data.Text("$intersected"))),
				),
			),

			// Ref Demo
			Div(Class("demo-section"),
				H2(Text("Element References")),
				P(Text("The data-ref attribute creates signals that reference DOM elements.")),
				Div(Class("demo-box"),
					Div(data.Ref("myDiv"), StyleAttr("padding: 10px; background: #e7f3ff; border-radius: 4px;"),
						Text("I am a referenced element"),
					),
					Button(
						data.On("click", "$myDiv.style.background = '#ffeb3b'"),
						Text("Change Background"),
					),
					Button(
						data.On("click", "$myDiv.scrollIntoView({ behavior: 'smooth' })"),
						Text("Scroll to Element"),
					),
				),
			),

			// PreserveAttr Demo
			Div(Class("demo-section"),
				H2(Text("Preserve Attributes")),
				P(Text("The data-preserve-attr attribute preserves attribute values during DOM morphing.")),
				Div(Class("demo-box"),
					Details(Attr("open"), data.PreserveAttr("open"),
						Summary(Text("Click to expand/collapse")),
						P(Text("This content is inside a details element. The 'open' attribute is preserved during DOM updates.")),
					),
				),
			),

			// Ignore Demo
			Div(Class("demo-section"),
				H2(Text("Ignore Elements")),
				P(Text("The data-ignore attribute tells Datastar to ignore an element and its descendants.")),
				Div(Class("demo-box"),
					Div(data.Ignore(),
						Text("This element is ignored by Datastar. Any data-* attributes here won't work."),
						Button(data.On("click", "$counter++"), Text("This won't work")),
					),
				),
			),

			// IgnoreMorph Demo
			Div(Class("demo-section"),
				H2(Text("Ignore Morphing")),
				P(Text("The data-ignore-morph attribute prevents elements from being morphed during DOM updates.")),
				Div(Class("demo-box"),
					Div(data.IgnoreMorph(), StyleAttr("padding: 10px; background: #fff3cd; border-radius: 4px;"),
						Text("This element will not be morphed during DOM updates."),
					),
				),
			),

			// OnSignalPatch Demo
			Div(Class("demo-section"),
				H2(Text("Signal Patch Events")),
				P(Text("The data-on-signal-patch attribute runs expressions when signals are patched.")),
				Div(Class("demo-box"),
					data.OnSignalPatchFilter(data.Filter{Include: "/counter/"}),
					data.OnSignalPatch("console.log('Counter signal patched:', patch)"),
					P(Text("Open the console and change the counter to see signal patch events.")),
					Button(data.On("click", "$counter++"), Text("Increment (check console)")),
				),
			),

			// Footer
			Div(Class("demo-section"),
				StyleAttr("text-align: center; background: #f8f9fa;"),
				P(Text("Built with "), A(Href("https://www.gomponents.com"), Text("gomponents")), Text(" and "), A(Href("https://data-star.dev"), Text("Datastar"))),
			),
		),
	)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = buildPage().Render(w)
}
