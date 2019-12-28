# Let's go to Japan! 日本にいこう！

![image](https://img.shields.io/badge/PullRequests-welcome-green.svg)
![image](https://img.shields.io/badge/FeatureRequests-welcome-green.svg)

I create this project as I want to know if there  some cheap flights to some cities of Japan in the incomming weekends.

There are three major LCCs which I'm planning to support in this project:

| Airline                                           | Status                 |
|---------------------------------------------------| :---:                  |
| [Peach](https://booking.flypeach.com/tw/search)   | :white_check_mark:     |
| [Jetstar](https://www.jetstar.com/tw/zh/home)     | In Progress            |
| [Tiger](https://www.tigerairtw.com/zh-tw/)        | :warning: (see [Tiger Airline](#tiger-airline)) |


Feel free to ask for features!


# Example

See [this example](//github.com/mkfsn/flyjapan/blob/master/examples/weekend.go)

![image](https://user-images.githubusercontent.com/667169/54473678-e8dfb980-4815-11e9-96c7-9ff463f5da9c.png)


# Tiger Airline

As [Tiger Airline](https://www.tigerairtw.com/zh-tw/) is using [reCaptcha](https://www.google.com/recaptcha/intro/v3.html) when querying the flight information, it's hard to write a crawler to fetch the detailed flight information.

The current implementation is to fetch the cache result (see [one of the cache page](https://static.tigerairtw.com/fare-cache/TPE:HND:TWD:2020-01)) so there's no flight information but only fare information per day.
