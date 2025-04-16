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

## :mag: Partial output example
```
Drive information:
   Vendor  : TANDBERG
   Model   : LTO-6 HH
   Firmware: 3579
Medium information:
  Cartridge Type: 0x00 - Data cartridge
  Medium format : 0x5a - LTO-6
  Formatted as  : 0x5a - LTO-6
  Assign. Org.  : LTO-CVE
  Manufacturer  : HPE
  Serial No     : F26VYYRDEP
  Barcode No    : FJK676L6
  Manuf. Date   : 2022-07-25 (roughly 2.7 years ago)
  Tape length   : 846 meters
  Tape width    : 12.7 mm
  MAM Capacity  : 16384 bytes
Format specs:
   Capacity  :  2500 GB native   -  6250 GB compressed with a 2.5:1 ratio
   R/W Speed :   160 MB/s native -   400 MB/s compressed
   Partitions:     4 max partitions supported
   Phy. specs: 4 bands/tape, 34 wraps/band, 16 tracks/wrap, 2176 total tracks
   Duration  : 4h20 to fill tape with 136 end-to-end passes (115 seconds/pass)
Usage information:
  Data written - alltime:            0 MiB (     0.00 GiB,   0.00 TiB)
  Data read    - alltime:            0 MiB (     0.00 GiB,   0.00 TiB)
  Data written - session:            0 MiB (     0.00 GiB,   0.00 TiB)
  Data read    - session:            0 MiB (     0.00 GiB,   0.00 TiB)
Previous sessions:
  Session N-0: Used in a device of vendor HP (serial XXXXXXXXXX)
  Session N-1: Used in a device of vendor HP
  Session N-2: Used in a device of vendor HP
  Session N-3: Used in a device of vendor HP
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
