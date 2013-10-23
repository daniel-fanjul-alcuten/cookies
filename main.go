package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

func read(filename string) *Cookieverse {
	var c Cookieverse
	if file, err := os.Open(filename); err == nil {
		reader := bufio.NewReader(file)
		decoder := json.NewDecoder(reader)
		if err := decoder.Decode(&c); err != nil {
			c = Cookieverse{}
		}
	}
	return &c
}

func write(filename string, c *Cookieverse) error {
	if bytes, err := json.MarshalIndent(c, "", "  "); err == nil {
		return ioutil.WriteFile(filename, bytes, 0666)
	} else {
		return err
	}
}

func work() error {

	filename := flag.String("f", "cookies.json", "the file")
	depth := flag.Int("d", 4, "the depth")
	number := flag.Int("n", 2, "the number")
	limit := flag.Int("l", 1e4, "the limit")
	flag.Parse()

	input := make(chan string)
	output := make(chan *Node, 1)
	go Search(read(*filename), input, output, *depth, *number, *limit)

	errs := make(chan error, 2)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			errs <- err
		}
		close(input)
	}()

	wait := make(chan struct{})
	go func() {
		cpsString := func(prefix string, c *Cookieverse, t Time) string {
			// fmt.Printf("%s%f cursor cps\n", prefix, c.CursorCpS())
			// fmt.Printf("%s%f grandma cps\n", prefix, c.GrandmaCpS())
			// fmt.Printf("%s%f farm cps\n", prefix, c.FarmCpS())
			// fmt.Printf("%s%f factory cps\n", prefix, c.FactoryCpS())
			// fmt.Printf("%s%f mine cps\n", prefix, c.MineCpS())
			// fmt.Printf("%s%f shipment cps\n", prefix, c.ShipmentCpS())
			// fmt.Printf("%s%f alchemy lab cps\n", prefix, c.AlchemyLabCpS())
			// fmt.Printf("%s%f portal cps\n", prefix, c.PortalCpS())
			// fmt.Printf("%s%f time machine cps\n", prefix, c.TimeMachineCpS())
			// fmt.Printf("%s%f antimatter condenser cps\n", prefix, c.AntimatterCondenserCpS())
			// fmt.Printf("%s%f normal multiplier \n", prefix, c.NormalMultiplier())
			// fmt.Printf("%s%f heavenly multiplier \n", prefix, c.HeavenlyMultiplier())
			// fmt.Printf("%s%f multiplier cps \n", prefix, c.MultiplierCpS())
			cps := c.Production()
			multiplier := int(cps * 100 / c.BuildingCpS())
			lucky := 12e2 * cps
			return fmt.Sprintf("%s%f cps (%d%%) [%.2g/%.2g] in %v\n", prefix,
				cps, multiplier, lucky, 7*lucky, t)
		}
		transitionString := func(i int, node *Node, t *Transition) string {
			cps1, cps2 := node.Production, t.Production
			lucky := Stock(7 * 12e3 * cps2)
			return fmt.Sprintf("%02d: %s(%d): -%.2g/%.3g +%.2g in %v\n", i+1,
				t.Description, len(node.Best), t.Cost, lucky+t.Cost, cps2-cps1, t.Time)
		}
		nodeString := func(node *Node) string {
			s := cpsString("< ", node.State.(*Cookieverse), 0)
			i, next, time := 0, node, Time(0)
			for len(next.Best) > 0 {
				t := next.Best[0]
				s += transitionString(i, next, t)
				time += t.Time
				i, next = i+1, t.Node
			}
			s += cpsString("> ", next.State.(*Cookieverse), time)
			return s
		}
		node := <-output
		lastNodeString := ""
		var lastMemStats runtime.MemStats
		for {
			select {
			case n, ok := <-output:
				if !ok {
					close(wait)
					return
				}
				if n != node {
					if err := write(*filename, n.State.(*Cookieverse)); err != nil {
						errs <- err
						close(wait)
						return
					}
					node = n
				}
				s := nodeString(node)
				if lastNodeString != s {
					fmt.Print(s)
					lastNodeString = s
				}
			case <-time.After(5 * time.Second):
				var mem runtime.MemStats
				runtime.ReadMemStats(&mem)
				fmt.Println("<< MemStats")
				fmt.Printf("HeapAlloc: %.3g (%.2g)\n", float64(mem.HeapAlloc),
					float64(mem.HeapAlloc)-float64(lastMemStats.HeapAlloc))
				fmt.Printf("HeapSys: %.3g (%.2g)\n", float64(mem.HeapSys),
					float64(mem.HeapSys)-float64(lastMemStats.HeapSys))
				fmt.Printf("HeapIdle: %.3g (%.2g)\n", float64(mem.HeapIdle),
					float64(mem.HeapIdle)-float64(lastMemStats.HeapIdle))
				fmt.Printf("HeapInuse: %.3g (%.2g)\n", float64(mem.HeapInuse),
					float64(mem.HeapInuse)-float64(lastMemStats.HeapInuse))
				fmt.Printf("HeapReleased: %.3g (%.2g)\n", float64(mem.HeapReleased),
					float64(mem.HeapReleased)-float64(lastMemStats.HeapReleased))
				fmt.Printf("HeapObjects: %.3g (%.2g)\n", float64(mem.HeapObjects),
					float64(mem.HeapObjects)-float64(lastMemStats.HeapObjects))
				fmt.Print(lastNodeString)
				lastMemStats = mem
			}
		}
	}()

	select {
	case err := <-errs:
		return err
	case <-wait:
		return nil
	}
}

func main() {
	if err := work(); err != nil {
		log.Fatal(err)
	}
}
