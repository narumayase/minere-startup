# minere-startup

Application to manage eth minering.

## use cases

- when it starts it sends an email that serves as an alert to know that the pc started
- if the [db.data](https://github.com/narumayase/minere-startup/blob/main/db.data) file has written a number less than the configured threshold, then the mining begins, otherwise, it sends an email that serves as an alert to know that something has failed many times and requires manual review.

# build
Install [Golang](https://golang.org/) to build this application and execute the command below.
```
$ go build 
```

## pre-requisites

* Copy the executable created by ```go build``` to the os startup folder or configure the os to start this when it starts
* For now this only works with teamredminer and with the arguments:
    * ```-a ethash -o pool -u address.trmtest -p x```
    * copy [conf.json](https://github.com/narumayase/minere-startup/blob/main/examples/conf.json) besides the executable of this application and configure it by filling this data:

## conf.json

| params        |  description
|---------------|--------------------------------------------------
| **Threshold** | system restart tolerance threshold
| **Email**     | see Email description below
| **Miner**     | see Miner description below

| Email params     |  description
|------------------|--------------------------------------------------
| **From**         | email address from where this application sends the email
| **To**           | email address where this application sends the email
| **Smtp**         | smtp server
| **SmtpPort**     | smtp port
| **PasswordFrom** | password from the email from where this application sends the email
| **Start**        | see Start description below
| **Stop**         | see Stop description below

| Start/Stop params |  description
|-------------------|--------------------------------------------------
| **Body**          | body to send in the email, it can be plane text and even html format
| **Subject**       | subjecto to send in the email
| **Send**          | if is allowed to send this email. You can configure that this application don't send start or stop email

| Miner params |  description
|--------------|--------------------------------------------------
| **Path**     | file path where the executable of teamredminer is
| **User**     | wallet address of the user + .trmtest string
| **Pool**     | mining pool

* example
```
{
  "Email": {
    "Start": {
      "Body": "Hi! I'm Ms. Minere Traba Jadora, I work for you, do you remember? Well.. I had some issues, I don't know what happened but you know, I just crashed.\nBut hey! I'm on again and I will try to start the mining again.\n\nBest regards.",
      "Subject": "ALERT - Minere needs your help!",
      "Send": true
    },
    "Stop": {
      "Body": "Errmm... It's me... again. \nI restarted the system two times and I'm afraid to continue.. Can you tell me what to do?\nI'm sorry",
      "Subject": "ALERT - Minere REALLY needs your help!",
      "Send": true
    },
    "From": "****@gmail.com",
    "To": "****@gmail.com",
    "Smtp": "smtp.gmail.com:587",
    "SmtpPort": 587,
    "PasswordFrom": "****"
  },
  "Miner": {
    "Path": "C:\\Users\\Default\\Documents\\teamredminer-v0.8.5-win\\teamredminer.exe",
    "User": "*****.trmtest",
    "Pool": "stratum+tcp://us1.ethermine.org:4444"
  },
  "Threshold": 2
}
```



