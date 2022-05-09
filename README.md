# prioritiser

Suppose I have a list of things I want to put in priority order—my ten favourite spy novels, my most urgent work tasks, the movies I most want to watch, whatever it might be. It’s actually quite hard for us to prioritize lists in that way. It’s easier if we just see two things side by side: we can usually pick which one of the two is higher priority. 

The Prioritizer will take a list of arbitrary things and present them to the user in just this way, asking enough questions to put the list in priority order, but no more than necessary. It should also be able to insert a new item to the list in the right place, again with the minimal number of questions. And, of course, it should be a reusable Go library that doesn’t depend on a terminal or a CLI or a database or any particular object type :slightly_smiling_face:


Keeping it simple. Use the minimum dependencies you need to get the job done. No gRPC, no complicated cloud services, no unnecessary packages or interfaces.

Focusing on key user stories first. Build and release a working app that does one thing, rather than a nearly-working app that nearly does five things. It's easier to get the basic design right with something very simple, and once you have the right design it's much easier to add features to it later.

