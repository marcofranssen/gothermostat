package main

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := &config{}

	err := cfg.load(configFile)
	check(err)

	nest := &nest{config: cfg}
	err = nest.authenticate()
	check(err)
	cfg.save(configFile)
}
