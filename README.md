# graper

Toolkit for calculating winemaking procedures. 
Built using Vue and Golang.

## Monitor Ferment from Refractometer Readings
A set of calculations are performed, first to calculate the initial gravity from the inital brix:

ig = 1.000898 + 0.003859118 * ib + 0.00001370735 * ib2 + 0.00000003742517 * ib3

And then to calculate the current gravity from the initial and current brix:

cg = 1.001843 - 0.002318474 *ib - 0.000007775 *ib2 - 0.000000034 * ib3 + 0.00574 *cb + 0.00003344 * cb2 + 0.000000086 *cb3

The alcohol by volume is calculated thus:

abv = 0.93 * ( ( 1017.5596 - ( 277.4 * cg ) + ( 1.33302 + 0.001427193 * cb + 0.000005791157 * cb2 ) * ( ( 937.8135 * ( 1.33302 + 0.001427193 * cb + 0.000005791157 * cb2 ) ) - 1805.1228 ) ) * ( cg / 0.794 ) )

Where:
abv = alcohol by volume
cg = current specific gravity
cb = current Brix reading (refractometer)
NOTE The 0.93 conversion factor was added based on experimental results to make the alcohol prediction for this particular calculator more accurate.

The residual sugar (in grams per litre) is calculated thus:

residual sugar = specific gravity * true brix

Spirit indication is calculated form the current alcohol thus, and used to adjust the current gravity, so that true Brix and residual sugar can be established:

si = (2 * SQRT ( 626159497 ) * SQRT ( 35209254016727200 * abv + 448667639342033000 ) - 33520822512398 ) / 841662180975

This is then used to calculate the corrected SG:

corrected_sg = current_sg - ( 1 - ( spirit_indication / 1000) ) + 1

Note that this alcohol calculation is different than the one used in the alcohol by refractometer and hydrometer as it gives more reliable results when coupled with the other calculations used to monitor fermentation using a refractometer.