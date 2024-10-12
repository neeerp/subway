# Random Subway Station Exploration

This is a simple webpage hosted on AWS Lambda using go templ + HTMX. It suggests a random subway station to explore in Toronto!

See this hosted at https://subway.davidg.xyz

## Build Instructions
Run `./bundle.sh` to regenerate templates, bundle lambda handlers, and compile CSS. 

Afterwards, you have to manually deploy the handlers to Lambda, and you need to manually drop `style.css` to the (hardcoded) S3 bucket. Yes, it's painful :)

## TODO List
[ ] Automate deployment to AWS so I don't have to manually upload this stuff every time I make a change
[ ] Use Tailwind CSS to make things ✨ p r e t t y ✨
