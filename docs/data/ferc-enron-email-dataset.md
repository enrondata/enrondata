# FERC Enron Email Dataset

The FERC Enron Email Data Set may be the second data set users typically find if they look for a more comprehensive data set than the CALO Enron Email Data Set. This is because googling “enron email” will bring up the CMU hosting page for the CALO email data set which refers to the FERC data set.

Using the FERC data set has a few challenges, namely:

1. Large size: The large size of the dataset (100+GB) means that it isn’t readily downloadable. An [online iCONECT interface](http://www.ferc.gov/industries/electric/indus-act/wec/enron/info-release.asp) is available for browsing with attachments. The site is hosted by [Lockheed Martin](http://fercic.bps-lmit.com/members/).
2. iCONECT format: The data comes as static images and in a flat file database format. The latter are “iCONECT24/7 / Concordance databases in delimited record format, with attachments,” not a standard email form such as MIME, PST, or NSF. The format is described in [this WMCU0356_UMD_Transmittal.pdf document](http://zaphod.mindlab.umd.edu:16080/JIKD/Integration/Data/WMCU0356_UMD_Transmittal.pdf).

The dataset is made available in the following formats which are described in the [Aspen Systems document](http://zaphod.mindlab.umd.edu:16080/JIKD/Integration/Data/WMCU0356_UMD_Transmittal.pdf).

1. Enron Email database
2. Enron Email (re-released) database
3. Enron Email (.pst) database
4. Enron Email (.pst) (re-released) database
5. Transcripts
6. Scanned Documents database
7. Scanned Documents (re-released) database

One of the EnronData Project’s goals is to take the FERC email and convert it into properly formatted PST and NSF formats, similar to their original states. A few software vendors have been contacted to see if iCONECT / Concordance databases can be reconstituted into PST / NSF files with attachments without success to date. Without an established solution, the EnronData Project is working on it’s own conversion utilities.
