# Prioritiser

Suppose I have a list of things I want to put in priority order—my ten favourite spy novels, my most urgent work tasks, the movies I most want to watch, whatever it might be. It’s actually quite hard for us to prioritize lists in that way. It’s easier if we just see two things side by side: we can usually pick which one of the two is higher priority. 


The Prioritizer will take a list of arbitrary things and present them to the user in just this way, asking enough questions to put the list in priority order, but no more than necessary.


1. Sort a list of unsorted items
```
prioritise itemsToSort.txt
...
Output:
Sorted Priorities:
Higginbotham - Midnight in Chernobyl
Martin - The True Believers
Pearson - The Profession of Violence
```

2. Add an item into a presorted list
```
prioritise -add sortedItems.txt

Please add your new item.
To exit, type Q and enter.

Sorted Priorities:
Higginbotham - Midnight in Chernobyl
Martin - The True Believers
New item! - The newest book
Pearson - The Profession of Violence
```

3. Modify an existing list on disk in place
```
// Sort the list in place, overwriting the original file 
prioritise -i books.txt

```
Prioritiser tries it's hardest to respect the users time by efficiently slotting a new priority into a previous list asking the minimum number of questions.