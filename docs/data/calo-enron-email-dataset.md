# CALO Enron Email Dataset

The [Carnegie Mellon University (CMU)](http://www.cmu.edu/) [CALO Project](https://archived.ai.sri.com/project/CALO.html) dataset is perhaps the most widely used dataset and is available for download at [http://www.cs.cmu.edu/~enron/](http://www.cs.cmu.edu/~enron/). This dataset is a derivative of the FERC dataset and has been referenced in many email research studies and is also used by many commercial E-Discovery organizations. The CMU page describes this dataset as follows:

1. This dataset was collected and prepared by the CALO Project (A Cognitive Assistant that Learns and Organizes).
2. It contains data from [150 custodians](https://github.com/enrondata/enrondata/blob/master/data/misc/edrp_calo-enron_maildir-users.txt), mostly senior management of Enron, organized into folders.
3. The corpus contains a total of about 0.5M messages.
4. This data was originally made public, and posted to the web, by the Federal Energy Regulatory Commission during its investigation.
5. The email dataset was later purchased by [Leslie Kaelbling](http://people.csail.mit.edu/lpk/) at MIT, and turned out to have a number of integrity problems. A number of folks at SRI, notably [Melinda Gervasio](https://www.sri.com/people/melinda-gervasio/), worked hard to correct these problems, and it is thanks to them (not me) that the dataset is available.
6. The dataset here
  1. does not include attachments, and
  2. some messages have been deleted “as part of a redaction effort due to requests from affected employees”.
  3. Invalid email addresses were converted to something of the form user@enron.com whenever possible (i.e., recipient is specified in some parse-able format like “Doe, John” or “Mary K. Smith”) and to no_address@enron.com when no recipient was specified.

CALO correctly identified 8 duplicate, misspelled custodians in the FERC dataset, resulting in 150 CALO custodians vs. 158 FERC custodians.

In addition to the above, the CALO dataset has a number of optimizations:

1. Message-ID: New Message-IDs have been created and used in place of existing Message-IDs
2. Date: Dates have been canonicalized replacing the raw dates
3. Headers: Some other headers are missing from the email

Removing the attachments makes the dataset much more manageable in size. However, attachments are still useful for certain investigations and [Mark Dredze](http://www.cs.jhu.edu/~mdredze/) has created a version of the CALO dataset with attachment information brought over from the FERC dataset.

K. Krasnow Waterman discusses how these changes affect the email in [Knowledge Discovery in Corporate Email: The Compliance Bot Meets Enron](https://dspace.mit.edu/handle/1721.1/37574), 2006.
