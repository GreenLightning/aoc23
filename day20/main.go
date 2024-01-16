package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	KindNone int = iota
	KindFlipFlop
	KindConjunction
	KindBroadcaster
	KindButton
)

type Module struct {
	Name                string
	Kind                int
	Inputs              []*Module
	Outputs             []*Module
	OutputInputIndexMap []int
	FlipFlopState       bool
	ConjunctionState    []bool
}

func (m *Module) Description() string {
	switch m.Kind {
	case KindFlipFlop:
		return "%" + m.Name
	case KindConjunction:
		return "&" + m.Name
	default:
		return m.Name
	}
}

func main() {
	lines := readLines("input.txt")

	var modules []*Module
	var broadcaster *Module

	regex := regexp.MustCompile(`^(?:(broadcaster)|([&%]?)(\w+)) -> (\w+(?:, \w+)*)$`)
	modulesByName := make(map[string]*Module)
	outputsByName := make(map[string]string)
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		module := new(Module)
		if matches[1] == "broadcaster" {
			module.Name = matches[1]
			module.Kind = KindBroadcaster
			broadcaster = module
		} else {
			module.Name = matches[3]
			if matches[2] == "%" {
				module.Kind = KindFlipFlop
			} else {
				module.Kind = KindConjunction
			}
		}
		modules = append(modules, module)
		modulesByName[module.Name] = module
		outputsByName[module.Name] = matches[4]
	}

	for _, module := range modules {
		outputs := outputsByName[module.Name]
		for _, outputName := range strings.Split(outputs, ", ") {
			output := modulesByName[outputName]
			if output == nil {
				output = new(Module)
				output.Name = outputName
				modulesByName[output.Name] = output
			}
			module.Outputs = append(module.Outputs, output)
			module.OutputInputIndexMap = append(module.OutputInputIndexMap, len(output.Inputs))
			output.Inputs = append(output.Inputs, module)
		}
	}

	visualizeInput(modules)

	{
		fmt.Println("--- Part One ---")
		var highCount, lowCount int
		simulate(modules, broadcaster,
			func(presses int, pulse Pulse) {
				if pulse.High {
					highCount++
				} else {
					lowCount++
				}
			},
			func(presses int) bool {
				if presses == 1000 {
					fmt.Println(lowCount * highCount)
					return true
				}
				return false
			},
		)
	}

	{
		fmt.Println("--- Part Two ---")
		rx := modulesByName["rx"]

		if len(rx.Inputs) != 1 {
			panic("invalid input")
		}

		collector := rx.Inputs[0]

		if collector.Kind != KindConjunction {
			panic("invalid input")
		}

		done := make([]bool, len(collector.Inputs))
		remaining := len(collector.Inputs)
		period := 1

		simulate(modules, broadcaster,
			func(presses int, pulse Pulse) {
				if pulse.High && pulse.Target == collector {
					if !done[pulse.Input] {
						done[pulse.Input] = true
						remaining--
						period = lcm(period, presses)
					}
				}
			},
			func(presses int) bool {
				if remaining == 0 {
					fmt.Println(period)
					return true
				}
				return false
			},
		)
	}
}

func visualizeInput(modules []*Module) {
	var buffer bytes.Buffer
	buffer.WriteString("digraph {\n")
	for _, m := range modules {
		for _, output := range m.Outputs {
			fmt.Fprintf(&buffer, "\"%s\" -> \"%s\"\n", m.Description(), output.Description())
		}
	}
	buffer.WriteString("}\n")

	check(os.WriteFile("input.gv", buffer.Bytes(), 0644))
	cmd := exec.Command("dot", "-o", "input.pdf", "-Tpdf", "input.gv")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	check(cmd.Run())
}

type Pulse struct {
	Source *Module
	Target *Module
	Input  int
	High   bool
}

func (pulse Pulse) String() string {
	s := "low"
	if pulse.High {
		s = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", pulse.Source.Name, s, pulse.Target.Name)
}

func simulate(modules []*Module, broadcaster *Module, onPulse func(presses int, pulse Pulse), afterPress func(presses int) bool) {
	// Fake module to populate pulse source for printing.
	button := new(Module)
	button.Name = "button"
	button.Kind = KindButton

	// Reset modules.
	for _, module := range modules {
		module.FlipFlopState = false
		if module.Kind == KindConjunction {
			module.ConjunctionState = make([]bool, len(module.Inputs))
		}
	}

	var queue []Pulse
	for presses := 1; ; presses++ {
		queue = append(queue, Pulse{button, broadcaster, 0, false})
		for len(queue) != 0 {
			pulse := queue[0]
			queue = queue[1:]

			onPulse(presses, pulse)

			// fmt.Println(pulse.String())

			module := pulse.Target

			send := func(high bool) {
				for i, output := range module.Outputs {
					queue = append(queue, Pulse{module, output, module.OutputInputIndexMap[i], high})
				}
			}

			switch module.Kind {
			case KindFlipFlop:
				if !pulse.High {
					module.FlipFlopState = !module.FlipFlopState
					send(module.FlipFlopState)
				}

			case KindConjunction:
				module.ConjunctionState[pulse.Input] = pulse.High
				all := true
				for _, state := range module.ConjunctionState {
					if !state {
						all = false
						break
					}
				}
				send(!all)

			case KindBroadcaster:
				send(pulse.High)
			}
		}

		if done := afterPress(presses); done {
			break
		}
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
