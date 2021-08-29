package main

func main() {

	g := createGame()

	if err := run(g); err != nil {
		g.Over(err)
	}

}
