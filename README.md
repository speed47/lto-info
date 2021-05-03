LTO Info Tool
=============

This tool reads the internal memory of LTO/Ultrium cartridges from a tape drive 🖭

More precisely, this tool can:
- Read, decode and display factory and usage information stored in the Cartridge Memory (`CM`) aka Medium Auxiliary Memory (`MAM`)
- View and modify the custom `User Medium Text Label` of a cartridge
- Display basic information about the tape drive device

## :hammer: How to build and use

You need to have a go build environment [properly set up](https://golang.org/doc/install), then just type:

```
make
```

And try the tool by typing:

```
./lto-info
```

By default, the tool will look for a tape device in `/dev/nst0`, or what is pointed to by the `TAPE` environment variable. To specify another device, use the `-d` option.

## :mag: Output example
```
Drive information:
   Vendor  : HP
   Model   : Ultrium 2-SCSI
   Firmware: F63D
Medium information:
  Cartridge Type: 0x00 - Data cartridge
  Medium format : 0x42 - LTO-2
  Formatted as  : 0x42 - LTO-2
  Assign. Org.  : LTO-CVE
  Manufacturer  : IMATION
  Serial No     : 0E00776390
  Manuf. Date   : 2011-12-07 (roughly 9.4 years ago)
  Tape length   : 609 meters
  Tape width    : 12.7 mm
  MAM Capacity  : 4096 bytes (1014 bytes remaining)
Format specs:
   Capacity  :   200 GB native   -   400 GB compressed with a 2:1 ratio
   R/W Speed :    40 MB/s native -    80 MB/s compressed
   Partitions:     1 max partitions supported
   Phy. specs: 4 bands/tape, 16 wraps/band, 8 tracks/wrap, 512 total tracks
   Duration  : 1h23 to fill tape with 64 end-to-end passes (78 seconds/pass)
Usage information:
  Partition space free  : 96% (193746/200448 MiB, 189/195 GiB, 0.18/0.19 TiB)
  Cartridge load count  : 13
  Data written - alltime:        81983 MiB (    80.06 GiB,   0.08 TiB, 0.41 FVE)
  Data read    - alltime:        20674 MiB (    20.19 GiB,   0.02 TiB, 0.10 FVE)
  Data written - session:            0 MiB (     0.00 GiB,   0.00 TiB, 0.00 FVE)
  Data read    - session:          139 MiB (     0.14 GiB,   0.00 TiB, 0.00 FVE)
Previous sessions:
  Session N-0: Used in a device of vendor HP (serial HU10625W0T)
  Session N-1: Used in a device of vendor HP (serial HU10625W0T)
  Session N-2: Used in a device of vendor HP (serial HUP9B067QF)
  Session N-3: Used in a device of vendor HP (serial HUP9B067QF)
```

## :gem: Build-time dependencies

- https://github.com/HewlettPackard/structex: encode/decode bitfields in SCSI structs in a readable way
- https://github.com/benmcclelland/mtio: go bindings for `mt` ioctls
- https://github.com/benmcclelland/sgio: go bindings for `sgio` ioctls
- https://github.com/jessevdk/go-flags: command-line options parser

## :hearts: Thanks

Inspired from other tools written in C:
- https://github.com/arogge/maminfo
- https://github.com/scangeo/lto-cm/

Big up to them!
