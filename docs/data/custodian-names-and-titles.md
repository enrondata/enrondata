# Custodian Names and Titles

In order to reconstruct the Enron email dataset accurately it is important to identify the correct number of custodians for which email exists. From this canonical list, we can build out user information including actual names, rank, title, etc.

Various datasets have used a string consisting generally of lastname and first initial to identify custodians. This ID appears in the FERC dataset as the ORIGIN field and in the CALO dataset as a directory name.

Using this ID, different versions of the Enron email dataset use different numbers of users. Notably, the following numbers of users are used:

1. 158 users: The FERC dataset identifies 158 unique users using the iCONECT ORIGIN database column.
2. 150 users: The CALO dataset identifies 150 unique users using the maildir user directory structure.
3. 149 users: Andrés Corrada-Emmanuel has identified 149 unique users noting that phanis-s is a misspelling of panus-s.
4. 148 users: EnronData.org has identified 148 unique users noting that whalley-l is a duplicate of whalley-g, both representing Lawrence “Greg” Whalley.

EnronData.org has verified the duplicates identified by CALO, identified two more duplicates, and corrected two misspellings as shown in [this Enron custodian list](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians.txt). This list was created before analyzing Corrada-Emmanuel’s custodian list which correctly identified one of additional duplicates.

After identifying custodians by ID, additional information can be associated with the custodians including names, ranks, titles, email addresses etc. A large part of the work has been performed by Jitesh Shetty and Jafar Adibi in their [Ex-Employee Status Report](https://web.archive.org/web/20131213191703/https://www.isi.edu/~adibi/Enron/Enron.htm). Combining this data with some additional data allows us to associate this information directly with the custodians in the dataset. The final results are available in the [this custodian information report](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.html). Corrada-Emmanuel has also created a [custodian ID to email address mapping](http://ciir.cs.umass.edu/~corrada/enron/folder-normalized-author.txt.gz). Email addresses will be incorporated by EnronData.org at a future date.
