# goxpdf

## Simply bulk generator certificate, using a base template and a csv file with names within it, written in golang!

### How to use it

1. Update the template part - `templatePath`
2. Update the csv path - `csvPath`
3. Need custom font? Update - `fontPath`
4. For font size & color - `fontSize` & `textColor`
5. To arrange text on the Y - axis and to allign it with the template, play along with the `textY`
6. X - axis is auto centered, to mantain consistency across the ceritificates

### Benchmark

| Quantity (PNGs) | Time Taken (in seconds) | Average time per certificate (ms) |
| ----------- | ----------- | ----------- |
| 100 | 23.96s | 239.61ms |
| 250 | 61.28s | 245.13ms |
| 1000 | 229.25s | 229.24ms |
