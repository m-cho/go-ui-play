package main

func main() {
	App(
		VerticalBox(
			Label("This is a label. Right now, labels can only span one line."),
			Button("This is button"),
		),
	).Window("libui GoLang", 640, 480, true).Start()
}
