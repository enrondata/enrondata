# Deduplication and Attachment Stripping – Reducing the Dataset

While this project’s initial goal is to create original, non-deduped datasets, oftentimes, the full dataset is not needed. Sometimes duplicates are not desired and sometimes attachments are not desired. The challenge is to meet this requirements while maintaining a realistic dataset.

One of the challenges with deduping is which duplicate do you remove and do you leave a link behind? For example, if Alice sends a message to Bob, it will typically exist in at least 3 places, in Alice’s Sent folder, in Alice’s Inbox and in Bob’s Inbox. If you were to remove two of those, which two would you remove and how representative would the resulting dataset be?

There seem to be three solutions to this depending on which problem you are solving.

1. Single-Instance Storage (Loss-less Storage Reduction): If storage is a problem, email archiving solutions solve this through SIS, or Single-Instance Storage where multiple copies of an email are stored only once. Emails on the mail server can be replaced by stubs which are emails where the body consists of a pointer to the full email. This way all records are accounted for but the storage costs are dramatically reduced for duplicates.
2. Attachment Elimination (Lossy Storage Reduction): Dataset size can be reduced even more by elimination attachments entirely while including attachment information either in the header or body. This can be created to simulate email with Attachment Stubbing where the attachments are no longer available.
3. Journal Email (Duplicate Elimination): If the goal is to eliminate duplicates entirely, maintaining user mailbox folders becomes problematic because certain folders will be missing email while others will not be. One way to address this problem is to eliminate folders entirely and move all the email into one or a few global folders, say organized by date instead of user. This is similar to how email archiving via journaling works. With journaling, a copy of each email sent or received, often eliminating duplicates, making it seem like an ideal approach for this requirement.

By using the above approaches, we can reduce the number of duplicates and the storage requirements while maintaining the characteristics of real world email datasets. I’ve added these to the proposed dataset list for consideration. Please let me know if you think these or other datasets would be of use. For now, I’ve added these to the projects page.
