package main

// vim: ts=4:sts=4

import (
	"encoding/json"
	"errors"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"os"
	"time"
)

type OptionsStruct struct {
	Device string `short:"f" long:"device" value-name:"DEV" description:"Tape device (default: /dev/nst0, or TAPE envvar)"`
	Mock   bool   `long:"mock" description:"Use a mocked tape drive (for tests only)"`
	Debug  bool   `short:"d" long:"debug" description:"Print debug information"`
	Dump   string `long:"dump" value-name:"FILE" description:"Dump SCSI raw data to a file"`

	Man bool `hidden:"1" long:"man"`
}

func main() {
	var err error

	options := &OptionsStruct{}
	parser := flags.NewParser(options, flags.Default)
	_, err = parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	if options.Man {
		fmt.Println("man")
		parser.WriteManPage(os.Stdout)
		return
	}

	//	err = attrSet0803(dev, "Cartouche de Test YaY! 3 spaces:   ")
	//	fmt.Println("Set:",err)

	var drive *TapeDrive

	syncerr := make(chan error)
	go func() {
		if options.Mock {
			drive, err = TapeDriveNewFake()
		} else {
			drive, err = TapeDriveNew(options.Device)
		}
		syncerr <- err
	}()

	openerr := errors.New("timeout")
	started := time.Now()
	waitupto := started.Add(time.Second * 20)
	lastprint := started
waitfor:
	for waitupto.After(time.Now()) && openerr != nil {
		select {
		case openerr = <-syncerr:
			if openerr != nil {
				fmt.Println("Failed")
				os.Exit(1)
			} else {
				// openerr
				break waitfor
			}
		default:
			if time.Since(lastprint) > time.Second {
				fmt.Printf("Still trying to open the device, aborting in %s...\n", time.Until(waitupto).Round(time.Second))
				lastprint = time.Now()
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	if openerr != nil {
		fmt.Println("Timed out opening device, is a tape inserted?")
		os.Exit(1)
	} else {
		fmt.Printf("Device %s opened\n", drive.DeviceName)
	}

	if options.Dump != "" {
		drive.SetDumpFile(options.Dump)
	}

	/*
			poh := &LogSenseType{0x3C, 0x0008}
			err = scsiLogSense(dev, poh)
			if err != nil {
				fmt.Println("logtest:",err)
		/	}
	*/

	//fmt.Println("Inquiry")
	/*err =*/
	drive.ScsiInquiry()
	//fmt.Println("Inquiry err:")
	//fmt.Println(err)

	drive.GetStatus()

	for i, a := range attributes {
		fmt.Printf("\rReading attribute %02d/%02d...", i+1, len(attributes))
		drive.GetAttribute(a)
		if options.Debug {
			j, _ := json.MarshalIndent(a, "", "   ")
			fmt.Println(string(j))
		}
	}
	fmt.Println("")
	//fmt.Printf("\r                              \r")

	//drive.CmList.Print()
	fmt.Println(drive)
}
