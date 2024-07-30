package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/process"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Basic Kill and Delete Command İmplementation CLI"
	app.Usage = "Let's you kill processes by name or id and delete files or folders"

	app.Commands = []cli.Command{
		{
			Name:        "kill",
			HelpName:    "kill",
			Action:      KillAction,
			ArgsUsage:   ` `,
			Usage:       `kills processes by process id or process name.`,
			Description: `Terminate a process.`,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:  "id",
					Usage: "kill process by process ID.",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "kill process by process name. ",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func KillAction(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments is expected, use flags")
	}

	if c.IsSet("id") && c.IsSet("name") {
		return errors.New("either pid or name flag must be provided")
	}

	if !c.IsSet("id") && c.String("name") == "" {
		return errors.New("name flag cannot be empty")
	}

	if err := killProcess(c); err != nil {
		return err
	}
	fmt.Println("Process killed successfully.")
	return nil
}

func killProcess(c *cli.Context) error {
	if c.IsSet("id") {
		proc, err := process.NewProcess(int32(c.Uint("id")))
		if err != nil {
			return err
		}

		return proc.Kill()
	}

	processes, err := process.Processes()
	if err != nil {
		return err
	}

	var (
		errs  []string
		found bool
	)

	target := c.String("name")
	for _, p := range processes {
		name, _ := p.Name()
		if name == "" {
			continue
		}

		if isEqualProcessName(name, target) {
			found = true
			if err := p.Kill(); err != nil {
				e := err.Error()
				errs = append(errs, e)
			}
		}
	}

	if !found {
		return errors.New("process not found")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

func isEqualProcessName(proc1 string, proc2 string) bool {
	if runtime.GOOS == "linux" {
		return proc1 == proc2
	}
	return strings.EqualFold(proc1, proc2)
}

type Volume struct {
	Name       string
	Total      uint64
	Used       uint64
	Available  uint64
	UsePercent float64
	Mount      string
}

func ActionVolumes(c *cli.Context) error {
	stats, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	var vols []*Volume

	for _, stat := range stats {
		usage, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			continue
		}

		vol := &Volume{
			Name:       stat.Device,
			Total:      usage.Total,
			Used:       usage.Used,
			Available:  usage.Free,
			UsePercent: usage.UsedPercent,
		}

		vols = append(vols, vol)
	}

	volsByteArr, err := json.MarshalIndent(vols, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(volsByteArr))
	return nil
}
