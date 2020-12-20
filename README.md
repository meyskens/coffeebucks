# Coffeebucks - pay off that coffee machine

This project started when I invested in a quite expensive coffee machine for my home. I took the amount out of my savings account as an investment.
Then I wondered, how long would it take to pay it back to myself if I were to pay Starbucks prices for my coffees every day.

[Demo](https://twitter.com/MaartjeME/status/1340697330639654912?s=20)

This project uses [Bunq](https://bunq.com/invite/Maartje) and an ESP32 (with giant button for fun!) to pay myself 2 euros every time I drink a coffee.
Due to the quite important part of having a good TLS stack (and lack of a good C library) I added a backend service in Go that does the actual transaction. This code is in Go and in `backend/`

## Hardware setup
I used an Adafruit Feather ESP32, which has a Cherry switch connected from A0 to GND. For fun it is inside a [3D printed panic button](https://www.thingiverse.com/thing:1406545). The ESP32 is setup to go into deepsleep and has an interupt when A0 goes low to connect to the WiFi and call the backend to make a payement.