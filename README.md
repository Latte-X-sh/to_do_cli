# to_do_cli

Commands:   
    - Are verbs (doing words)
    - Short and clear

Args:   
    - Input required
    - Space separated
    - Can be more than one
    - Nouns

    - Language of Commands:
        - Verb args0 args1
        - They are pronounceable
        - verb plus nouns

Options:
    - Flags:
        - Prefixed (double hyphen, /Flag etc)
        - Long and Short options
        - Stackable or can be combined
        - Input to Modify Behaviour
        - Language of Flags:
            - verb adverb noun
            (rm --force [file])
            - pronounceable
            (forcefully remove this file)
    - Adverbs:
        - Express manner, place, time , frequency, degree, level of certainty.
        - They answer, the how, in what way, when, where and to what extent

CLI APPS:
    - The name of the cli is a Noun.For example git, emacs, vi, http
    - Apps will help you launch something, do more than one thing.
    - It is a collection of commands
    Sub-commands:
        - for example svn add, brew install, git clone etc.
        Example:
        brew  fetch   -v    hugo
        (app) (cmd) (flag)  (arf)
        brew   install  hugo
        (noun) (verb)   (object)

App Design:
    - App features:
        - Add Todo
        - List Todos
        - Mark as done
        - Search and Filter
        - Priorities
        - Archive
        - Edit
        - Create Dates
        - Due Dates
        - Tags
        - Projects
