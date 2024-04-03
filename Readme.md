# DownCalc

Downcalc is a little cli tool that allows you to calculate the estimated download time for a file of a given filesize.
> Notice file sizes are treated as the n^2 sizes so 1GB == 10243 Bytes

## Usage
```bash
downcal 100gb
# => Download time with 100 MBit/s for 100gb:  2h:23m:9s
```

### Set the download speed
```bash
downcal config 100 #sets the download speed to 100 MBit/s. not giving a value prompts a speed test
```

### Set the download speed once
```bash
downcal -s 500 20gb # calculates the time needed to downlaod 20GB with 500 MBit/s
```

## Configure

To set the download speed use the `config` command

```bash
downcal config 100 #(optional download speed in MBit/s. If not set the tool will prompt you to run a speed test)
```

you can either let the tool run a speed test or input the speed yourself



