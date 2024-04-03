# DownCalc

Downcalc is a little cli tool that allows you to calculate the estimated download time for a file of a given filesize

## Usage
```bash
downcal 10gb

# Set the download speed
downcal config 100 #sets the download speed to 100 MBit/s. not giving a value prompts a speed test
=> Download time with 100 MBit/s for 100gb:  2h:23m:9s

# Set the download speed ones
downcal -s 500 20gb # calculates the time needed to downlaod 20GB with 500 MBit/s
```

## Configure

To set the download speed use the `config` command

```bash
downcal config #100 (optional download speed in MBit/s)
```

you can either let the tool run a speed test or input the speed yourself


