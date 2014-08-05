go-meh
======

extremely basic scraper for daily deals site meh.com. i've only tested it on mavericks, with brew's go 1.3 and i'm sure it's not following best standards of anything but it works


## Return structure
```

{
  "Name": "2-for-Tuesday: Truefire Gourmet 18-Piece Set of Cedar Grilling Planks \u0026 Wraps",
  "Price": "$18",
  "Features": [
    "Recommended for one use only",
    "Must be soaked in water before use",
    "Smoky flavor for gas or charcoal grills, so there's less incentive to charcoal grill",
    "Keeps your grill clean: one less excuse for avoiding your family",
    "Sets a high bar for future dinner-table presentation"
  ],
  "Images": [
    "https://res.cloudinary.com/mediocre/image/upload/c_pad,f_auto,h_600,w_600/xzdldvkhxf90fxcdrruo.png",
    "https://res.cloudinary.com/mediocre/image/upload/c_pad,f_auto,h_600,w_600/rzhecayzyu5xgp7xlna0.png",
    "https://res.cloudinary.com/mediocre/image/upload/c_pad,f_auto,h_600,w_600/tmmur1dlc3owsx0goi5m.png",
    "https://res.cloudinary.com/mediocre/image/upload/c_pad,f_auto,h_600,w_600/i8sdkce4cu7g0f6dwgrf.png",
    "https://res.cloudinary.com/mediocre/image/upload/c_pad,f_auto,h_600,w_600/jonqe9bofrhyhp2pdwu9.png"
  ],
  "ForumLink": "https://meh.com/forum/topics/2-for-tuesday-truefire-gourmet-18-piece-set-of-cedar-grilling-planks--wrap",
  "VideoLink": "//www.youtube.com/embed/MPt9OQ6p7p4?autohide=1\u0026color=white\u0026showinfo=0\u0026theme=light",
  "StoryTitle": "You don't plank your own cedar?",
  "Story": [
    "I'm one of those people who likes to know where my cedar comes from. With all our \"modern\" \"civilized\" \"technology\", we've lost that feeling of connection to our grilling planks. Instead of passively consuming the mass production of the cedar-industrial complex, I harvest it from my own cedar grove. I care for each tree from seed to grill, learn their peculiarities, share their hopes and their fears.",
    "That plank your salmon was cooked on? Her name was Jennifer. Such a great sense of humor, a real jokester. Those wood chips in the yard? My friend Levi. He was such a cute sapling. And the cedar infusion in your cocktail, that's Mark. He had a great voice. If you lean down to your glass and listen very closely, you might hear him singing.  ",
    "Sure, butchering them is always difficult. But I honor their sacrifice with a Native American prayer of thanks that I made up, so we're cool. ",
    "Or, you know, you could just buy some cedar grilling planks, if you have the almighty dollar where your soul should be. These 12 planks will give your grilled meals the same sweet, smoky flavor, if decadent self-indulgence is all you care about. Shovel that lavish feast into your unthinking maw while you genuflect before your one true god, Consumerism. Mmmmm, yummy.",
    "I guess some people just don't want to put in a little effort for the 30 years it takes for a cedar tree to mature. Whatever. I don't judge. I'm a better person than that."
  ]
}
```

## Example
```
package main

import "github.com/go-martini/martini"
import "meh"

func main() {
  m := martini.Classic()
  m.Get("/meh.json", func() string {
                return meh.GetMeh()
        })
  m.Run()
}
```

## License

The [BSD 3-Clause license][bsd], the same as goquery and golang.
