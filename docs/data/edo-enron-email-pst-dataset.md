# EDO Enron Email PST Dataset

Although much of the original Enron Email came in PST files, the most common form to get this email in today is in [MIME format](https://www.cs.cmu.edu/~enron/) from the CMU [CALO Project](https://archived.ai.sri.com/project/CALO.html).

Previously, the CMU / CALO dataset was converted to PST format by Pete Warden [earlier PST conversion](https://petewarden.com/2008/03/25/how-to-conver-1/). Pete’s PST is similar to journal email in that per-user delineation and folder structure of the user email stores have been removed.

To preserve the user information associated with the email, EnronData.org has converted the CALO Enron Email Dataset to the form of 148 custodian PST files with folder structure, preserving the information in the CALO dataset. Email for each of the [148 identified custodians](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.html) is available in per-custodian PST files. A few minor changes were made to correct names and merge duplicate users where both correct and incorrect names existed. Custom X-headers have also been added including unique IDs to facilitate testing.

The files are currently available for download from the EnronData.org homepage as a 734MB 7z archive. 7z is an archival format similar to ZIP, BZIP2 and RAR that generally achieves higher compression rates. The uncompressed size for this dataset is roughly 8.6 gigabytes. For verification, please use the [MD5 checksums](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_pst_dataset_md5sum.log).

This dataset is licensed under the [Creative Commons Attribution 3.0 United States license](http://creativecommons.org/licenses/by/3.0/us/). To provide attribution, please cite to “EnronData.org.”
