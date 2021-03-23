import smtplib
import sys
import json
import time

data = json.load(sys.stdin)


host = data["host"]["value"]
username = data["username"]["value"]
password = data["password"]["value"]
port = data["port"]["value"]

sender = "Private Person <from@example.com>"
receiver = "A Test User <to@example.com>"
time = time.time()

message = f"""\
Subject: Hi Mailtrap {time}
To: {receiver}
From: {sender}

This is a test e-mail message from python using the terraform provider. If you see this it is working."""

with smtplib.SMTP(host, 2525) as server:
    print(server.login(username, password))
    print(server.sendmail(sender, receiver, message))
